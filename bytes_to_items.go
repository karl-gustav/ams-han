package ams

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	isVariableLenght = byte(0x09)
	isFixedLenght    = byte(0x06)
	fixedBytesLenght = 4
)

func bytesToItem(bytes []byte) (v interface{}, err error) {
	if bytes[17] != 0x09 {
		return nil, fmt.Errorf("Unknown header type: %02x\n\n\n", bytes[14])
	}
	year := int(byteArrayToInt64(bytes[19:21]))
	month := int(bytes[21])
	day := int(bytes[22])
	hour := int(bytes[24])
	min := int(bytes[25])
	sec := int(bytes[26])
	messageTimestamp := time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local)

	offset := 17 + 2 + int(bytes[18])
	if bytes[offset] != 0x02 {
		return nil, fmt.Errorf("Unknown message type: %02x\n\n\n", bytes[offset])
	}

	messageType := messageTypes(bytes[offset+1])
	offset += 2

	now := time.Now()

	baseItem := BaseItem{
		MeterTime:   messageTimestamp,
		HostTime:    now,
		MessageType: messageType,
	}

	switch messageType {
	case messageType1:
		v = &MessageType1{
			BaseItem: baseItem,
		}
	case twoFasesMessageType2:
		v = &TwoFasesMessageType2{
			BaseItem: baseItem,
		}
	case threeFasesMessageType2:
		v = &ThreeFasesMessageType2{
			BaseItem: baseItem,
		}
	case twoFasesMessageType3:
		v = &TwoFasesMessageType3{
			BaseItem: baseItem,
		}
	case threeFasesMessageType3:
		v = &ThreeFasesMessageType3{
			BaseItem: baseItem,
		}
	}

	pointer := reflect.ValueOf(v)
	if pointer.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("You need to provide a pointer to MashalItems")
	}
	val := pointer.Elem()
	eTyp := reflect.TypeOf(v).Elem()
	for i := 0; i < eTyp.NumField(); i++ {
		field := eTyp.Field(i)
		var endOfMessage int
		switch field.Type.Kind() {
		case reflect.Int:
			if isFixedLenght != bytes[offset] {
				return nil, wrongOffsetMarkerError(bytes[offset:], isFixedLenght, field.Name)
			}
			endOfMessage = offset + fixedBytesLenght + 1
			offset += 1
			val.Field(i).SetInt(byteArrayToInt64(bytes[offset:endOfMessage]))
		case reflect.String:
			if isVariableLenght != bytes[offset] {
				return nil, wrongOffsetMarkerError(bytes[offset:], isVariableLenght, field.Name)
			}
			lengthOfMessage := int(bytes[offset+1]) + 2
			endOfMessage = offset + lengthOfMessage
			offset += 2
			val.Field(i).SetString(string(bytes[offset:endOfMessage]))
		case reflect.Struct:
			valField := val.Field(i)
			switch valField.Type() {
			case reflect.TypeOf((*BaseItem)(nil)).Elem():
				continue // Skip BaseItem embedded struct, don't move offset
			case reflect.TypeOf((*time.Time)(nil)).Elem():
				if isVariableLenght != bytes[offset] {
					return nil, wrongOffsetMarkerError(bytes[offset:], isVariableLenght, field.Name)
				}
				lengthOfMessage := int(bytes[offset+1]) + 2
				endOfMessage = offset + lengthOfMessage
				offset += 2
				valField.Set(reflect.ValueOf(byteArrayToTime(bytes[offset:endOfMessage])))
			default:
				throwUnknownFieldTypeError(field.Name, baseItem.MessageType, "struct")
			}
		default:
			throwUnknownFieldTypeError(field.Name, baseItem.MessageType, "built-in")
		}
		offset = endOfMessage
	}
	return
}

func byteArrayToInt64(bytes []byte) int64 {
	if len(bytes) > 8 {
		panic("Tried to convert byte array greater than 4 to int")
	}
	data := int64(0)
	for _, b := range bytes {
		data = (data << 8) | int64(b)
	}
	return data
}

func byteArrayToTime(bytes []byte) time.Time {
	if len(bytes) < 7 {
		panic("Tried to convert byte array less than 8 to time.Time")
	}
	year := int(bytes[0])<<8 + int(bytes[1])
	month := time.Month(bytes[2])
	day := int(bytes[3])
	hour := int(bytes[4])
	min := int(bytes[5])
	sec := int(bytes[6])
	return time.Date(year, month, day, hour, min, sec, 0, time.Local)
}

func throwUnknownFieldTypeError(fieldName string, messageType messageTypes, fieldType string) {
	panic(fmt.Sprintf(
		"Don't recognice the type of field %s (message type %v, field type %s), will only recognice int, string and time.Time!",
		fieldName,
		messageType,
		fieldType,
	))
}

func wrongOffsetMarkerError(restOfBytes []byte, isFixedLenght byte, fieldName string) error {
	return fmt.Errorf(
		"Offset wasn't what it was supposed to be, was %02x, expected %02x for %s\nRest of bytes: [%s]\n",
		restOfBytes[0],
		isFixedLenght,
		fieldName,
		strings.Join(bytesToHexStrings(restOfBytes), ", "),
	)
}

func bytesToHexStrings(bytes []byte) (output []string) {
	for _, b := range bytes {
		output = append(output, fmt.Sprintf("0x%02x", b))
	}
	return
}
