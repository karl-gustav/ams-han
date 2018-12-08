What is this
============


This libary converts a byte stream (chan byte) into structs for parsing messages over MBus for the Advanced Metering System (AMS), also known as a Smart meter for measuring power consumption.


Usage
-----

	func main() {
		byteStream := getByteChannel() // byteStream is a `chan byte`, read more below
		next := ams.ByteReader(byteStream)
		for {
			bytePackage, _ := next()
			message, _ := ams.BytesParser(bytePackage)
			fmt.Printf("%+v\n", message)
		}
	}

`byteStream` is a `chan byte` from the serial port, you need to create this yourself a good example can be found here:
https://github.com/karl-gustav/ams-han-cmd/blob/master/main.go

With error handlin:

	func main() {
		byteStream := byteChannel()
		next := ams.ByteReader(byteStream)
		for {
			bytePackage, err := next()
			if err != nil {
				fmt.Println(err)
				if err == ams.CHANNEL_IS_CLOSED_ERROR {
					return
				}
				continue
			}

			message, err := ams.BytesParser(bytePackage)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%+v\n", message)
		}
	}

Helper function to make printing []byte easier:

    func byteArrayToHexStringArray(bytes []byte) (strings []string) {
        for _, b := range bytes {
            strings = append(strings, fmt.Sprintf("0x%02x", b))
        }
        return
    }



