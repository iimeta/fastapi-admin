package email

import (
	"crypto/tls"

	"github.com/gogf/gf/v2/text/gregex"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/errors"
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

func NewDialer(host string, port int, userName, password, fromName string) *Dialer {
	return &Dialer{
		Host:     host,
		Port:     port,
		UserName: userName,
		Password: password,
		FromName: fromName,
	}
}

func NewDefaultDialer() *Dialer {
	return &Dialer{
		Host:     config.Cfg.Email.Host,
		Port:     config.Cfg.Email.Port,
		UserName: config.Cfg.Email.UserName,
		Password: config.Cfg.Email.Password,
		FromName: config.Cfg.Email.FromName,
	}
}

func SendMail(message *Message, dialer *Dialer, opt ...OptionFunc) error {

	m := gomail.NewMessage()

	if dialer == nil {
		dialer = NewDefaultDialer()
	}

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

func Verify(email string) error {

	if ok := gregex.IsMatchString(`^[a-zA-Z0-9_\-\.+]+@[a-zA-Z0-9_\-]+(\.[a-zA-Z0-9_\-]+)+$`, email); ok {
		return nil
	}

	return errors.Newf("The `%s` is not a valid email address", email)
}
