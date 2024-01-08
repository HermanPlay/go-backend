package main

import (
	"strconv"

	"github.com/HermanPlay/backend/internal/api/http"
	"github.com/HermanPlay/backend/internal/api/http/server"
	"github.com/HermanPlay/backend/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("././.env")
	if err != nil {
		panic(err)
	}
	cfg := config.GetConfig()

	init := http.Init(cfg)
	app := server.Init(init)

	app.Run(":" + strconv.Itoa(cfg.App.Port))
}
