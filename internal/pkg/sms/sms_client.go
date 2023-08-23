package sms

import (
	"fmt"
	sms "github.com/dmitriy-borisov/go-smsru"
	"log"
)

type Client interface {
	Send(phone, code string)
}

type smsClient struct {
	client  *sms.SmsClient
	isDebug bool
}

func CreateClient(c *sms.SmsClient, d bool) Client {
	return &smsClient{c, d}
}

func (c *smsClient) Send(phone, code string) {
	if c.isDebug {
		return
	}

	msg := sms.NewSms(phone, fmt.Sprintf("ОТП код: %s", code))

	res, err := c.client.SmsSend(msg)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}

	log.Printf("Status = %d, Ids = %v, Balance = %f", res.Status, res.Ids, res.Balance)
}
