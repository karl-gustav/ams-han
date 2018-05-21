package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"bitbucket.org/karlgustav/ams-han-mbus/decode_mbus"

	"github.com/goburrow/serial"
)

var (
	address  string
	baudrate int
	databits int
	stopbits int
	parity   string
	verbose  bool
)

const (
	messageType1 = 1
	messageType2 = 13
	messageType3 = 18
)

var debug = os.Getenv("DEBUG") != ""

func main() {
	flag.StringVar(&address, "a", "/dev/ttyUSB2", "address")
	flag.IntVar(&baudrate, "b", 2400, "baud rate")
	flag.IntVar(&databits, "d", 8, "data bits")
	flag.IntVar(&stopbits, "s", 1, "stop bits")
	flag.StringVar(&parity, "p", "E", "parity (N/E/O)")
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.Parse()

	config := serial.Config{
		Address:  address,
		BaudRate: baudrate,
		DataBits: databits,
		StopBits: stopbits,
		Parity:   parity,
		Timeout:  60 * time.Second,
	}
	if debug {
		log.Printf("connecting %+v", config)
	}
	port, err := serial.Open(&config)
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		log.Println("connected")
	}
	defer func() {
		err := port.Close()
		log.Println("Closed connection!")
		if err != nil {
			log.Fatal(err)
		}
	}()

	serialChannel := make(chan byte)

	go readReader(serialChannel)

	for {
		var buf [8]byte
		n, err := port.Read(buf[:])
		if err == io.EOF {
			log.Fatalln("Reached end of stream")
			return
		} else if err != nil {
			log.Println("[ERROR]:", err)
			return
		}
		for i := 0; i < n; i++ {
			serialChannel <- buf[i]
		}
	}
}

func readReader(ch chan byte) {
	startBuf := make([]byte, 3)
	readFull(ch, startBuf)

	for {
		lenght, findStartError := decode_mbus.FindStart(startBuf)
		if findStartError != nil {
			fmt.Println(findStartError)
			startBuf[0], startBuf[1] = startBuf[1], startBuf[2]
			startBuf[2] = <-ch
			if verbose {
				fmt.Printf("Read byte %02x %08b\n", startBuf[2], startBuf[2])
			}
			continue
		}
		if verbose {
			fmt.Printf("0 byte was: %02[1]x (%08[1]b)\n", startBuf[0])
			fmt.Printf("1 byte was: %02[1]x (%08[1]b)\n", startBuf[1])
			fmt.Printf("2 byte was: %02[1]x (%08[1]b) %[1]d\n", startBuf[2])
			fmt.Println("lenght:", lenght)
		}
		buf := make([]byte, lenght)
		readFull(ch, buf)
		ts := time.Now().Format("2006-01-02T15:04:05.999")
		if verbose {
			fmt.Println(ts, "RestBuffer:")
			for i, v := range buf {
				fmt.Printf("%d\t %08[2]b (%02[2]x)", i, v)
				fmt.Println("")
			}
		}

		buf = append(startBuf, buf...)
		if buf[17] == 0x09 {
			var messageType byte
			year := byteArrayToInt(buf[19:21])
			month := int(buf[21])
			day := int(buf[22])
			hour := int(buf[24])
			min := int(buf[25])
			sec := int(buf[26])
			messageTimestamp := time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local)

			offset := 17 + 2 + buf[18]
			if buf[offset] == 0x02 {
				messageType = buf[offset+1]
			}
			offset += 2
			if verbose {
				fmt.Println("message time stamp:", messageTimestamp)
				fmt.Println("Number of items:", messageType)
			}
			if messageType == 1 {
				msg := decode_mbus.Items1{}
				err := decode_mbus.MarshalItems(buf[offset:], &msg)
				printIfNotError(msg, err)
			} else if messageType == 9 {
				msg := decode_mbus.Items9{}
				err := decode_mbus.MarshalItems(buf[offset:], &msg)
				printIfNotError(msg, err)
			} else if messageType == 13 {
				msg := decode_mbus.Items13{}
				err := decode_mbus.MarshalItems(buf[offset:], &msg)
				printIfNotError(msg, err)
			} else if messageType == 14 {
				msg := decode_mbus.Items14{}
				err := decode_mbus.MarshalItems(buf[offset:], &msg)
				printIfNotError(msg, err)
			} else if messageType == 18 {
				msg := decode_mbus.Items18{}
				err := decode_mbus.MarshalItems(buf[offset:], &msg)
				printIfNotError(msg, err)
			}
		} else {
			fmt.Printf("[ERROR] Unknown message type: %02x\n", buf[14])
		}
		readFull(ch, startBuf)
	}
}

func readFull(ch chan byte, buf []byte) {
	for i, _ := range buf {
		buf[i] = <-ch
	}
}

func byteArrayToInt(bytes []byte) int {
	if len(bytes) > 4 {
		panic("Tried to convert byte array greater than 4 to int")
	}
	data := int(0)
	for _, b := range bytes {
		data = (data << 8) | int(b)
	}
	return data
}
func printIfNotError(msg interface{}, err error) {
	if err != nil {
		fmt.Println("[ERROR]", err)
		err = nil
	} else {
		jsonString, _ := json.Marshal(msg)
		fmt.Printf("%s\n", jsonString)
	}
}
