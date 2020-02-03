package cmd

import (
	"encoding/json"
	"os"
	"sync"
)

type Config struct {
	DbUrl string `json:"db_url"`
}

type ConfigFlags struct {
	ConfigFile string `flag:"config" required:"true" default:"config.json" usage:"Path to the config file"`
}

var config *Config
var configOnce sync.Once

func readConfig() {
	var configFlags = ConfigFlags{}
	ReadFlags(&configFlags)
	file, err := os.Open(configFlags.ConfigFile)
	if err != nil {
		panic(err)
	}
	config = &Config{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	configOnce.Do(func() {
		readConfig()
	})
	return config
}
