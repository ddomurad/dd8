package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tarm/serial"
)

// this simple program reads a intel hex file
// and sends the bytes to a serial port in hex ascii format
// every line is sent with format addr_h, addr_l, data, data, ...., '$'
// every character sent is echoed back by the serial port, and should be validated if the data was sent correctly
// the program uses flag package to parse the command line arguments
// program arguments: -port, -file, -entry , where the entry is the entry addres of the uploaded code

func loadHexFile(file string) (string, error) {
	data, err := os.ReadFile(file)
	return string(data), err
}

func parseHexLine(line string) ([]byte, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil
	}

	if line == ":00000001FF" {
		return nil, nil
	}

	if len(line) < 11 {
		return nil, fmt.Errorf("invalid hex line, too short: %s", line)
	}

	if line[0] != ':' {
		return nil, errors.New("invalid hex line, missing ':'")
	}

	addr := line[3:7]
	data := line[9 : len(line)-2]
	out := addr + data + "$"
	out = strings.ToLower(out)
	return []byte(out), nil
}

func sendSerialData(port string, data []byte) error {
	c := &serial.Config{Name: port, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		return err
	}

	defer s.Close()

	echoReadBuffer := make([]byte, 1)
	fmt.Println()
	for _, r := range data {
		fmt.Printf("%c", r)
		_, err = s.Write([]byte{r})
		if err != nil {
			return err
		}
		// read the echo
		_, err = s.Read(echoReadBuffer)
		if err != nil {
			return err
		}

		if echoReadBuffer[0] != r {
			fmt.Println()
			fmt.Printf("echo mismatch: %c != %c\n", r, echoReadBuffer[0])
			return errors.New("echo mismatch")
		}
	}

	return nil
}

func main() {
	serialPort := flag.String("port", "", "serial port to use")
	hexFile := flag.String("file", "", "hex file to upload")
	entry := flag.String("entry", "", "entry address of the code")

	flag.Parse()

	if *serialPort == "" || *hexFile == "" || *entry == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	hexFileData, err := loadHexFile(*hexFile)
	if err != nil {
		panic(err)
	}

	hexLines := strings.Split(hexFileData, "\n")
	sendBuffer := make([]byte, 0)

	for _, line := range hexLines {
		data, err := parseHexLine(line)
		if err != nil {
			panic(err)
		}
		sendBuffer = append(sendBuffer, data...)
	}

	entryAddrBytes := []byte(*entry)
	sendBuffer = append(sendBuffer, entryAddrBytes...)
	sendBuffer = append(sendBuffer, '!')

	err = sendSerialData(*serialPort, sendBuffer)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
