package methods

import "github.com/go-gomail/gomail"

type EmailTo struct {
	To      string //收件人
	Subject string //主题
	Body    string //内容
}

// SendEmail 发送邮件
func SendEmail(e *EmailTo) {
	// 创建一个发送邮件的对象
	msg := gomail.NewMessage()
	// 设置发件人
	msg.SetHeader("From", msg.FormatAddress("xu756top@163.com", "阿新网"))
	// 设置收件人
	msg.SetHeader("To", e.To)
	// 设置主题
	msg.SetHeader("Subject", e.Subject)
	// 设置内容
	msg.SetBody("text/html", e.Body)
	// 创建一个发送邮件的对象
	d := gomail.NewDialer("smtp.163.com", 25, "xu756top@163.com", "SCYWTIRMYVWATFMN")
	// 发送邮件
	if err := d.DialAndSend(msg); err != nil {
		println("发送邮件失败", err.Error())
		return
	}
}
