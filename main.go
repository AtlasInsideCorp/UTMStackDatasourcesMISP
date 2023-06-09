package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/KbaYero/UTMStackDatasourcesMISP/config"
	"github.com/KbaYero/UTMStackDatasourcesMISP/controller"
	"github.com/KbaYero/UTMStackDatasourcesMISP/serv"
)

func main() {
	cnf := config.GetConfig()
	go controller.ProcessData(cnf)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	go func() {
		<-sigCh
		fmt.Println("finished...")
		os.Exit(0)
	}()

	serv.ShowEvents(cnf)
}
