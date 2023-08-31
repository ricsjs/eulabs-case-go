package config

import (
	"fmt"
)

//optei por deixar esse arquivo visualizável no repositório do github para ter acesso às configurações

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewConfig() *Config {
	return &Config{
		User:     "root",
		Password: "root",
		Host:     "db-container",
		Port:     "3306",
		DBName:   "eulabscase",
	}
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
}
