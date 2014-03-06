package main

import (
    "fmt"
	"encoding/json"
    "os"
)

type Settings interface {
	Load(filename string)
	PortNo() string
}

type SettingsImpl struct {
	Port string
}

func (settings *SettingsImpl) Load(filename string)  {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	decoder.Decode(&settings)
}

func (settings *SettingsImpl) PortNo() string {
	return settings.Port;
}
