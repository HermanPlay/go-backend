package main

import (
	"log"
	"strconv"

	"github.com/HermanPlay/backend/internal/api/http"
	"github.com/HermanPlay/backend/internal/api/http/server"
	"github.com/HermanPlay/backend/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	env, err := godotenv.Read("././.env")
	if err != nil {
		panic(err)
	}
	cfg, err := config.GetConfig(env)
	if err != nil {
		log.Fatal("No config!")
	}

	init := http.Init(cfg)
	app := server.Init(init)

	app.Run(":" + strconv.Itoa(cfg.App.Port))
}
