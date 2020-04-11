# emailpub

emailpub is a Google Cloud Platform Cloud Function. It gets triggered by
messages published to a Pub/Sub topic. Each message represents an email to
send.

It sends email via SMTP. The Pub/Sub messages include the subject,
recipient, and text body.

I created it because I wanted to send emails from hosts where I don't want
to expose SMTP credentials. With emailpub, I can give these hosts a GCP
service account allowing them to publish messages to send emails.

## Setup

Create `env.yml`:

```yaml
---
SMTP_HOST: smtp.gmail.com
SMTP_PORT: '587'
SMTP_USER: ...@gmail.com
SMTP_PASS: password...
```

Deploy the function (also creates the Pub/Sub topic):

```
gcloud functions deploy EmailPub \
  --runtime go113 \
    --trigger-topic emailpub \
    --source . \
    --env-vars-file env.yml
```
