package solutions

import (
	"errors"
	"log"
)

var hexadecimal map[byte][]bit = map[byte][]bit{
	'0': {false, false, false, false},
	'1': {false, false, false, true},
	'2': {false, false, true, false},
	'3': {false, false, true, true},
	'4': {false, true, false, false},
	'5': {false, true, false, true},
	'6': {false, true, true, false},
	'7': {false, true, true, true},
	'8': {true, false, false, false},
	'9': {true, false, false, true},
	'A': {true, false, true, false},
	'B': {true, false, true, true},
	'C': {true, true, false, false},
	'D': {true, true, false, true},
	'E': {true, true, true, false},
	'F': {true, true, true, true},
}

const versionBits int = 3
const typeIDbits int = 3
const literalID uint64 = 4

type packet struct {
	version  uint64
	typeID   uint64
	value    *uint64
	operator *operator
}

const byteHigh = byte(1) << 7

type operator struct {
	lengthType uint64
	size       uint64
	subpackets []packet
}

type bit bool // true == 1
type literal int

func F16a(input string) int {
	bits := parseBits(input)
	p, bits, err := parsePacket(bits)
	if err != nil {
		log.Fatalf("Error parsing packet: %s", err)
	}
	return int(sumVersions(p))
}

func F16b(input string) int {
	bits := parseBits(input)
	p, bits, err := parsePacket(bits)
	if err != nil {
		log.Fatalf("Error parsing packet: %s", err)
	}
	result := evaluatePacket(p)
	return int(result)
}

func evaluatePacket(p packet) uint64 {
	var result uint64
	switch p.typeID {
	case 0:
		result = 0
		for i := range p.operator.subpackets {
			result += evaluatePacket(p.operator.subpackets[i])
		}
		return result
	case 1:
		result = 1
		for i := range p.operator.subpackets {
			result *= evaluatePacket(p.operator.subpackets[i])
		}
		return result
	case 2:
		result := evaluatePacket(p.operator.subpackets[0])
		for i := 1; i < len(p.operator.subpackets); i++ {
			current := evaluatePacket(p.operator.subpackets[i])
			if current < result {
				result = current
			}
		}
		return result
	case 3:
		result := evaluatePacket(p.operator.subpackets[0])
		for i := 1; i < len(p.operator.subpackets); i++ {
			current := evaluatePacket(p.operator.subpackets[i])
			if current > result {
				result = current
			}
		}
		return result
	case 4:
		return uint64(*p.value)
	case 5:
		if len(p.operator.subpackets) != 2 {
			log.Fatalf("Wrong number of subpackets for type %d", p.typeID)
		}
		if evaluatePacket(p.operator.subpackets[0]) > evaluatePacket(p.operator.subpackets[1]) {
			return 1
		} else {
			return 0
		}
	case 6:
		if len(p.operator.subpackets) != 2 {
			log.Fatalf("Wrong number of subpackets for type %d", p.typeID)
		}
		if evaluatePacket(p.operator.subpackets[0]) < evaluatePacket(p.operator.subpackets[1]) {
			return 1
		} else {
			return 0
		}
	case 7:
		if len(p.operator.subpackets) != 2 {
			log.Fatalf("Wrong number of subpackets for type %d", p.typeID)
		}
		if evaluatePacket(p.operator.subpackets[0]) == evaluatePacket(p.operator.subpackets[1]) {
			return 1
		} else {
			return 0
		}
	}
	log.Fatalf("This should never happen")
	return uint64(0)
}

func parsePacket(b []bit) (packet, []bit, error) {
	var version, typeID uint64
	var value *uint64 = nil
	var op *operator = nil
	var err error
	version, b, err = parseVersion(b)
	typeID, b, err = parseTypeID(b)
	if typeID == literalID {
		value, b, err = parseLiteral(b)
	} else {
		op, b, err = parseOperator(b)
	}
	if err != nil {
		log.Fatalf("Error encountered parsing packet: %s", err)
	}
	p := packet{
		version,
		typeID,
		value,
		op,
	}
	return p, b, err
}

func parseOperator(b []bit) (*operator, []bit, error) {
	var lengthType, size uint64
	var err error
	var current packet
	subpackets := make([]packet, 0)
	lengthType, b, err = getBits(b, 1)
	if lengthType == 0 {
		size, b, err = getBits(b, 15)
		payload := b[0:size]
		b = b[size:]
		for len(payload) > 0 {
			current, payload, err = parsePacket(payload)
			subpackets = append(subpackets, current)
		}
	} else {
		size, b, err = getBits(b, 11)
		for i := uint64(0); i < size; i++ {
			current, b, err = parsePacket(b)
			subpackets = append(subpackets, current)
		}
	}
	return &operator{
		lengthType,
		size,
		subpackets,
	}, b, err
}

func parseLiteral(b []bit) (*uint64, []bit, error) {
	var result uint64 = 0
	var done uint64 = 1
	var err error
	var current uint64
	for done != 0 {
		done, b, err = getBits(b, 1)
		result <<= 4
		current, b, err = getBits(b, 4)
		result += current
	}
	return &(result), b, err
}

func parseTypeID(b []bit) (uint64, []bit, error) {
	return getBits(b, typeIDbits)
}

func parseVersion(b []bit) (uint64, []bit, error) {
	return getBits(b, versionBits)
}

func getBits(b []bit, n int) (uint64, []bit, error) {
	var result uint64 = 0
	for i := 0; i < n; i++ {
		if len(b) == 0 {
			return result, b, errors.New("No more bits")
		}
		result <<= 1
		if b[0] {
			result += 1
		}
		b = b[1:]
	}
	return result, b, nil
}

func sumVersions(p packet) uint64 {
	if p.typeID == literalID {
		return p.version
	}
	result := p.version
	for i := range p.operator.subpackets {
		result += sumVersions(p.operator.subpackets[i])
	}
	return result
}

func bytesToBits(b []byte) []bit {
	result := make([]bit, 0)
	for i := range b {
		for j := 0; j < 8; j++ {
			result = append(result, (b[i]&1<<(7-j) > 0))
		}
	}
	return result
}

func parseBits(input string) []bit {
	result := make([]bit, 0)
	for i := range input {
		current := hexadecimal[input[i]]
		for j := range current {
			result = append(result, current[j])
		}
	}
	return result
}
