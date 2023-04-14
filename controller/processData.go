package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/KbaYero/UTMStackDatasourcesMISP/config"
	"github.com/KbaYero/UTMStackDatasourcesMISP/models"
	"github.com/KbaYero/UTMStackDatasourcesMISP/utils"
)

// ProcessAttributes processes data of type attribute
func ProcessData(cnf config.Configuration) {
	for {
		if utils.CheckIfExistPath(cnf.EventsPath) {
			os.RemoveAll(cnf.EventsPath)
		}
		os.Mkdir(cnf.EventsPath, 0700)

		var wg sync.WaitGroup
		concurrency := 10
		semaphore := make(chan bool, concurrency)
		i := -1
		stop := false

		for !stop {
			semaphore <- true
			wg.Add(1)
			i++

			go func(i int) {
				defer func() {
					wg.Done()
					<-semaphore
				}()

				var listEvents models.ResponseBodyGetEvents
				reqBody := models.RequestEventsRestSearch{
					Page:         i,
					Limit:        100,
					ReturnFormat: "json",
				}
				err := getData(cnf.ApiKey, "https://"+cnf.Instance+"/events/restSearch", &listEvents, reqBody)
				if err != nil {
					log.Printf("error getting data for page %d: %v", i, err)
					return
				}
				if len(listEvents.Response) == 0 {
					stop = true
				} else {
					var eventCleaned []models.CleanedEventsBody
					for _, event := range listEvents.Response {
						eventCleaned = append(eventCleaned, cleanData(event))
					}
					file, err := json.MarshalIndent(eventCleaned, "", " ")
					if err != nil {
						log.Printf("error saving data to json file for page %d: %v", i, err)
						return
					}
					currentTime := fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))
					filename := filepath.Join(cnf.EventsPath, "misp-instance-"+currentTime+".json")
					err = os.WriteFile(filename, file, 0644)
					if err != nil {
						log.Printf("error creating json file for page %d: %v", i, err)
						return
					}
				}
			}(i)
			if stop {
				break
			}
		}
		wg.Wait()
		log.Println("Event update process finished")
		time.Sleep(time.Duration(cnf.TimeCheck) * time.Hour)
	}
}
