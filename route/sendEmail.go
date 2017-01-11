package route

import (
	"fmt"
	"net/smtp"
	"strings"
	"net/http"
)

func SendToMail(user, password, host, to, subject, body, mailtype string, chanErr chan bool) {
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
		chanErr <- true
	}
}

func SendEmail(w http.ResponseWriter, r *http.Request) {

	user := "1047887945@qq.com"
	password := "abcd1234"
	host := "smtp.exmail.qq.com:465"
	to := "harrytang@vipabc.com"

	subject := "使用Golang发送邮件"

	body := `
		<html>
		<body>
		<h3>
		"Test send to email"
		</h3>
		</body>
		</html>
		`
	fmt.Println("send email")
	chanErr := make(chan bool)
	go SendToMail(user, password, host, to, subject, body, "html", chanErr)

	for {
		select {
		case <-chanErr:
			fmt.Printf("c1:%d ", chanErr)
			w.Write([]byte("success"))
			return
		}
	}

}
