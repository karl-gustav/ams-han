package ams

func ByteParser(bytePackages chan []byte) (chan interface{}, chan error) {
	messages := make(chan interface{})
	errors := make(chan error)
	go func() {
		for bytePackage := range bytePackages {
			item, err := bytesToItem(bytePackage)
			if err != nil {
				errors <- err
			} else {
				messages <- item
			}
		}
		close(messages)
		close(errors)
	}()
	return messages, errors
}
