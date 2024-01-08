package config

import (
	"log"
	"os"
	"strconv"
)

type (
	Config struct {
		App App
		Db  Db
	}

	App struct {
		Port       int
		Api_secret string
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}
)

func GetConfig() *Config {
	app_port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		log.Fatal("Error parsing env variable port. Error: ", err)
	}
	api_secret := os.Getenv("api_secret")
	app := App{
		Port:       app_port,
		Api_secret: api_secret,
	}
	db_host := os.Getenv("db_host")
	if len(db_host) == 0 {
		log.Fatal("Error parsing env variable db_host. Error: db_host is empty")
	}
	db_port, err := strconv.Atoi(os.Getenv("db_port"))
	if err != nil {
		log.Fatal("Error parsing env variable db_port. Error: ", err)
	}
	db_user := os.Getenv("db_user")
	if len(db_host) == 0 {
		log.Fatal("Error parsing env variable db_user. Error: db_user is empty")
	}
	db_password := os.Getenv("db_password")
	if len(db_host) == 0 {
		log.Fatal("Error parsing env variable db_password. Error: db_password is empty")
	}
	db_name := os.Getenv("db_name")
	if len(db_host) == 0 {
		log.Fatal("Error parsing env variable db_name. Error: db_name is empty")
	}

	db := Db{
		Host:     db_host,
		Port:     db_port,
		User:     db_user,
		Password: db_password,
		DBName:   db_name,
	}

	return &Config{
		App: app,
		Db:  db,
	}
}
