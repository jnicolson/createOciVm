package main

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/gomail.v2"
)

type SmtpServer struct {
	Address  string
	Port     int
	Username string
	Password string
}

func SendNotification(from, to string, server SmtpServer) {

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Oracle VM Create")
	m.SetBody("text/plain", "A VM was created")

	d := gomail.NewDialer(server.Address, server.Port, server.Username, server.Password)

	err := d.DialAndSend(m)

	if err != nil {
		log.Fatal().Err(err).Msg("Sending Email")
	}

}
