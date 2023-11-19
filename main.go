package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/tarikcaliskan/3d-printer-go/pkg/printer"
)

func main() {
	var portName string
	var baudRate int
	var logFileName string

	flag.StringVar(&portName, "port", "/dev/tty.usbserial-1140", "The serial port name")
	flag.IntVar(&baudRate, "baud-rate", 250000, "The baud rate for serial communication")
	flag.StringVar(&logFileName, "log-file", "printer_log.txt", "The log file name")
	flag.Parse()

	printer, err := printer.New(portName, baudRate, logFileName)
	if err != nil {
		fmt.Println("Error initializing printer:", err)
		return
	}
	defer printer.Close()

	go printer.Read()
	go printer.CheckTemp()
	go printer.ExecuteInputCmd()

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

	printer.CloseLogFile()
	os.Exit(0)
}
