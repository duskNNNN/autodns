package MailSend

import (
	"log"
	"strconv"

	"github.com/spf13/viper"

	"gopkg.in/gomail.v2"
)

// 用于发送邮件
func MailSend(msg string) {
	admin_email := viper.GetString("mail.adminEmail")
	// 发送给管理员
	mailTo := []string{
		admin_email,
	}
	// 邮件主题
	var mail_subject string
	mail_body := msg
	// 设置邮箱主体
	mail_all := map[string]string{
		"user": viper.GetString("mail.sendEmail"),
		"pass": viper.GetString("mail.token"),
		"host": viper.GetString("mail.host"),
	}
	mailer := gomail.NewMessage(
		gomail.SetEncoding(gomail.Base64),
	)
	port, _ := strconv.Atoi(viper.GetString("mail.port"))
	// 设置邮件负载
	mailer.SetHeader("From", mail_all["user"])
	mailer.SetHeader("To", mailTo...)
	mailer.SetHeader("Subject", mail_subject)
	mailer.SetBody("text/html", mail_body)
	sender := gomail.NewDialer(mail_all["host"], port, mail_all["user"], mail_all["pass"])
	err := sender.DialAndSend(mailer)
	if err != nil {
		log.Println(err)
		return
	}
}
