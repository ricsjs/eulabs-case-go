package config

import (
	"fmt"
	"log"
)

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
		Host:     "localhost",
		Port:     "3306",
		DBName:   "eulabscase",
	}
}

func (c *Config) GetDSN() string {
	log.Println(c.Host)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
}
