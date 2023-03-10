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
