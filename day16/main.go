package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	bytes := make([]byte, 10240)
	stream := strings.NewReader(input)
	packet := ""
	read, _ := stream.Read(bytes)
	for i := range bytes[:read] {
		hex, _ := strconv.ParseInt(string(bytes[i]), 16, 64)
		packet += fmt.Sprintf("%04b", hex)
	}

	parser := Parser{packet}
	bitsread, sumVersions, total := parser.ParsePacket(0, 0, []int64{}, 0)

	fmt.Printf("#%v,%v,%v", bitsread, sumVersions, total)
}
