package main

import (
	"fmt"
	"net/url"

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

	urlParse, err := url.Parse("https://mai.ru/education/schedule")
	if err != nil{
		fmt.Println("url!")
		return
	}
	query := urlParse.Query()
	query.Add("department","151")
	query.Add("course", "3")

	urlParse.RawQuery = query.Encode()

	service.GetGroupsByInstituteAndCourse(urlParse.String())

}
