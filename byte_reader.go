package ams

import (
)

var buf = make([]byte, 1024)

func ByteReader(ch chan byte) (chan []byte, chan error) {
	bytePackages := make(chan []byte)
	errors := make(chan error)
	channelIsOpen := false
	go func() {
		readBytes(ch, buf, 0, 3)

		for {
			if err := VerifyStart(buf); err != nil {
				buf[0], buf[1] = buf[1], buf[2]
				buf[2], channelIsOpen = <-ch
				if !channelIsOpen {
					break
				}
				errors <- err
				continue
			}

			numberOfRemainingBytes, err := ReadLenght(buf)
			if err != nil {
				errors <- err
				continue
			}

			readBytes(ch, buf, 3, numberOfRemainingBytes)
			bytePackage := buf[:3+numberOfRemainingBytes]
			if err := VerifyEnd(bytePackage); err != nil {
				errors <- err
				continue
			}
			bytePackages <- bytePackage

			readBytes(ch, buf, 0, 3)
		}
		close(bytePackages)
		close(errors)
	}()
	return bytePackages, errors
}

func readBytes(ch chan byte, buf []byte, start int, stop int) {
	for i := start; i < start+stop; i++ {
		buf[i] = <-ch
	}
}
