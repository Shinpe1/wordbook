package util

import (
	"log"
	"net/smtp"
	"os"
)

func SendMail(to string, token string) error {
	log.Println("#SendMail start;")

	a := auth{
		username: os.Getenv("MAIL_USERNAME"),
		from:     os.Getenv("MAIL_FROM"),
		password: os.Getenv("MAIL_PASS"),
		server:   os.Getenv("MAIL_SERVER"),
	}

	auth := smtp.PlainAuth(a.username, a.from, a.password, a.server)

	m := mail{
		from: os.Getenv("MAIL_FROM"),
		to:   to,
		sub:  "Please confirm your e-mail address!!",
		host: os.Getenv("MAIL_HOST"),
		msg: "Please confirm your e-mail address by clicking the following URL\r\n\r\n" +
			createURL(token) +
			"\r\n\r\n(this URL will be expired in 60 min.)" +
			"\r\n\r\nBest regards.\twordbook-anywhere.",
	}

	if err := smtp.SendMail(m.host, auth, m.from, []string{m.to}, []byte(m.body())); err != nil {
		log.Println("sending mail failed;")
		log.Fatal(err.Error())
		return err
	}

	log.Println("#SendMail end;")
	return nil
}

type mail struct {
	from string
	to   string
	sub  string
	host string
	msg  string
}

func (m mail) body() string {
	return "To: " + m.to + "\r\n" +
		"Subject: " + m.sub + "\r\n\r\n" +
		m.msg + "\r\n"
}

type auth struct {
	username string
	from     string
	password string
	server   string
}

func createURL(token string) string {
	return "http://localhost:" + os.Getenv("PORT") + "/api/v1.0/auth/temp?token=" + token
}
