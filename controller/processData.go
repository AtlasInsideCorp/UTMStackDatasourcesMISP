package controller

import (
	"log"
	"sync"
	"time"

	"github.com/KbaYero/UTMStackDatasourcesMISP/database"
	"github.com/KbaYero/UTMStackDatasourcesMISP/models"
)

// ProcessNewData requests the data from the MISP API, to later process it and save it in the database
// If the data exists, check if it can be updated
func ProcessNewData(db database.ISqlite, apiKey, apiUrl string, timeCheck int, mu *sync.Mutex) {
	go func() {
	ExternalLoop:
		for {
			var lisAttrJson []models.AttributesJson
			var listAttrDB []models.AttributeDB
			err := getData(apiKey, apiUrl+"/attributes", &lisAttrJson)
			if err != nil {
				log.Println("error getting data: ", err)
				time.Sleep(time.Minute * time.Duration(timeCheck))
				continue
			}
			for _, attJson := range lisAttrJson {
				attDB := AttrFromJsonToDB(attJson)
				listAttrDB = append(listAttrDB, attDB)
			}
			mu.Lock()
			for _, attDB := range listAttrDB {
				if db.CheckIfExist(&attDB, attDB.ItemID) {
					err := db.Update(&attDB, attDB.ItemID)
					if err != nil {
						mu.Unlock()
						log.Println("error updating item in the database: ", err)
						time.Sleep(time.Minute * time.Duration(timeCheck))
						continue ExternalLoop
					}
				} else {
					attDB.SendToCorrelation = false
					err = db.AddNew(&attDB)
					if err != nil {
						mu.Unlock()
						log.Println("error adding new item in the database: ", err)
						time.Sleep(time.Minute * time.Duration(timeCheck))
						continue ExternalLoop
					}
				}
			}
			mu.Unlock()
			log.Println("Data updated successfully")
			time.Sleep(time.Minute * time.Duration(timeCheck))
		}
	}()
}
