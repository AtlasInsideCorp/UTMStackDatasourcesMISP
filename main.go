package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/KbaYero/UTMStackDatasourcesMISP/controller"
	"github.com/KbaYero/UTMStackDatasourcesMISP/database"
	"github.com/KbaYero/UTMStackDatasourcesMISP/models"
	"github.com/KbaYero/UTMStackDatasourcesMISP/utils"
)

const (
	TimeCheck int = 2
)

var mu sync.Mutex

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

	//Get Environment Variables
	apiKey := utils.Getenv("API_KEY")
	apiUrl := utils.Getenv("API_URL")
	//var dataTypes []string
	//dataTypes = append(dataTypes, utils.Getenv("DATA_TYPES"))
	//if strings.Contains(dataTypes[0], ",") {
	//	dataTypes = strings.Split(dataTypes[0], ",")
	//}

	//-----------Init DB for attributes--------------------
	db, err := database.NewDBInstance("attributes", pathDB)
	if err != nil {
		log.Fatal("error creating database: ", err)
	}
	db.Migrate(models.AttributeDB{})
	defer db.Close()

	//--------Process Data from Misp API and sent to correlation-------
	controller.ProcessNewData(db, apiKey, apiUrl, TimeCheck, &mu)

	//----------Control close db----------
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	go func() {
		<-sigCh
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("database closed...")
		os.Exit(0)
	}()
}
