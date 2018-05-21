package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/goburrow/serial"
)

var (
	address  string
	baudrate int
	databits int
	stopbits int
	parity   string

	message string
)

func main() {
	flag.StringVar(&address, "a", "/dev/ttyUSB2", "address")
	flag.IntVar(&baudrate, "b", 2400, "baud rate")
	flag.IntVar(&databits, "d", 8, "data bits")
	flag.IntVar(&stopbits, "s", 1, "stop bits")
	flag.StringVar(&parity, "p", "N", "parity (N/E/O)")
	flag.StringVar(&message, "m", "serial", "message")
	flag.Parse()

	config := serial.Config{
		Address:  address,
		BaudRate: baudrate,
		DataBits: databits,
		StopBits: stopbits,
		Parity:   parity,
		Timeout:  2200 * time.Millisecond,
	}
	log.Printf("connecting %+v", config)
	port, err := serial.Open(&config)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected")
	defer func() {
		err := port.Close()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("closed")
	}()

	b := make([]byte, 8)
	for {
		ts := time.Now().Format("2006-01-02T15:04:05.999")
		n, err := port.Read(b)
		if err == io.EOF {
			log.Fatalln("Reached end of stream")
			break
		} else if err != nil {
			log.Println("[ERROR]:", err)
			continue
		}
		if n == 0 {
			panic("n was 0")
		}
		for _, v := range b[:n] {
			fmt.Printf("%s\t %08b (%02x)", ts, v, v)

			if v == byte(0x7e) {
				fmt.Print("  <-- was 0x7e (126)")
			}
			fmt.Println("")
		}
	}
}
