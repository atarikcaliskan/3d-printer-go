package printer

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jacobsa/go-serial/serial"
)

type Printer struct {
	Port      io.ReadWriteCloser
	LogWriter *bufio.Writer
	LogFile   *os.File
}

func New(portName string, baudRate int, logFileName string) (*Printer, error) {

	options := serial.OpenOptions{
		PortName:        portName,
		BaudRate:        uint(baudRate),
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	port, err := serial.Open(options)
	if err != nil {
		return nil, err
	}

	logFile, err := os.Create(logFileName)
	if err != nil {
		return nil, err
	}

	logWriter := bufio.NewWriter(logFile)

	return &Printer{
		Port:      port,
		LogFile:   logFile,
		LogWriter: logWriter,
	}, nil
}

func (p *Printer) Read() {
	scanner := bufio.NewScanner(p.Port)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Printer Message:", line)

		_, err := p.LogWriter.WriteString("Printer Message: " + line + "\n")
		if err != nil {
			log.Fatal(err)
		}
		p.LogWriter.Flush()
	}
}

func (p *Printer) CheckTemp() {
	for {
		command := "M105\n"
		_, err := p.Port.Write([]byte(command))
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second)
	}
}

func (p *Printer) ExecuteInputCmd() {
	for {
		fmt.Print("Enter custom command (or press Enter to skip): ")
		userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		userInput = userInput[:len(userInput)-1]

		if userInput != "" {
			fmt.Println("Sending User Command:", userInput)
			_, err := p.Port.Write([]byte(userInput + "\n"))
			if err != nil {
				log.Fatal(err)
			}
		}

		time.Sleep(time.Second)
	}
}

func (p *Printer) Close() {
	if p.Port != nil {
		p.Port.Close()
	}
	p.LogWriter.Flush()
	if p.LogFile != nil {
		p.LogFile.Close()
	}
}

func (p *Printer) CloseLogFile() {
	p.LogFile.Close()
}
