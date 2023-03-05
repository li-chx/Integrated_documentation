package service

import (
	"fmt"
	"github.com/Haroxa/Integrated_documentation/common"
	"github.com/Haroxa/Integrated_documentation/config"
	"time"
)

func getHtmlContent(MsgType int, DesireContent string, MessageContent string) string {
	SendTime := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", time.Now().In(common.ChinaTime).Year(), time.Now().In(common.ChinaTime).Month(), time.Now().In(common.ChinaTime).Day(), time.Now().In(common.ChinaTime).Hour(), time.Now().In(common.ChinaTime).Minute(), time.Now().In(common.ChinaTime).Second())
	html := ""
	switch MsgType {
	case 1:
		html = fmt.Sprintf(`<div>
						<div>
							叮咚ding~
						</div>
						<div>
							<p> 属于你的小幸运已经被签收! </p>
						</div>
						<div>
							<p> 有人点亮你在小幸运活动中的愿望啦~ </p>
							<p> 快去看看Ta是谁吧! </p>
						</div>
						<div>
							<p> 此邮箱为系统邮箱，请勿回复。</p>
							<p> 发送于 %s </p>
						</div>
					<div>`, SendTime)
	case 2:
		html = fmt.Sprintf(`<div>
						<div>
							叮咚ding~
						</div>
						<div>
							<p> Ta取消了对你 </p>
						</div>
						<div style="padding: 8px 40px 8px 50px;">
							<p>" %s "</p>
						</div>
						<div> 愿望的点亮，并留言: "%s" </div>
						<div>
							<p>该愿望重新被投入到愿望池</p>
							<p>耐心等待下一个有缘人点亮它吧</p>
							<p>惊喜总在不经意间~</p>
						</div>
						<div>
							<p> 此邮箱为系统邮箱，请勿回复。</p>
							<p> 发送于 %s </p>
						</div>
					<div>`, DesireContent, MessageContent, SendTime)
	case 3:
		html = fmt.Sprintf(`<div>
								<div>
									叮咚ding~
								</div>
								<div>
									<p> 你的愿望已经成功被%s同学实现了 </p>
								</div>
								<div>
									<p> 再去愿望池看看其它愿望吧~ </p>
								</div>
								<div>
									<p> 此邮箱为系统邮箱，请勿回复。</p>
									<p> 发送于 %s </p>
								</div>
							<div>`, DesireContent, SendTime)
	case 4:
		html = fmt.Sprintf(`<div>
								<div>
									叮咚ding~
								</div>
								<div>
									<p> Ta删除了愿望 </p>
								</div>
								<div style="padding: 8px 40px 8px 50px;">
									<p> "%s" </p>
								</div>
								<div>
									<p> 再去愿望池看看其它愿望吧~ </p>
								</div>
								<div>
									<p> 此邮箱为系统邮箱，请勿回复。</p>
									<p> 发送于 %s </p>
								</div>
							<div>`, DesireContent, SendTime)
	}
	return html
}

func SendMail(EmailAddress string, MsgType int, DesireContent string, MessageContent string) error {

	mailConfig := config.GetMailConfig()
	html := getHtmlContent(MsgType, DesireContent, MessageContent)
	message := gomail.NewMessage()
	message.SetAddressHeader("From", mailConfig.From, mailConfig.FromName)
	message.SetHeader("To", EmailAddress)
	message.SetHeader("Subject", "[集成文档2023]邮件通知")
	message.SetBody("text/html", html)

	dia := gomail.NewDialer(mailConfig.Host, mailConfig.Port, mailConfig.Username, mailConfig.Password)

	err := dia.DialAndSend(message)

	return err
}
