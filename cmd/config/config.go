package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Host     string
		Port     string
		Username string
		Password string
		Schema   string
	}

	RabbitMQ struct {
		Host     string
		Port     string
		Username string
		Password string
	}
}

func NewConfig() (*Config, error) {
	file, err := os.Open("resources/config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	config := &Config{}
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	config.Database.Host = os.Getenv("MYSQL_HOST")
	config.Database.Port = os.Getenv("MYSQL_PORT")
	config.Database.Username = os.Getenv("MYSQL_USERNAME")
	config.Database.Password = os.Getenv("MYSQL_PASSWORD")
	config.Database.Schema = os.Getenv("MYSQL_SCHEMA")

	config.RabbitMQ.Host = os.Getenv("RABBITMQ_HOST")
	config.RabbitMQ.Port = os.Getenv("RABBITMQ_PORT")
	config.RabbitMQ.Username = os.Getenv("RABBITMQ_USERNAME")
	config.RabbitMQ.Password = os.Getenv("RABBITMQ_PASSWORD")

	return config, nil
}
