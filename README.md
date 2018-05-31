What is this
============


This libary converts a byte stream (chan byte) into structs for parsing messages over MBus for the Advanced Metering System (AMS), also known as a Smart meter for measuring power consumption.


Usage
-----

    // byteStream is a chan byte, you need to create this yourself
    bytePackages := ams.ByteReader(byteStream)
    messages := ams.ByteParser(bytePackages)

    for message := range messages {
        if message.Error != nil {
            fmt.Println("[ERROR]", message.Error)
        } else {
            jsonString, _ := json.Marshal(message.Data)
            fmt.Printf("%s\n", jsonString)
        }
    }

If you need logging you can easily add that your self:

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
        }()
        return out
    }

    func byteArrayToHexStringArray(bytes []byte) (strings []string) {
        for _, b := range bytes {
            strings = append(strings, fmt.Sprintf("0x%02x", b))
        }
        return
    }



