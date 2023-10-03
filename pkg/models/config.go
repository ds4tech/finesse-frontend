package models

import (
	"io/ioutil"
	"encoding/json"
 )
 
 type Config struct  {
	Db map[string]string `json:"db"`
	Server map[string]string `json:"server"`
 }
 
 
 func NewConfig(fname string) *Config{
	data,err := ioutil.ReadFile(fname)
	if err != nil{
	   panic(err)
	}
	config := Config{}
	err = json.Unmarshal(data,&config)
	if err != nil {
	panic(err)
 }
 return config