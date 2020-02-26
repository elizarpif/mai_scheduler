package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Address     string      `yaml:"address"`
	Departments Departments `yaml:"departments"`
}

type Departments struct {
	Inst1  int `yaml:"inst1"`
	Inst2  int `yaml:"inst2"`
	Inst3  int `yaml:"inst3"`
	Inst4  int `yaml:"inst4"`
	Inst5  int `yaml:"inst5"`
	Inst6  int `yaml:"inst6"`
	Inst7  int `yaml:"inst7"`
	Inst8  int `yaml:"inst8"`
	Inst9  int `yaml:"inst9"`
	Inst10 int `yaml:"inst10"`
	Inst11 int `yaml:"inst11"`
	Inst12 int `yaml:"inst12"`
}

//func getSchedule
func GetConfig(config string) (*Config, error) {
	cnf := new(Config)
	data, err := ioutil.ReadFile(config)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, cnf)
	if err != nil {
		return nil, err
	}

	return cnf, err
}
