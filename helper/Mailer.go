package helper

import (
	"fmt"
	"github.com/Haroxa/Integrated_documentation/common"
	"github.com/Haroxa/Integrated_documentation/config"
	"gopkg.in/gomail.v2"
	"time"
)

func getHtmlContent(MsgType int, MessageContent string) string {
	SendTime := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", time.Now().In(common.ChinaTime).Year(), time.Now().In(common.ChinaTime).Month(), time.Now().In(common.ChinaTime).Day(), time.Now().In(common.ChinaTime).Hour(), time.Now().In(common.ChinaTime).Minute(), time.Now().In(common.ChinaTime).Second())
	html := ""
	switch MsgType {
	case 1:
		html = fmt.Sprintf(`<div>
						<div>
							叮咚ding~
						</div>
						<div>
							<p> 你正在注册集成文档的用户! </p>
						</div>
						<div>
							<p> 你的注册验证码为 </p>
							<p> %s </p>
						</div>
						<div>
							<p> 请在一分钟内使用，逾期将失效</p>
							<p> 发送于 %s </p>
						</div>
					<div>`, MessageContent, SendTime)
	}
	return html
}

func SendMail(EmailAddress string, MsgType int, MessageContent string) error {

	mailConfig := config.GetMailConfig()
	html := getHtmlContent(MsgType, MessageContent)
	message := gomail.NewMessage()
	message.SetAddressHeader("From", mailConfig.From, mailConfig.FromName)
	message.SetHeader("To", EmailAddress)
	message.SetHeader("Subject", "[集成文档2023]邮件通知")
	message.SetBody("text/html", html)

	dia := gomail.NewDialer(mailConfig.Host, mailConfig.Port, mailConfig.Username, mailConfig.Password)

	err := dia.DialAndSend(message)

	return err
}
