package main

import (
	"fmt"
	"mai_scheduler/cmd"
	"mai_scheduler/config"
	"mai_scheduler/telegramBot"
	"os"
)

func main() {
	cfg, err := config.GetConfig("config/config.yaml")
	if err != nil {
		fmt.Println("config!")
		return
	}

	service := new(cmd.MaiClient)
	service.Config = cfg
	service.InitClient()


	telegramBot.TelegramBot(service, os.Getenv("TOKEN"))
}
