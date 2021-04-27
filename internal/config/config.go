package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Environment string

type Postgres struct {
	User     string
	Password string
	DBName   string
	Host     string
}

type Auth struct {
	Service string
	Host    string
}

type Server struct {
	Port uint
}

var c struct {
	Environment
	Postgres
	Auth
	Server
}

const (
	ENV_DEV  = "dev"
	ENV_PROD = "prod"
)

func init() {
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("can't open config file: %v", err)
	}

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&c); err != nil {
		log.Fatalf("can't decode config: %v", err)
	}
}

func GetEnvironment() Environment {
	return c.Environment
}

func GetPostgres() Postgres {
	return c.Postgres
}

func GetAuth() Auth {
	return c.Auth
}

func GetServer() Server {
	return c.Server
}
