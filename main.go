package main

import (
	"log"
	"os"

	"github.com/KbaYero/UTMStackDatasourcesMISP/database"
	"github.com/KbaYero/UTMStackDatasourcesMISP/models"
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

	//Get Environment Variables
	//apiKey := utils.Getenv("API_KEY")
	//apiUrl := utils.Getenv("API_URL")
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
}
