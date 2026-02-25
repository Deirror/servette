package mail

import (
	"fmt"
	"net/smtp"
)

type SMTPClient struct {
	host     string
	port     string
	username string
	password string
	from     string
}

func NewSMTPClient(cfg *Config) (*SMTPClient, error) {
	return &SMTPClient{
		host:     cfg.Host,
		port:     cfg.Port,
		username: cfg.Username,
		password: cfg.Password,
		from:     cfg.From,
	}, nil
}

func (s *SMTPClient) Send(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.username, s.password, s.host)

	msg := []byte("From: <" + s.from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		body + "\r\n")

	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	return smtp.SendMail(addr, auth, s.from, []string{to}, msg)
}
