package ams

import (
	"fmt"
)

const (
	frameDelimiter          = 0x7E
	bytesAlreadyRead        = 3
	numberOfFrameDelimiters = 2 /* start byte 0x7e + end byte 0x7e*/
)

const (
	Type3 = 0xA
)

var (
	ErrWrongLength           = "Wrong length of buffer, should be > 3, was %d"
	ErrWrongStartByte        = "Expected 0x7E as first byte, was 0x%02x"
	ErrWrongEndByte          = "Expected 0x7E as last byte, was 0x%02x"
	ErrWrongFrameFormatField = "Expected second byte to start with 0xA?, was 0x%02x (first was 0x%02x)"
)

// 2 Byte Frame Format Field
// 16 BITS: "TTTTSLLLLLLLLLLL"
// - T=Type bits: TTTT = 0101 (0xa0) = Type 3
// - S=Segmentation=0 (Segment = 1)
// - L=11 Length Bits
func VerifyStart(buf []byte) error {
	if len(buf) < 3 {
		return fmt.Errorf(ErrWrongLength, len(buf))
	}
	if buf[0] != frameDelimiter {
		return fmt.Errorf(ErrWrongStartByte, buf[0])
	}
	if getFirst4Bits(buf[1], 4) != Type3 {
		return fmt.Errorf(ErrWrongFrameFormatField, buf[1], buf[0])
	}
	return nil
}

func VerifyEnd(buf []byte) error {
	if buf[len(buf)-1] != frameDelimiter {
		return fmt.Errorf(ErrWrongEndByte, buf[len(buf)-1])
	}
	return nil
}

func ReadLenght(buf []byte) (int, error) {
	if err := VerifyStart(buf); err != nil {
		return 0, err
	}
	first3Bits := int(buf[1] & 0x7)
	lenght := first3Bits<<8 + int(buf[2])
	return lenght + numberOfFrameDelimiters - bytesAlreadyRead, nil
}

func ReadMessageType(buf []byte) (MessageType, error) {
	if len(buf) <= 18 {
		return 0, fmt.Errorf(ErrWrongLength, len(buf))
	}
	offset := 17 + 2 + int(buf[18])
	if len(buf) <= offset {
		return 0, fmt.Errorf(ErrWrongLength, len(buf))
	}
	return MessageType(buf[offset+1]), nil
}

func getFirst4Bits(input byte, n uint) byte {
	return input >> 4
}
