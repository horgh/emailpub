// Package emailpub is a Cloud Function that gets triggered by Pub/Sub
// messages. It sends an email.
package emailpub

import (
	"context"
	"encoding/json"
	"net/smtp"
	"os"

	"github.com/jhillyerd/enmime"
	"github.com/pkg/errors"
)

// EmailPub consumes a Pub/Sub message.
func EmailPub(ctx context.Context, psm pubSubMessage) error {
	var m message
	if err := json.Unmarshal(psm.Data, &m); err != nil {
		return errors.Wrap(err, "error unmarshaling")
	}

	if m.Body == "" || m.Subject == "" || m.To == "" {
		return errors.New("message is not fully populated")
	}

	if err := sendEmail(m.Subject, m.To, m.Body); err != nil {
		return err
	}

	return nil
}

type pubSubMessage struct {
	Data []byte `json:"data"`
}

type message struct {
	Body    string
	Subject string
	To      string
}

func sendEmail(subject, to, body string) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	if host == "" || port == "" || user == "" || pass == "" {
		return errors.New("missing SMTP environment variables")
	}

	auth := smtp.PlainAuth("", user, pass, host)

	master := enmime.Builder().
		From("", user).
		Subject(subject).
		Text([]byte(body))

	if err := master.To("", to).Send(host+":"+port, auth); err != nil {
		return errors.Wrap(err, "error sending email")
	}

	return nil
}
