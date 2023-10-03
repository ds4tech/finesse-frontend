package models

import (
	"encoding/json"
	"io/ioutil"
)

var (
	Configuration Config
)

init() {
    args := os.Args[1:]
    if len(args) == 0{
       fmt.Println("********************\nMust specify a config file   in args\n********************")
       os.Exit(1)
   }

   Configuration = NewConfig(args[0]) // configuration initialized here
}

type Config struct {
	Db     map[string]string `json:"db"`
	Server map[string]string `json:"server"`
}

func NewConfig(fname string) *Config {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}
