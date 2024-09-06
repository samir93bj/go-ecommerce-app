package main

import (
	config "go-ecommerce-app/configs"
	"go-ecommerce-app/internal/api"
)

func main() {
	config, err := config.SetupEnv()

	if err != nil {
		panic(err)
	}

	api.StartServer(config)
}
