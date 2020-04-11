package emailpub

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendEmail(t *testing.T) {
	if os.Getenv("LIVE") == "" {
		t.Skip("Skipping test since LIVE is not set")
	}

	err := sendEmail("test subject", "will@summercat.com", "test body")
	require.NoError(t, err, "sent email")
}
