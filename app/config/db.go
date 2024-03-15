package config

import (
	"fmt"
	"log"
	"os"
	"yugod-backend/app/util"

	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	Username     string `yaml:"Username"`
	Password     string `yaml:"Password"`
	Connection   string `yaml:"Connection"`
	DatabaseName string `yaml:"DatabaseName"`
}

var DB DBConfig

func (d DBConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True", d.Username, d.Password, d.Connection, d.DatabaseName)
}

func init() {
	// Read config from YAML
	var settings DBConfig

	if util.GetEnvBooleanValue("DB_USE_ENV") {
		settings.Username = os.Getenv("DB_USERNAME")
		settings.Password = os.Getenv("DB_PASSWORD")
		settings.Connection = os.Getenv("DB_CONNECTION")
		settings.DatabaseName = os.Getenv("DB_NAME")
	} else {
		config, err := os.ReadFile("config/db.yml")

		if err != nil {
			log.Fatal("DB config not set.")
		}
		yamlErr := yaml.Unmarshal(config, &settings)
		if yamlErr != nil {
			log.Fatal("DB config read error.")
		}
	}

	DB = settings
}
