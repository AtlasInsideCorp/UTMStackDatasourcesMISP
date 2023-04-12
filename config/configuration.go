package config

import (
	"os"
	"sync"

	"github.com/KbaYero/UTMStackDatasourcesMISP/utils"
)

type Configuration struct {
	ApiKey     string
	ApiUrl     string
	EventsPath string
}

var (
	configurationInstance Configuration
	configurationOnce     sync.Once
)

// GetConfig return a configuration object
func GetConfig() Configuration {
	configurationOnce.Do(func() {
		if len(os.Args) > 1 {
			arg := os.Args[1]
			switch arg {
			case "run":
				utils.LoadEnv()
			}
		}

		configurationInstance.ApiKey = utils.Getenv("API_KEY")
		configurationInstance.ApiUrl = utils.Getenv("API_URL")
		configurationInstance.EventsPath = utils.Getenv("EVENTS_PATH")
	})

	return configurationInstance
}
