package main

import (
	"fmt"
	"os"

	"mai_scheduler/config"
	"mai_scheduler/mai"
	"mai_scheduler/telegramBot"
)

func main() {
	cfg, err := config.GetConfig("config/config.yaml")
	if err != nil {
		fmt.Println("config!")
		return
	}

	service := new(mai.MaiClient)
	service.Config = cfg
	service.InitClient()

	telegramBot.TelegramBot(service, os.Getenv("BOTTOKEN"))
}
