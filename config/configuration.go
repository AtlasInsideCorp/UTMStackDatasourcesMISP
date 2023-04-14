package config

import (
	"os"
	"strconv"
	"sync"

	"github.com/KbaYero/UTMStackDatasourcesMISP/utils"
)

type Configuration struct {
	ApiKey     string
	Instance   string
	EventsPath string
	ServerPort string
	TimeCheck  int
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
		configurationInstance.Instance = utils.Getenv("MISP_INSTANCE")
		configurationInstance.EventsPath = utils.Getenv("EVENTS_PATH")
		configurationInstance.ServerPort = utils.Getenv("SERVER_PORT")
		checkString := utils.Getenv("TIME_CHECK")
		configurationInstance.TimeCheck, _ = strconv.Atoi(checkString)
	})
	return configurationInstance
}
