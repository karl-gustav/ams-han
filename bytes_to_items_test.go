package ams

import (
	"testing"
	"time"

	"github.com/karl-gustav/ams-han/crc16"
)

const layout = "2006-01-02T15:04:05"

func TestItem1(t *testing.T) {
	m := []byte{0x7e, 0xa0, 0x27, 0x01, 0x02, 0x01, 0x10, 0x5a, 0x87, 0xe6, 0xe7, 0x00, 0x0f, 0x40, 0x00, 0x00, 0x00, 0x09, 0x0c, 0x07, 0xe2, 0x05, 0x1b, 0x07, 0x13, 0x25, 0x26, 0xff, 0x80, 0x00, 0x00, 0x02, 0x01, 0x06, 0x00, 0x00, 0x03, 0x52, 0x81, 0x3d, 0x7e}
	message, err := BytesParser(m)
	if err != nil {
		t.Errorf("Didn't expect error... %v", err)
		return
	}
	messageType := messageTypes(1)
	meterTime, _ := time.Parse(layout, "2018-05-27T19:37:38")
	actPowPos := 850

	switch v := message.(type) {
	case *MessageType1:
		if v.MessageType != messageType {
			t.Errorf("Unexpected value in MessageType, was %v, expected %v", v.MessageType, messageType)
		}
		if v.MeterTime.Format(layout) != meterTime.Format(layout) {
			t.Errorf("Unexpected value in MeterTime, was %v, expected %v", v.MeterTime.Format(layout), meterTime.Format(layout))
		}
		if v.HostTime == (time.Time{}) {
			t.Errorf("Unexpected value in HostTime, was %v", v.HostTime)
		}
		if v.ActPowPos != actPowPos {
			t.Errorf("Unexpected value in ActPowPos, was %v, expected %v", v.ActPowPos, actPowPos)
		}
	default:
		t.Errorf("Only expected ThreeFasesMessageType2, got %T", message)
	}
}

func TestItem13(t *testing.T) {
	m := []byte{0x7e, 0xa0, 0x78, 0x01, 0x02, 0x01, 0x10, 0xc4, 0x98, 0xe6, 0xe7, 0x00, 0x0f, 0x40, 0x00, 0x00, 0x00, 0x09, 0x0c, 0x07, 0xe2, 0x05, 0x19, 0x05, 0x0c, 0x2b, 0x0a, 0xff, 0x80, 0x00, 0x00, 0x02, 0x0d, 0x09, 0x07, 0x4b, 0x46, 0x4d, 0x5f, 0x30, 0x30, 0x31, 0x09, 0x10, 0x36, 0x39, 0x37, 0x30, 0x36, 0x33, 0x31, 0x34, 0x30, 0x31, 0x34, 0x36, 0x38, 0x38, 0x30, 0x33, 0x09, 0x07, 0x4d, 0x41, 0x33, 0x30, 0x34, 0x48, 0x34, 0x06, 0x00, 0x00, 0x02, 0xa0, 0x06, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x01, 0xb8, 0x06, 0x00, 0x00, 0x08, 0x5e, 0x06, 0x00, 0x00, 0x06, 0x1a, 0x06, 0x00, 0x00, 0x01, 0x7b, 0x06, 0x00, 0x00, 0x09, 0x56, 0x06, 0x00, 0x00, 0x09, 0x66, 0x06, 0x00, 0x00, 0x09, 0x5a, 0x36, 0x45, 0x7e}
	message, err := BytesParser(m)
	if err != nil {
		t.Errorf("Didn't expect error... %v", err)
		return
	}

	messageType := messageTypes(13)
	meterTime, _ := time.Parse(layout, "2018-05-25T12:43:10")
	obisListVersion := "KFM_001"
	gs1 := "6970631401468803"
	meterModel := "MA304H4"
	actPowPos := 672
	actPowNeg := 0
	reactPowPos := 0
	reactPowNeg := 440
	currL1 := 2142
	currL2 := 1562
	currL3 := 379
	voltL1 := 2390
	voltL2 := 2406
	voltL3 := 2394

	switch v := message.(type) {
	case *ThreeFaseMessageType2:
		if v.MessageType != messageType {
			t.Errorf("Unexpected value in MessageType, was %v, expected %v", v.MessageType, messageType)
		}
		if v.MeterTime.Format(layout) != meterTime.Format(layout) {
			t.Errorf("Unexpected value in MeterTime, was %v, expected %v", v.MeterTime.Format(layout), meterTime.Format(layout))
		}
		if v.HostTime == (time.Time{}) {
			t.Errorf("Unexpected value in HostTime, was %v", v.HostTime)
		}
		if v.ObisListVersion != obisListVersion {
			t.Errorf("Unexpected value in ObisListVersion, was %v, expected %v", v.ObisListVersion, obisListVersion)
		}
		if v.Gs1 != gs1 {
			t.Errorf("Unexpected value in Gs1, was %v, expected %v", v.Gs1, gs1)
		}
		if v.MeterModel != meterModel {
			t.Errorf("Unexpected value in MeterModel, was %v, expected %v", v.MeterModel, meterModel)
		}
		if v.ActPowPos != actPowPos {
			t.Errorf("Unexpected value in ActPowPos, was %v, expected %v", v.ActPowPos, actPowPos)
		}
		if v.ActPowNeg != actPowNeg {
			t.Errorf("Unexpected value in ActPowNeg, was %v, expected %v", v.ActPowNeg, actPowNeg)
		}
		if v.ReactPowPos != reactPowPos {
			t.Errorf("Unexpected value in ReactPowPos, was %v, expected %v", v.ReactPowPos, reactPowPos)
		}
		if v.ReactPowNeg != reactPowNeg {
			t.Errorf("Unexpected value in ReactPowNeg, was %v, expected %v", v.ReactPowNeg, reactPowNeg)
		}
		if v.CurrL1 != currL1 {
			t.Errorf("Unexpected value in CurrL1, was %v, expected %v", v.CurrL1, currL1)
		}
		if v.CurrL2 != currL2 {
			t.Errorf("Unexpected value in CurrL2, was %v, expected %v", v.CurrL2, currL2)
		}
		if v.CurrL3 != currL3 {
			t.Errorf("Unexpected value in CurrL3, was %v, expected %v", v.CurrL3, currL3)
		}
		if v.VoltL1 != voltL1 {
			t.Errorf("Unexpected value in VoltL1, was %v, expected %v", v.VoltL1, voltL1)
		}
		if v.VoltL2 != voltL2 {
			t.Errorf("Unexpected value in VoltL2, was %v, expected %v", v.VoltL2, voltL2)
		}
		if v.VoltL3 != voltL3 {
			t.Errorf("Unexpected value in VoltL3, was %v, expected %v", v.VoltL3, voltL3)
		}
	default:
		t.Errorf("Only expected ThreeFasesMessageType2, got %T", message)
	}
}
func TestItem18(t *testing.T) {
	m := []byte{0x7e, 0xa0, 0x9a, 0x01, 0x02, 0x01, 0x10, 0xaa, 0xa5, 0xe6, 0xe7, 0x00, 0x0f, 0x40, 0x00, 0x00, 0x00, 0x09, 0x0c, 0x07, 0xe2, 0x05, 0x1a, 0x06, 0x15, 0x00, 0x0a, 0xff, 0x80, 0x00, 0x00, 0x02, 0x12, 0x09, 0x07, 0x4b, 0x46, 0x4d, 0x5f, 0x30, 0x30, 0x31, 0x09, 0x10, 0x36, 0x39, 0x37, 0x30, 0x36, 0x33, 0x31, 0x34, 0x30, 0x31, 0x34, 0x36, 0x38, 0x38, 0x30, 0x33, 0x09, 0x07, 0x4d, 0x41, 0x33, 0x30, 0x34, 0x48, 0x34, 0x06, 0x00, 0x00, 0x04, 0x3e, 0x06, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x02, 0x27, 0x06, 0x00, 0x00, 0x0c, 0x57, 0x06, 0x00, 0x00, 0x07, 0x14, 0x06, 0x00, 0x00, 0x02, 0xe6, 0x06, 0x00, 0x00, 0x09, 0x6b, 0x06, 0x00, 0x00, 0x09, 0x87, 0x06, 0x00, 0x00, 0x09, 0x7d, 0x09, 0x0c, 0x07, 0xe2, 0x05, 0x1a, 0x06, 0x15, 0x00, 0x0a, 0xff, 0x80, 0x00, 0x00, 0x06, 0x01, 0x3a, 0x28, 0xcf, 0x06, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x13, 0x00, 0x06, 0x00, 0x1c, 0xd0, 0x52, 0x89, 0x6b, 0x7e}
	//m := []byte{0x7e, 0xa0, 0x9a, 0x01, 0x02, 0x01, 0x10, 0xaa, 0xa5, 0xe6, 0xe7, 0x00, 0x0f, 0x40, 0x00, 0x00, 0x00, 0x09, 0x0c, 0x07, 0xe2, 0x05, 0x1b, 0x07, 0x14, 0x00, 0x0a, 0xff, 0x80, 0x00, 0x00, 0x02, 0x12, 0x09, 0x07, 0x4b, 0x46, 0x4d, 0x5f, 0x30, 0x30, 0x31, 0x09, 0x10, 0x36, 0x39, 0x37, 0x30, 0x36, 0x33, 0x31, 0x34, 0x30, 0x31, 0x34, 0x36, 0x38, 0x38, 0x30, 0x33, 0x09, 0x07, 0x4d, 0x41, 0x33, 0x30, 0x34, 0x48, 0x34, 0x06, 0x00, 0x00, 0x03, 0x65, 0x06, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x01, 0xe0, 0x06, 0x00, 0x00, 0x0a, 0x11, 0x06, 0x00, 0x00, 0x07, 0x0c, 0x06, 0x00, 0x00, 0x01, 0xd2, 0x06, 0x00, 0x00, 0x09, 0x4b, 0x06, 0x00, 0x00, 0x09, 0x5f, 0x06, 0x00, 0x00, 0x09, 0x58, 0x09, 0x0c, 0x07, 0xe2, 0x05, 0x1b, 0x07, 0x14, 0x00, 0x0a, 0xff, 0x80, 0x00, 0x00, 0x06, 0x01, 0x3a, 0xaf, 0x1c, 0x06, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x13, 0x00, 0x06, 0x00, 0x1c, 0xfb, 0xe8, 0x15, 0x8e, 0x7e}
	message, err := BytesParser(m)
	if err != nil {
		t.Errorf("Didn't expect error... %v", err)
		return
	}

	messageType := messageTypes(18)
	meterTime, _ := time.Parse(layout, "2018-05-26T21:00:10")
	obisListVersion := "KFM_001"
	gs1 := "6970631401468803"
	meterModel := "MA304H4"
	actPowPos := 1086
	actPowNeg := 0
	reactPowPos := 0
	reactPowNeg := 551
	currL1 := 3159
	currL2 := 1812
	currL3 := 742
	voltL1 := 2411
	voltL2 := 2439
	voltL3 := 2429
	dateTime, _ := time.Parse(layout, "2018-05-26T06:21:00")
	actEnergyPa := 20588751
	actEnergyMa := 0
	actEnergyPr := 4864
	actEnergyMr := 1888338

	switch v := message.(type) {
	case *ThreeFaseMessageType3:
		if v.MessageType != messageType {
			t.Errorf("Unexpected value in MessageType, was %v, expected %v", v.MessageType, messageType)
		}
		if v.MeterTime.Format(layout) != meterTime.Format(layout) {
			t.Errorf("Unexpected value in MeterTime, was %v, expected %v", v.MeterTime.Format(layout), meterTime.Format(layout))
		}
		if v.HostTime == (time.Time{}) {
			t.Errorf("Unexpected value in HostTime, was %v", v.HostTime)
		}
		if v.ObisListVersion != obisListVersion {
			t.Errorf("Unexpected value in ObisListVersion, was %v, expected %v", v.ObisListVersion, obisListVersion)
		}
		if v.Gs1 != gs1 {
			t.Errorf("Unexpected value in Gs1, was %v, expected %v", v.Gs1, gs1)
		}
		if v.MeterModel != meterModel {
			t.Errorf("Unexpected value in MeterModel, was %v, expected %v", v.MeterModel, meterModel)
		}
		if v.ActPowPos != actPowPos {
			t.Errorf("Unexpected value in ActPowPos, was %v, expected %v", v.ActPowPos, actPowPos)
		}
		if v.ActPowNeg != actPowNeg {
			t.Errorf("Unexpected value in ActPowNeg, was %v, expected %v", v.ActPowNeg, actPowNeg)
		}
		if v.ReactPowPos != reactPowPos {
			t.Errorf("Unexpected value in ReactPowPos, was %v, expected %v", v.ReactPowPos, reactPowPos)
		}
		if v.ReactPowNeg != reactPowNeg {
			t.Errorf("Unexpected value in ReactPowNeg, was %v, expected %v", v.ReactPowNeg, reactPowNeg)
		}
		if v.CurrL1 != currL1 {
			t.Errorf("Unexpected value in CurrL1, was %v, expected %v", v.CurrL1, currL1)
		}
		if v.CurrL2 != currL2 {
			t.Errorf("Unexpected value in CurrL2, was %v, expected %v", v.CurrL2, currL2)
		}
		if v.CurrL3 != currL3 {
			t.Errorf("Unexpected value in CurrL3, was %v, expected %v", v.CurrL3, currL3)
		}
		if v.VoltL1 != voltL1 {
			t.Errorf("Unexpected value in VoltL1, was %v, expected %v", v.VoltL1, voltL1)
		}
		if v.VoltL2 != voltL2 {
			t.Errorf("Unexpected value in VoltL2, was %v, expected %v", v.VoltL2, voltL2)
		}
		if v.VoltL3 != voltL3 {
			t.Errorf("Unexpected value in VoltL3, was %v, expected %v", v.VoltL3, voltL3)
		}
		if v.DateTime.Format(layout) != dateTime.Format(layout) {
			t.Errorf("Unexpected value in DateTime, was %v, expected %v", v.DateTime, dateTime)
		}
		if v.ActEnergyPa != actEnergyPa {
			t.Errorf("Unexpected value in ActEnergyPa, was %v, expected %v", v.ActEnergyPa, actEnergyPa)
		}
		if v.ActEnergyMa != actEnergyMa {
			t.Errorf("Unexpected value in ActEnergyMa, was %v, expected %v", v.ActEnergyMa, actEnergyMa)
		}
		if v.ActEnergyPr != actEnergyPr {
			t.Errorf("Unexpected value in ActEnergyPr, was %v, expected %v", v.ActEnergyPr, actEnergyPr)
		}
		if v.ActEnergyMr != actEnergyMr {
			t.Errorf("Unexpected value in ActEnergyMr, was %v, expected %v", v.ActEnergyMr, actEnergyMr)
		}
	default:
		t.Errorf("Only expected ThreeFasesMessageType2, got %T", message)
	}
}

func TestCrc16Valid(t *testing.T) {
	m := []byte{0x7e, 0xa0, 0x27, 0x01, 0x02, 0x01, 0x10, 0x5a, 0x87, 0xe6, 0xe7, 0x00, 0x0f, 0x40, 0x00, 0x00, 0x00, 0x09, 0x0c, 0x07, 0xe2, 0x05, 0x1b, 0x07, 0x13, 0x25, 0x26, 0xff, 0x80, 0x00, 0x00, 0x02, 0x01, 0x06, 0x00, 0x00, 0x03, 0x52, 0x81, 0x3d, 0x7e}
	length := len(m)
	msb := uint16(m[length-2])
	lsb := uint16(m[length-3])
	expecting := (msb << 8) | lsb
	actual := crc16.ChecksumCCITT(m[1 : length-3])

	if expecting != actual {
		t.Errorf("expected checksum to be %x but was %x", expecting, actual)
	}
}

func TestCrc16Invalid(t *testing.T) {
	m := []byte{0x7e, 0xa1, 0x27, 0x01, 0x02, 0x01, 0x10, 0x5a, 0x87, 0xe6, 0xe7, 0x00, 0x0f, 0x40, 0x00, 0x00, 0x00, 0x09, 0x0c, 0x07, 0xe2, 0x05, 0x1b, 0x07, 0x13, 0x25, 0x26, 0xff, 0x80, 0x00, 0x00, 0x02, 0x01, 0x06, 0x00, 0x00, 0x03, 0x52, 0x81, 0x3d, 0x7e}
	length := len(m)
	msb := uint16(m[length-2])
	lsb := uint16(m[length-3])
	expecting := (msb << 8) | lsb
	actual := crc16.ChecksumCCITT(m[1 : length-3])

	if expecting == actual {
		t.Errorf("expected checksums to be different but both was %x", expecting)
	}
}
