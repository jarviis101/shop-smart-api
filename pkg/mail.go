package pkg

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

func CreateMailDialer(cfg Mailer) *gomail.Dialer {
	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d
}
