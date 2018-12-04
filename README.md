What is this
============


This libary converts a byte stream (chan byte) into structs for parsing messages over MBus for the Advanced Metering System (AMS), also known as a Smart meter for measuring power consumption.


Usage
-----

	func main() {
		byteStream := getByteChannel() // byteStream is a `chan byte`, you need to create this yourself
		next := ams.ByteReader(byteStream)
		for {
			bytePackage, _ := next()
			message, _ := ams.BytesParser(bytePackage)
			fmt.Printf("%+v\n", message)
		}
	}

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



