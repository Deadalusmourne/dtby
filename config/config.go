package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	DatabaseUser string `yaml:"db_user"`
	DatabasePwd string `yaml:"db_password"`
	DatabaseName string `yaml:"db_name"`
	DatabaseHost string `yaml:"db_host"`
	DatabasePort string `yaml:"db_port"`
	DSN string
}

var Config Configuration

func LoadConfig(path string) error{
	data, err := ioutil.ReadFile(path)
	if err != nil{
		return err
	}
	err = yaml.Unmarshal(data, &Config)
	if err!=nil{
		return err
	}
	dsn := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		Config.DatabaseUser, Config.DatabasePwd, Config.DatabaseHost, Config.DatabasePort, Config.DatabaseName)
	fmt.Printf("DSN: %v\n", dsn)
	Config.DSN = dsn
	return nil
}