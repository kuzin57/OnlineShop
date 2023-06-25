package auth

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"

	"github.com/kuzin57/OnlineShop/pkg/db"
	"gopkg.in/yaml.v2"
)

const (
	pathToConf = "./cmd/config/service_email.yaml"
)

const (
	codeMessage = `
	Hello!
	Your email was specified to restore the password to log in to the personal account of the online store BShop.
	Your code: %d

	Sincerely,
	the BShop team.
	`

	orderMessage = `
	Hello!
	You have made an order on our site! Order ID is %d

	Sincerely,
	the BShop team.
	`

	subjectCodeMessage = "Subject: Password recovery in BShop account\r\n"

	subjectNotificationMessage = "Subject: You've just made an order in BShop\r\n"
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
		auth, s.Email, []string{to}, []byte(subjectCodeMessage+fmt.Sprintf(codeMessage, code)))

	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceEmail) SendOrderNotification(to string, order *db.Order) error {
	auth := smtp.PlainAuth("", s.Email, s.Password, s.Host)
	message := []byte(subjectNotificationMessage + fmt.Sprintf(orderMessage, order.Id))
	message = append(message, []byte("\nProducts: \n")...)

	for _, product := range order.Products {
		message = append(message, []byte(
			fmt.Sprintf(
				"-%s %s %d × %d руб. \n",
				product.Name, product.Brand, product.Amount, product.Price))...)
	}

	message = append(message, []byte(fmt.Sprintf("Total: %d руб.\n", order.TotalSum))...)

	err := smtp.SendMail(
		s.Host+":"+strconv.Itoa(int(s.Port)),
		auth, s.Email, []string{to},
		message)
	if err != nil {
		return err
	}

	return nil
}
