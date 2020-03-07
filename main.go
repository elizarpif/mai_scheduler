package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

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

	info, err := service.GetGroups(8, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info.Groups[20])

	schs, _ := service.GetScheduleByGroup(info.Groups[20])
	spew.Dump(schs)
}
