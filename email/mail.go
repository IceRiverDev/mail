package email

import (
	"gopkg.in/gomail.v2"
)

func SendMail() {
	mc := &MailConfigs{}
	mc.setMailConfigs()

	m := gomail.NewMessage()
	m.SetHeader("From", mc.cmd.from)
	m.SetHeader("To", mc.cmd.to)
	m.SetHeader("Subject", mc.cmd.subject)
	m.SetBody("text/html", mc.cmd.body)

	d := gomail.NewDialer(mc.Address, mc.Port, mc.User, mc.Password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
