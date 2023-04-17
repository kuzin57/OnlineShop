package db

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Connection struct {
		Host     string `yaml:"host"`
		Port     uint32 `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBname   string `yaml:"dbname"`
	} `yaml:"connection"`
}

func getDBConfig(pathToConfig string) DatabaseConfig {
	f, err := os.ReadFile(pathToConfig)
	if err != nil {
		log.Fatal(err)
	}

	var dc DatabaseConfig
	if err = yaml.Unmarshal(f, &dc); err != nil {
		log.Fatal(err)
	}
	return dc
}
