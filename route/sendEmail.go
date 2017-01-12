package route

import (
	"fmt"
	"net/http"
	"go-run/services"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {

	user := "1047887945@qq.com"
	password := "yimtpishgivdbbhg"
	host := "smtp.qq.com:25"
	to := "towne766@126.com"

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
	chanFail := make(chan int)
	chanSuccess := make(chan int)
	go services.SendToMail(user, password, host, to, subject, body, "html", chanFail, chanSuccess)

	for {
		select {
		case <-chanFail:
			fmt.Printf("c1:%d ", chanFail)
			w.Write([]byte("error"))
			return
		case <-chanSuccess:
			fmt.Printf("c1:%d ", chanSuccess)
			w.Write([]byte("success"))
			return
		}
	}

}
