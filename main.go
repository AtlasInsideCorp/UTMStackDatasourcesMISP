package main

import (
	"os"

	"github.com/KbaYero/UTMStackDatasourcesMISP/utils"
)

func main() {
	//Load environment variables
	var pathDB string
	if len(os.Args) > 1 {
		arg := os.Args[1]
		switch arg {
		case "run":
			utils.LoadEnv()
			pathDB = "db"
		}
	} else {
		pathDB = "/usr/src/app/db"
	}

	if !utils.CheckIfExistPath(pathDB) {
		os.Mkdir(pathDB, 0700)
	}
}
