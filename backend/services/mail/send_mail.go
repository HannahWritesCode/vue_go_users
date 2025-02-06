package mail

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendResetEmail(receiver string, reset_token string) error {
	// Sender data
	from := os.Getenv("SENDER_EMAIL")
	password := os.Getenv("APP_PASSWORD")

	// Receiver email address
	to := []string{
		receiver,
	}

	// smtp server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Subject and body
	subject := "Reset Your Password"
	body := fmt.Sprintf("Reset your password here: http://localhost:5173/reset-password/%s", reset_token)

	// Build the message
	message := fmt.Sprintf("From: %s\r\n", from)
	message += fmt.Sprintf("To: %s\r\n", to)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += fmt.Sprintf("\r\n%s\r\n", body)

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	return err
}
