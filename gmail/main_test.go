package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmailWithGmail(t *testing.T) {
	sender := NewGmailSender(Name, FromEmailAddress, FromEmailPassword)
	subject := "Test email to Nilllavees"
	content := `
	<h1>Hello dude</h1>
	<p>Just codding around here! </p>
	`
	to := []string{"myrachanto@gmail.com", "mutiapatrick35@gmail.com"}
	err := sender.SendEmail(subject, content, to, nil, nil, nil)
	require.NoError(t, err)
}
