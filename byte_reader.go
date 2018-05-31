package ams

import (
	"fmt"
)

var buf = make([]byte, 1024)

func ByteReader(ch chan byte) chan []byte {
	bytePackages := make(chan []byte)
	go func() {
		readBytes(ch, buf, 0, 3)

		for {
			if err := VerifyStart(buf); err != nil {
				fmt.Println("[ERROR]", err)
				buf[0], buf[1] = buf[1], buf[2]
				buf[2] = <-ch
				continue
			}

			numberOfRemainingBytes, err := ReadLenght(buf)
			if err != nil {
				fmt.Println(err)
			}

			readBytes(ch, buf, 3, numberOfRemainingBytes)
			bytePackage := buf[:3+numberOfRemainingBytes]
			if err := VerifyEnd(bytePackage); err != nil {
				fmt.Println(err)
			}
			bytePackages <- bytePackage

			readBytes(ch, buf, 0, 3)
		}
	}()
	return bytePackages
}

func readBytes(ch chan byte, buf []byte, start int, stop int) {
	for i := start; i < start+stop; i++ {
		buf[i] = <-ch
	}
}
