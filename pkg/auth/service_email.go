package auth

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

const (
	pathToConf = "./cmd/config/service_email.yaml"
)

const (
	message = `
	Hello!
	Your email was specified to restore the password to log in to the personal account of the online store BShop.
	Your code: %d

	Sincerely,
	the BShop team.
	`

	subject = "Subject: Password recovery in BShop account\r\n"
)

type ServiceEmail struct {
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint32 `yaml:"port"`
}

func InitServiceEmail() *ServiceEmail {
	f, err := os.ReadFile(pathToConf)
	if err != nil {
		log.Fatal(err)
	}

	var se ServiceEmail
	if err = yaml.Unmarshal(f, &se); err != nil {
		log.Fatal(err)
	}
	return &se
}

func (s *ServiceEmail) SendCode(to string, code int) error {
	auth := smtp.PlainAuth("", s.Email, s.Password, s.Host)

	err := smtp.SendMail(
		s.Host+":"+strconv.Itoa(int(s.Port)),
		auth, s.Email, []string{to}, []byte(subject+fmt.Sprintf(message, code)))

	if err != nil {
		return err
	}

	return nil
}
