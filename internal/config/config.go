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
type Server struct {
	Host string
	Port uint
}

var c struct {
	Environment
	Postgres
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

func GetServer() Server {
	return c.Server
}
