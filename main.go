package main

import (
	"fmt"
	"mai_scheduler/cmd"
	"mai_scheduler/config"
	"net/url"
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

	info, err := service.GetGroupsByInstituteAndCourse("https://mai.ru/education/schedule/?department=151&course=3")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)

	//https://mai.ru/education/schedule/detail.php?group=%D0%9C8%D0%9E-311%D0%91-17
	str := fmt.Sprintf("https://mai.ru/education/schedule/detail.php?group=%s", url.QueryEscape(info.Groups[22]))

	fmt.Println(str)

	_,err = service.GetSchedule(str)
	if err != nil {
		fmt.Println(err)
	}

}
