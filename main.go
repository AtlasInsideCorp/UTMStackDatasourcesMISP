package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/KbaYero/UTMStackDatasourcesMISP/config"
	"github.com/KbaYero/UTMStackDatasourcesMISP/controller"
	"github.com/KbaYero/UTMStackDatasourcesMISP/fileserver"
)

const (
	TimeCheck int = 12
)

func main() {
	cnf := config.GetConfig()
	go controller.ProcessData(cnf, TimeCheck)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	go func() {
		<-sigCh
		fmt.Println("finished...")
		os.Exit(0)
	}()

	fileserver.ShowEvents(cnf)
}
