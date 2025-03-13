package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
)

var input string
var unencoded string

func main() {
	flag.StringVar(&unencoded, "bytes", "", "A byte string to convert in 0x format, comma separated (i.e. \"0x01, 0x02, 0x03\").")
	flag.Parse()

	if unencoded == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter a byte string in hex format (i.e. '0x01, 0x02, 0x03'):")
		input, _ = reader.ReadString('\n')
	} else {
		input = unencoded
	}

	input = strings.TrimSpace(input)

	input = strings.ReplaceAll(input, "0x", "")
	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, " ", "")

	buf, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println("Invalid hex string:", err)
		return
	}

	encoded := make([]byte, len(buf))
	for i := 0; i < len(buf); i++ {
		encoded[i] = (buf[i] + 2) & 0xFF
	}

	hexString := ""
	for i, b := range encoded {
		hexString += fmt.Sprintf("0x%02x", b)
		if i < len(encoded)-1 {
			hexString += ","
		}
	}

	fmt.Println("Encoded string:", hexString)
}
