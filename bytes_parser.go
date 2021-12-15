package ams

import (
	"fmt"
	"strings"
	"time"

	"github.com/karl-gustav/ams-han/crc16"
)

const (
	isVariableLenght = byte(0x09)
	isFixedLenght    = byte(0x06)
	fixedBytesLenght = 4
)

type cursorType struct {
	position int
}

// BytesParser converts the mbus byte array into go structs
func BytesParser(byteList []byte) (v interface{}, err error) {
	if byteList[17] != 0x09 {
		return nil, fmt.Errorf("unknown header type %02x, expecting %02x", byteList[14], 0x09)
	}
	// Don't know what's stored in byteList[23] ¯\_(ツ)_/¯
	messageTimestamp := byteArrayToTime(append(append([]byte{}, byteList[19:23]...), byteList[24:27]...))

	cursor := cursorType{17 + 2 + int(byteList[18])}
	if byteList[cursor.position] != 0x02 {
		return nil, fmt.Errorf("unknown message type %02x, expecting %02x", byteList[cursor.position], 0x02)
	}

	// Checksum everything except delimiters (start & end bytes) and checksum bytes
	packageChecksum := uint16(byteList[len(byteList)-2])<<8 + uint16(byteList[len(byteList)-3])
	calculatedChecksum := crc16.ChecksumCCITT(byteList[1 : len(byteList)-3])
	if calculatedChecksum != packageChecksum {
		return nil, fmt.Errorf("Calculated checksum was %x but the transmitted checksum was %x\n%s", calculatedChecksum, packageChecksum, bytesToHexStrings(byteList[1:len(byteList)-1]))
	}

	messageType := messageTypes(byteList[cursor.position+1])
	cursor.position += 2

	baseItem := BaseItem{
		MeterTime:   messageTimestamp,
		HostTime:    time.Now(),
		MessageType: messageType,
	}

	switch messageType {
	case messageType1:
		v = &MessageType1{
			BaseItem:  baseItem,
			ActPowPos: extractInt(&cursor, byteList),
		}
	case singleFaseMessageType2:
		v = &SingleFaseMessageType2{
			BaseItem:        baseItem,
			ObisListVersion: extractString(&cursor, byteList),
			Gs1:             extractString(&cursor, byteList),
			MeterModel:      extractString(&cursor, byteList),
			ActPowPos:       extractInt(&cursor, byteList),
			ActPowNeg:       extractInt(&cursor, byteList),
			ReactPowPos:     extractInt(&cursor, byteList),
			ReactPowNeg:     extractInt(&cursor, byteList),
			CurrL1:          extractInt(&cursor, byteList),
			VoltL1:          extractInt(&cursor, byteList),
		}
	case threeFaseMessageType2:
		v = &ThreeFaseMessageType2{
			BaseItem:        baseItem,
			ObisListVersion: extractString(&cursor, byteList),
			Gs1:             extractString(&cursor, byteList),
			MeterModel:      extractString(&cursor, byteList),
			ActPowPos:       extractInt(&cursor, byteList),
			ActPowNeg:       extractInt(&cursor, byteList),
			ReactPowPos:     extractInt(&cursor, byteList),
			ReactPowNeg:     extractInt(&cursor, byteList),
			CurrL1:          extractInt(&cursor, byteList),
			CurrL2:          extractInt(&cursor, byteList),
			CurrL3:          extractInt(&cursor, byteList),
			VoltL1:          extractInt(&cursor, byteList),
			VoltL2:          extractInt(&cursor, byteList),
			VoltL3:          extractInt(&cursor, byteList),
		}
	case singleFaseMessageType3:
		v = &SingleFaseMessageType3{
			BaseItem:        baseItem,
			ObisListVersion: extractString(&cursor, byteList),
			Gs1:             extractString(&cursor, byteList),
			MeterModel:      extractString(&cursor, byteList),
			ActPowPos:       extractInt(&cursor, byteList),
			ActPowNeg:       extractInt(&cursor, byteList),
			ReactPowPos:     extractInt(&cursor, byteList),
			ReactPowNeg:     extractInt(&cursor, byteList),
			CurrL1:          extractInt(&cursor, byteList),
			VoltL1:          extractInt(&cursor, byteList),
			DateTime:        extractTime(&cursor, byteList),
			ActEnergyPos:    extractInt(&cursor, byteList),
			ActEnergyNeg:    extractInt(&cursor, byteList),
			ReactEnergyPos:  extractInt(&cursor, byteList),
			ReactEnergyNeg:  extractInt(&cursor, byteList),
		}
	case threeFaseMessageType3:
		v = &ThreeFaseMessageType3{
			BaseItem:        baseItem,
			ObisListVersion: extractString(&cursor, byteList),
			Gs1:             extractString(&cursor, byteList),
			MeterModel:      extractString(&cursor, byteList),
			ActPowPos:       extractInt(&cursor, byteList),
			ActPowNeg:       extractInt(&cursor, byteList),
			ReactPowPos:     extractInt(&cursor, byteList),
			ReactPowNeg:     extractInt(&cursor, byteList),
			CurrL1:          extractInt(&cursor, byteList),
			CurrL2:          extractInt(&cursor, byteList),
			CurrL3:          extractInt(&cursor, byteList),
			VoltL1:          extractInt(&cursor, byteList),
			VoltL2:          extractInt(&cursor, byteList),
			VoltL3:          extractInt(&cursor, byteList),
			DateTime:        extractTime(&cursor, byteList),
			ActEnergyPa:     extractInt(&cursor, byteList),
			ActEnergyMa:     extractInt(&cursor, byteList),
			ActEnergyPr:     extractInt(&cursor, byteList),
			ActEnergyMr:     extractInt(&cursor, byteList),
		}
	}
	return
}

func extractInt(cursor *cursorType, byteList []byte) (result int) {
	if byteList[cursor.position] != isFixedLenght {
		panic(fmt.Sprintf("Expected the first byte to be %02x on mbus int, was %02x", isFixedLenght, byteList[cursor.position]))
	}
	if cursor.position+5+1 > len(byteList) {
		panic("Cursor +6 is bigger than the size of byteList")
	}
	result = bytesToInt(byteList[cursor.position+1 : cursor.position+5])
	cursor.position += 5
	return
}

func extractString(cursor *cursorType, byteList []byte) (result string) {
	if byteList[cursor.position] != isVariableLenght {
		panic(fmt.Sprintf("Expected the first byte to be %02x on mbus string, was %02x", isVariableLenght, byteList[cursor.position]))
	}
	lengthOfMessage := int(byteList[cursor.position+1]) + 2
	if cursor.position+lengthOfMessage+1 > len(byteList) {
		panic(fmt.Sprintf("Cursor +%d is bigger than the size of byteList", lengthOfMessage))
	}
	endOfMessage := cursor.position + lengthOfMessage
	result = string(byteList[cursor.position+2 : endOfMessage])
	cursor.position += lengthOfMessage
	return
}

func extractTime(cursor *cursorType, byteList []byte) (result time.Time) {
	if byteList[cursor.position] != isVariableLenght {
		panic(fmt.Sprintf("Expected the first byte to be %02x on mbus string, was %02x", isVariableLenght, byteList[cursor.position]))
	}
	lengthOfMessage := int(byteList[cursor.position+1]) + 2
	if cursor.position+lengthOfMessage+1 > len(byteList) {
		panic(fmt.Sprintf("Cursor +%d is bigger than the size of byteList", lengthOfMessage))
	}
	endOfMessage := cursor.position + lengthOfMessage
	result = byteArrayToTime(byteList[cursor.position+2 : endOfMessage])
	cursor.position += lengthOfMessage
	return
}

func byteArrayToTime(byteList []byte) time.Time {
	if len(byteList) < 7 {
		panic("Tried to convert byte array less than 8 to time.Time")
	}
	year := int(byteList[0])<<8 + int(byteList[1])
	month := time.Month(byteList[2])
	day := int(byteList[3])
	hour := int(byteList[4])
	min := int(byteList[5])
	sec := int(byteList[6])
	return time.Date(year, month, day, hour, min, sec, 0, time.Local)
}

func bytesToInt(byteList []byte) int {
	if len(byteList) > 4 {
		panic("Tried to convert byte array greater than 4 to int")
	}
	data := int(0)
	for _, b := range byteList {
		data = (data << 8) | int(b)
	}
	return data
}

func bytesToHexStrings(bytes []byte) string {
	var output []string
	for _, b := range bytes {
		output = append(output, fmt.Sprintf("0x%02x", b))
	}
	return strings.Join(output, " ")
}
