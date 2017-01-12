package services

import (
	"strings"
	"net/smtp"
	"fmt"
)

func SendToMail(user, password, host, to, subject, body, mailtype string, chanFail, chanSuccess chan int) bool {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	if err != nil {
		fmt.Println(err)
		chanFail <- 1
		return false
	}
	chanSuccess <- 2
	return true
}