package auth

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	pathToConf = "./cmd/config/service_email.yaml"
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
