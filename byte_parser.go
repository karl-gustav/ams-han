package ams

func ByteParser(bytePackages chan []byte) chan Message {
	messages := make(chan Message)
	go func() {
		for bytePackage := range bytePackages {
			item, err := bytesToItem(bytePackage)
			messages <- Message{
				Data:  item,
				Error: err,
			}
		}
		close(messages)
	}()
	return messages
}
