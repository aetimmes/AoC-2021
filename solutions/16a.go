package solutions

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

const versionBits int = 3
const typeIDbits int = 3
const literalID int = 4

type packet struct {
	version  int
	typeID   int
	value    *literal
	operator *operator
}

const byteHigh = byte(1) << 7

type operator struct {
	lengthType int
	size       int
	subpackets []packet
}

type bit bool // true == 1
type literal int

func F16a(input string) int {
	bytes, err := hex.DecodeString(input)
	bits := bytesToBits(bytes)
	if err != nil {
		log.Fatalf("Error parsing hex: %s", err)
	}
	p, bits, err := parsePacket(bits)
	if err != nil {
		log.Fatalf("Error parsing packet: %s", err)
	}
	return sumVersions(*p)
}

func parsePacket(b []bit) (*packet, []bit, error) {
	var version, typeID int
	var value *literal = nil
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
		fmt.Printf("Error encountered parsing packet: %s", err)
		return nil, b, err
	}
	return &packet{
		version,
		typeID,
		value,
		op,
	}, b, err
}

func parseOperator(b []bit) (*operator, []bit, error) {
	var lengthType, size int
	var err error
	var current *packet
	subpackets := make([]packet, 0)
	lengthType, b, err = getBits(b, 1)
	if lengthType == 0 {
		size, b, err = getBits(b, 15)
		payload := b[0:size]
		b = b[size:]
		for len(payload) > 0 {
			current, payload, err = parsePacket(payload)
			subpackets = append(subpackets, *(current))
		}
	} else {
		size, b, err = getBits(b, 11)
	}
	return &operator{
		lengthType,
		size,
		subpackets,
	}, b, err
}

func parseLiteral(b []bit) (*literal, []bit, error) {
	sum := 0
	done := 1
	var err error
	var current int
	for done != 0 {
		done, b, err = getBits(b, 1)
		sum <<= 8
		current, b, err = getBits(b, 4)
		sum += current
	}
	result := literal(sum)
	return &(result), b, err
}

func parseTypeID(b []bit) (int, []bit, error) {
	return getBits(b, typeIDbits)
}

func parseVersion(b []bit) (int, []bit, error) {
	return getBits(b, versionBits)
}

func getBits(b []bit, n int) (int, []bit, error) {
	result := 0
	for i := 0; i < n; i++ {
		if len(b) == 0 {
			return result, b, errors.New("No more bits")
		}
		result = result << 1
		if b[0] {
			result += 1
		}
		b = b[1:]
	}
	return result, b, nil
}

func sumVersions(p packet) int {
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
	result := make([]bit, 0, 8*len(b))
	for i := range b {
		for j := 0; j < 8; j++ {
			result[8*i+j] = (b[i]&1<<(7-j) > 0)
		}
	}
	return result
}
