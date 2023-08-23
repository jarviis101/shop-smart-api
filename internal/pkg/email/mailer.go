package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
)

type Mailer interface {
	Send(email, code string)
}

type mailer struct {
	d *gomail.Dialer
}

func CreateMailer(d *gomail.Dialer) Mailer {
	return &mailer{d}
}

func (m *mailer) Send(email, code string) {
	message := gomail.NewMessage()
	message.SetHeader("From", "from@gmail.com")
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Shop smart authentication")
	message.SetBody("text/plain", fmt.Sprintf("ОТП код: %s", code))

	if err := m.d.DialAndSend(message); err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}
}
