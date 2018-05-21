package decode_mbus

import (
	"fmt"
	"reflect"
	"time"
)

const tagName = "bytes"

const (
	isVariableLenght = byte(0x09)
	isFixedLenght    = byte(0x06)
	fixedBytesLenght = 4
)

func MarshalItems(bytes []byte, v interface{}) (err error) {
	pointer := reflect.ValueOf(v)
	if pointer.Kind() != reflect.Ptr {
		return fmt.Errorf("You need to provide a pointer to MashalItems")
	}

	val := pointer.Elem()
	eTyp := reflect.TypeOf(v).Elem()
	offset := 0
	for i := 0; i < eTyp.NumField(); i++ {
		field := eTyp.Field(i)
		var endOfMessage int
		switch field.Type.Kind() {
		case reflect.Int:
			if isFixedLenght != bytes[offset] {
				return fmt.Errorf("Offset wasn't what it was supposed to be, was %02x, expected %02x for %s\n", bytes[offset], isFixedLenght, field.Name)
			}
			endOfMessage = offset + fixedBytesLenght + 1
			offset += 1
			val.Field(i).SetInt(byteArrayToInt64(bytes[offset:endOfMessage]))
		case reflect.String:
			if isVariableLenght != bytes[offset] {
				return fmt.Errorf("Offset wasn't what it was supposed to be, was %02x, expected %02x for %s\n", bytes[offset], isVariableLenght, field.Name)
			}
			lengthOfMessage := int(bytes[offset+1]) + 2
			endOfMessage = offset + lengthOfMessage
			offset += 2
			val.Field(i).SetString(string(bytes[offset:endOfMessage]))
		case reflect.Struct:
			valField := val.Field(0)
			switch valField.Type() {
			case reflect.TypeOf(time.Time{}):
				if isVariableLenght != bytes[offset] {
					return fmt.Errorf("Offset wasn't what it was supposed to be, was %02x, expected %02x for %s\n", bytes[offset], isVariableLenght, field.Name)
				}
				lengthOfMessage := int(bytes[offset+1]) + 2
				endOfMessage = offset + lengthOfMessage
				offset += 2
				valField.Set(reflect.ValueOf(byteArrayToTime(bytes[offset:endOfMessage])))
			}
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
