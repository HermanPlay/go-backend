package config

import (
	"errors"
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

var errApiPort = errors.New("error parsing env variable port")
var errApiSecret = errors.New("error parsing env variable api_secret")
var errDbHost = errors.New("error parsing env variable db_host")
var errDbPort = errors.New("error parsing env variable db_port")
var errDbUser = errors.New("error parsing env variable db_user")
var errDbPassword = errors.New("error parsing env variable db_password")
var errDbName = errors.New("error parsing env variable db_name")

func GetConfig(env map[string]string) (*Config, error) {
	app_port, err := strconv.Atoi(env["port"])
	if err != nil {
		return nil, errApiPort
	}
	api_secret := env["api_secret"]
	if len(api_secret) == 0 {
		return nil, errApiSecret
	}
	app := App{
		Port:       app_port,
		Api_secret: api_secret,
	}
	db_host := env["db_host"]
	if len(db_host) == 0 {
		return nil, errDbHost
	}
	db_port, err := strconv.Atoi(env["db_port"])
	if err != nil {
		return nil, errDbPort
	}
	db_user := env["db_user"]
	if len(db_user) == 0 {
		return nil, errDbUser
	}
	db_password := env["db_password"]
	if len(db_password) == 0 {
		return nil, errDbPassword
	}
	db_name := env["db_name"]
	if len(db_name) == 0 {
		return nil, errDbName
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
	}, nil
}
