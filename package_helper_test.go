package ams

import (
	"testing"
)

func TestEmptyBuffer(t *testing.T) {
	err := VerifyStart([]byte{})

	if err == nil {
		t.Errorf("Expected error because buffer was empty")
	}
}

func TestStartByte(t *testing.T) {
	err := VerifyStart([]byte{1, 2, 3})

	if err == nil {
		t.Errorf("Expected error because first byte wasn't 126")
	}
}

func TestFrameFormatField(t *testing.T) {
	err := VerifyStart([]byte{0x7e, 0x00, 0x00})
	if err == nil {
		t.Errorf("Expected error because second byte didn't start with 0101 (0xA)")
	}
}

func TestLenght(t *testing.T) {
	lenght, _ := ReadLenght([]byte{0x7e, 0xa1, 0x78})
	if lenght != 375 {
		t.Errorf("Expected lenght to be 375 (0x01 <<8 0x78), was %v", lenght)
	}
}
