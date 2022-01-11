package main

import (
	"fmt"
	"strconv"
)

type Parser struct {
	packet string
}

func showPacketType(packetType int64) string {
	switch packetType {
	case 0: // sum
		return "sum"
	case 1: // product
		return "product"
	case 2: // minimum
		return "minimum"
	case 3: // maximum
		return "maximum"
	case 4: // maximum
		return "literal"
	case 5: // greater than, 1 if value of first packet > 2nd, else 0
		return "greater than"
	case 6: // less than, 1 if value of first packet < second, else 0
		return "less than"
	case 7: // equal to, 1 if values of packets are equal, else 0
		return "equals"
	default: // equal to, 1 if values of packets are equal, else 0
		return "unknown"
	}
}

// ParsePacket interprets the binary literal (0100110) according to the Buoyancy Interchange Transmission System (BITS) rules
// limit sets a max number of packets to parse, and is set when parsing a sub-packet stream. Otherwise, 0 indicates no limit
// versions is the sum of all the version headers parsed so far
func (parser Parser) ParsePacket(limit int64, versions int64, literals []int64, packetsRead int64) (int64, int64, []int64) {
	var read int64 = 0
	sequence := 0
	for {
		if read >= int64(len(parser.packet)) {
			break
		}

		remaining := parser.packet[read:]
		if lastBits, _ := strconv.ParseInt(remaining, 2, 64); lastBits == 0 {
			fmt.Printf("Last %d bits are 0s. End of the packet\n", len(remaining))
			read += int64(len(remaining))
			break
		}

		if limit > 0 && packetsRead >= limit {
			fmt.Printf("Sub-packet limit reached %d (limit: %d)\n", packetsRead, limit)
			break
		}

		if sequence == 0 {
			packetsRead += 1
			bits, err := strconv.ParseInt(parser.packet[read:read+3], 2, 64)
			if err != nil {
				fmt.Printf("Cannot parse version header: %s", err.Error())
			} else {
				fmt.Printf("Version %d\n", bits)
			}
			read += 3
			versions += bits
			sequence++
		}

		if sequence == 1 {
			packetType, err := strconv.ParseInt(parser.packet[read:3+read], 2, 64)
			if err != nil {
				fmt.Printf("Cannot parse packet type header: %s", err.Error())
			} else {
				fmt.Printf("Packet type %s\n", showPacketType(packetType))
			}
			read += 3

			switch packetType {
			case 4:
				literal := ""
				// read literal value 5 bits at a time
				for {
					fmt.Printf("%v", parser.packet[read:1+read])
					header, _ := strconv.ParseInt(parser.packet[read:1+read], 2, 64)
					read += 1
					fmt.Printf("%v", parser.packet[read:4+read])
					literal += parser.packet[read : read+4]
					read += 4
					if header == 0 {
						sequence = 0
						break
					}
				}
				fmt.Println()
				dec, _ := strconv.ParseInt(literal[:], 2, 64)
				literals = append(literals, dec)
				fmt.Printf("literal: %v\n", dec)
			default:
				// sub-packets
				header, _ := strconv.ParseInt(parser.packet[read:1+read], 2, 64)
				read += 1
				var operands []int64
				if header == 0 {
					size, _ := strconv.ParseInt(parser.packet[read:15+read], 2, 64)
					fmt.Printf("sub-packets (%s) length: %v\n", parser.packet[read:15+read], size)
					read += 15
					fmt.Printf("sub-packets: %s\n", parser.packet[read:read+size])
					subParser := Parser{packet: parser.packet[read : read+size]}
					fmt.Printf("read before: %v\n", read)
					subRead, subVersions, values := subParser.ParsePacket(0, 0, []int64{}, 0)
					for i := range values {
						operands = append(operands, values[i])
					}
					read += subRead
					versions += subVersions
					fmt.Printf("read after: %v\n", read)
				} else if header == 1 {
					size, _ := strconv.ParseInt(parser.packet[read:11+read], 2, 64)
					fmt.Printf("# sub-packets: %v\n", size)
					read += 11
					subParser := Parser{packet: parser.packet[read:]}
					fmt.Printf("read before: %v\n", read)
					subRead, subVersions, values := subParser.ParsePacket(size, 0, []int64{}, 0)
					for i := range values {
						operands = append(operands, values[i])
					}
					read += subRead
					versions += subVersions
					fmt.Printf("read after: %v\n", read)
				}
				fmt.Printf("Operate %s on %d\n", showPacketType(packetType), operands)
				switch packetType {
				case 0: // sum
					sum := calc(operands, func(val1, val2 int64) int64 {
						return val1 + val2
					}, 0)
					literals = append(literals, sum)
					fmt.Printf("Sum returning %d\n", sum)
				case 1: // product
					product := calc(operands, func(val1, val2 int64) int64 {
						return val1 * val2
					}, 1)
					fmt.Printf("Product returning %d\n", product)
					literals = append(literals, product)
				case 2: // minimum
					minimum := calc(operands, func(val1, val2 int64) int64 {
						if val1 < val2 {
							return val1
						}
						return val2
					}, 9999999)
					fmt.Printf("Minimum returning %d\n", minimum)
					literals = append(literals, minimum)
				case 3: // maximum
					maximum := calc(operands, func(val1, val2 int64) int64 {
						if val1 > val2 {
							return val1
						}
						return val2
					}, -1)
					fmt.Printf("Maximum returning %d\n", maximum)
					literals = append(literals, maximum)
				case 5: // greater than, 1 if value of first packet > 2nd, else 0
					literals = append(literals, pairwiseOperation(operands, func(val1, val2 int64) int64 {
						if val1 > val2 {
							fmt.Printf("%d > %d returning 1\n", val1, val2)
							return 1
						}
						fmt.Printf("%d < %d returning 0\n", val1, val2)
						return 0
					}))
				case 6: // less than, 1 if value of first packet < second, else 0
					literals = append(literals, pairwiseOperation(operands, func(val1, val2 int64) int64 {
						if val1 < val2 {
							fmt.Printf("%d < %d returning 1\n", val1, val2)
							return 1
						}
						fmt.Printf("%d > %d returning 0\n", val1, val2)
						return 0
					}))
				case 7: // equal to, 1 if values of packets are equal, else 0
					literals = append(literals, pairwiseOperation(operands, func(val1, val2 int64) int64 {
						if val1 == val2 {
							fmt.Printf("%d = %d returning 1\n", val1, val2)
							return 1
						}

						fmt.Printf("%d != %d returning 0\n", val1, val2)
						return 0
					}))
				}
				sequence = 0
			}
		}
	}

	return read, versions, literals
}

type Operator func(val1, val2 int64) int64

func calc(values []int64, fn Operator, starting int64) int64 {
	total := starting
	for _, v := range values {
		total = fn(total, v)
	}
	return total
}

func pairwiseOperation(values []int64, fn Operator) int64 {
	if len(values) != 2 {
		return 0
	}

	return fn(values[0], values[1])
}
