package goutils

import (
	"net/smtp"
	"strings"
)

func sendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func SendEmail(to_email []string, subject, body string) error {
	user := "1460586781@qq.com"
	password := "xmlpdvspdkxeiefh"
	host := "smtp.qq.com:587"

	to_stirng := strings.Join(to_email, ";")
	return sendToMail(user, password, host, to_stirng, subject, body, "text")
}

func Notify(subject, message string, emails []string) {
	SendEmail(emails, subject, message)
}
