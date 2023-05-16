package handlers

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type PagesConfig struct {
	Home struct {
		Path      string   `yaml:"path"`
		Templates []string `yaml:"templates"`
	} `yaml:"home"`
	Catalogue struct {
		Path      string   `yaml:"path"`
		Templates []string `yaml:"templates"`
	} `yaml:"catalogue"`
	Auth struct {
		Path      string   `yaml:"path"`
		Templates []string `yaml:"templates"`
		AuthType  string   `yaml:"auth_type"`
	} `yaml:"auth"`
	Registration struct {
		Path      string   `yaml:"path"`
		Templates []string `yaml:"templates"`
	} `yaml:"registration"`
	PasswordRecovery struct {
		Path      string   `yaml:"path"`
		Templates []string `yaml:"templates"`
	} `yaml:"password_recovery"`
	Settings struct {
		Path      string   `yaml:"path"`
		Templates []string `yaml:"templates"`
	} `yaml:"settings"`
	Order struct {
		Path      string
		Templates []string `yaml:"templates"`
	} `yaml:"order"`
}

func GetHandlersParameters(pathToConf string) PagesConfig {
	f, err := os.ReadFile(pathToConf)
	if err != nil {
		log.Fatal(err)
	}

	var pc PagesConfig
	if err = yaml.Unmarshal(f, &pc); err != nil {
		log.Fatal(err)
	}
	return pc
}
