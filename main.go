package main

import (
	"fmt"

	"mai_scheduler/cmd"
	"mai_scheduler/config"
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

	service.GetGroupsByInstituteAndCourse("https://mai.ru/education/schedule/?department=151&course=3")

}
