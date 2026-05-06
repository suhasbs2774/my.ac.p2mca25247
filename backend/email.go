package main

import "gopkg.in/gomail.v2"

func SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", "your-email@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587,
		"suhasbs099@gmail.com",
		"", // use Gmail App Password
	)

	return d.DialAndSend(m)
}
