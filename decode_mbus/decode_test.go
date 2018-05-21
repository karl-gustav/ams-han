package decode_mbus

import (
	"testing"
)

func TestEmptyBuffer(t *testing.T) {
	_, err := FindStart([]byte{})

	if err != ErrWrongLength {
		t.Errorf("Expected error because buffer was empty")
	}
}

func TestStartByte(t *testing.T) {
	_, err := FindStart([]byte{33, 22, 33})

	if err != ErrWrongStartByte {
		t.Errorf("Expected error because first byte wasn't 126")
	}
}

func TestFrameFormatField(t *testing.T) {
	_, err := FindStart([]byte{0x7e, 0x00, 0x00})
	if err != ErrWrongFrameFormatField {
		t.Errorf("Expected error because second byte didn't start with 0101 (0xA)")
	}
}

func TestLenght(t *testing.T) {
	lenght, _ := FindStart([]byte{0x7e, 0xa1, 0x78})
	if lenght != 375 {
		t.Errorf("Expected lenght to be 375 (0x01 + 0x78), was %v", lenght)
	}
}
