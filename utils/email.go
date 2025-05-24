package utils

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, body string) {
	email := os.Getenv("HOST_EMAIL")
	appPassword := os.Getenv("APP_PASSWORD")
	if email == "" || appPassword == "" {
		log.Println("Email or App Password not set in environment variables")
		return
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "sovajitr@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, email, appPassword)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send email to %s: %v", to, err)
	} else {
		log.Printf("Email sent to %s", to)
	}
}
