package config

import (
	"encoding/json"
	"io/ioutil"
)

var Configuration Config

type Config struct {
	MaxOriginalHeight int `json:"maxOriginalHeight"`
	MaxOriginalWidth  int `json:"maxOriginalWidth"`
	MaxResizedHeight  int `json:"maxResizedHeight"`
	MaxResizedWidth   int `json:"maxResizedWidth"`
}

func LoadConfig() {
	data, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &Configuration)
	if err != nil {
		panic(err)
	}
}
