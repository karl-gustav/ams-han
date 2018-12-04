package ams

import (
	"fmt"
)

var buf = make([]byte, 1024)

var CHANNEL_IS_CLOSED_ERROR = fmt.Errorf("Channel is closed!")

func ByteReader(ch chan byte) (next func() ([]byte, error)) {
	var channelIsOpen bool
	return func() ([]byte, error) {
		readBytes(ch, buf, 0, 3)
		if err := VerifyStart(buf); err != nil {
			buf[0], buf[1] = buf[1], buf[2]
			buf[2], channelIsOpen = <-ch
			if !channelIsOpen {
				return nil, CHANNEL_IS_CLOSED_ERROR
			}
			return nil, err
		}

		numberOfRemainingBytes, err := ReadLenght(buf)
		if err != nil {
			return nil, err
		}

		readBytes(ch, buf, 3, numberOfRemainingBytes)
		bytePackage := buf[:3+numberOfRemainingBytes]
		if err := VerifyEnd(bytePackage); err != nil {
			return nil, err
		}
		return bytePackage, err
	}
}

func readBytes(ch chan byte, buf []byte, start int, stop int) {
	for i := start; i < start+stop; i++ {
		buf[i] = <-ch
	}
}
