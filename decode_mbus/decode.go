package decode_mbus

import "fmt"

type HanMessage struct{}

const (
	startByte       = 0x7E
	bytesRead       = 3
	frameDelimiters = 2 /* start byte 0x7e + end byte 0x7e*/
)

const (
	Type3 = 0xA
)

var (
	ErrWrongLength           = fmt.Errorf("Wrong length of buffer, should be 3")
	ErrWrongStartByte        = fmt.Errorf("Expected 126 as first byte")
	ErrWrongFrameFormatField = fmt.Errorf("Expected second byte to start with 0xA")
)

// 2 Byte Frame Format Field
// 16 BITS: "TTTTSLLLLLLLLLLL"
// - T=Type bits: TTTT = 0101 (0xa0) = Type 3
// - S=Segmentation=0 (Segment = 1)
// - L=11 Length Bits
func FindStart(buffer []byte) (int, error) {
	if len(buffer) != 3 {
		return 0, ErrWrongLength
	}
	if buffer[0] != startByte {
		// fmt.Printf("start byte was: %02x (%08b)\n", buffer[0], buffer[0])
		return 0, ErrWrongStartByte
	}
	if getFirst4Bits(buffer[1], 4) != Type3 {
		return 0, ErrWrongFrameFormatField
	}

	first3Bits := int(buffer[1] & 0x7)
	lenght := first3Bits<<8 + int(buffer[2])
	return lenght + frameDelimiters - bytesRead, nil
}

func getFirst4Bits(input byte, n uint) byte {
	return input >> 4
}

/*

   // 2 Byte Frame Format Field
   // 16 BITS: "TTTTSLLLLLLLLLLL"
   // - T=Type bits: TTTT = 0101 (0xa0) = Type 3
   // - S=Segmentation=0 (Segment = 1)
   // - L=11 Length Bits
     if ((buf[1] & 0xa0) == 0xa0) { // 16 bit frame format field = "0101SLLLLLLLLLLL" - Type 3 = 0101 (0xa0), S=Segmentation, L=Length bits
   n = my_read(read_fd, buf+2, 1);
   if (debug) {printf(" C(%02x)", buf[2]); fflush(stdout);}
   if (n<=0) return 0; // EOF or ERROR
   length = ((buf[1] & 0x07) << 8) + buf[2]; // length is 11 bits
   break;
     }
*/
