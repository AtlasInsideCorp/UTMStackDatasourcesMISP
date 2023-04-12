package main

import (
	"os"

	"github.com/KbaYero/UTMStackDatasourcesMISP/config"
	"github.com/KbaYero/UTMStackDatasourcesMISP/utils"
)

const (
	TimeCheck int = 2
)

func main() {
	cnf := config.GetConfig()
	if !utils.CheckIfExistPath(cnf.EventsPath) {
		os.Mkdir(cnf.EventsPath, 0700)
	}

}
