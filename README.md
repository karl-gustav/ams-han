What is this
============


This libary converts a byte stream (chan byte) into structs for parsing messages over MBus for the Advanced Metering System (AMS), also known as a Smart meter for measuring power consumption.


Usage
-----

    func main() {
        byteStream := getByteChannel() // byteStream is a `chan byte`, you need to create this yourself

        bytePackages, errors := ams.ByteReader(byteStream)
        printErrors(errors)
        messages, errors := ams.ByteParser(bytePackages)
        printErrors(errors)

        for message := range messages {
            jsonString, _ := json.Marshal(message)
            fmt.Printf("%s\n", jsonString)
        }
    }

    func printErrors(errors chan error) {
        go func() {
            for err := range errors {
                fmt.Println("[ERROR]", err)
            }
        }()
    }

If you need logging of the byte channel you can easily add that your self:

    bytePackages := ams.ByteReader(byteStream)
    if verbose {
        bytePackages = channelLogger(bytePackages)
    }
    messages := ams.ByteParser(bytePackages)

    func channelLogger(in chan []byte) chan []byte {
        out := make(chan []byte)
        go func() {
            for bytes := range in {
                fmt.Printf("\nBuffer(%d): \n[%s]\n", len(bytes), strings.Join(byteArrayToHexStringArray(bytes), ", "))
                out <- bytes
            }
            close(out)
        }()
        return out
    }

    func byteArrayToHexStringArray(bytes []byte) (strings []string) {
        for _, b := range bytes {
            strings = append(strings, fmt.Sprintf("0x%02x", b))
        }
        return
    }



