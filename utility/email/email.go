package email

import (
	"crypto/tls"
	"github.com/iimeta/fastapi-admin/internal/config"
	"gopkg.in/gomail.v2"
)

type Message struct {
	To      []string // 收件人
	Subject string   // 邮件主题
	Body    string   // 邮件正文
}

type Dialer struct {
	Host     string // smtp.xxx.com
	Port     int    // 端口号
	UserName string // 发信账号
	Password string // 发信密码
	FromName string // 发信人名称
}

type OptionFunc func(msg *gomail.Message)

func NewMessage(to []string, subject, body string) *Message {
	return &Message{
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

func NewDialer(host string, port int, username, password string) *Dialer {
	return &Dialer{
		Host:     host,
		Port:     port,
		UserName: username,
		Password: password,
	}
}

func NewDefaultDialer() *Dialer {
	return &Dialer{
		Host:     config.Cfg.Email.Host,
		Port:     config.Cfg.Email.Port,
		UserName: config.Cfg.Email.UserName,
		Password: config.Cfg.Email.Password,
	}
}

func SendMail(message *Message, dialer *Dialer, opt ...OptionFunc) error {

	m := gomail.NewMessage()

	if dialer == nil {
		dialer = NewDefaultDialer()
	}

	// 这种方式可以添加别名, 即“XX官方”
	m.SetHeader("From", m.FormatAddress(dialer.UserName, dialer.FromName))

	if len(message.To) > 0 {
		m.SetHeader("To", message.To...)
	}

	if len(message.Subject) > 0 {
		m.SetHeader("Subject", message.Subject)
	}

	if len(message.Body) > 0 {
		m.SetBody("text/html", message.Body)
	}

	// m.SetHeader("Cc", m.FormatAddress("xxxx@foxmail.com", "收件人")) //抄送
	// m.SetHeader("Bcc", m.FormatAddress("xxxx@gmail.com", "收件人"))  //暗送

	for _, o := range opt {
		o(m)
	}

	return do(dialer, m)
}

func do(dialer *Dialer, msg *gomail.Message) error {

	d := gomail.NewDialer(dialer.Host, dialer.Port, dialer.UserName, dialer.Password)

	// 默认开启SSL
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d.DialAndSend(msg)
}
