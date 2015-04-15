package utils

import (
	"fmt"
	"github.com/nymoral/gothere/config"
	"net/smtp"
)

var auth smtp.Auth

func init() {
	auth = smtp.PlainAuth("",
		config.Config.MailUsername,
		config.Config.MailPassword,
		config.Config.MailHost)
}

type Message struct {
	To      string
	Subject string
	Body    string
}

func (m *Message) recipient() []string {
	strSlice := make([]string, 1)
	strSlice[0] = m.To
	return strSlice
}

func (m *Message) msg() []byte {
	return []byte(fmt.Sprintf("Subject: %s\n%s", m.Subject, m.Body))
}

func (msg *Message) Send() error {
	err := smtp.SendMail(config.Config.MailHost+":"+config.Config.MailPort,
		auth,
		config.Config.MailUsername,
		msg.recipient(),
		msg.msg())
	return err
}
