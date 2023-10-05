package main

import (
	"encoding/binary"
	"fmt"
	"log"
)

type decodedPacket struct {
	Short1  int16
	Twelve  string
	Single  byte
	Eight   string
	Short2  int16
	Fifteen string
	Long    int32
}

func decodePacket(packet []byte) (decodedPacket, error) {
	if len(packet) != 44 {
		return decodedPacket{}, fmt.Errorf("packet size should be 44 bytes")
	}

	p := decodedPacket{}
	p.Short1 = int16(binary.BigEndian.Uint16(packet[:2]))
	p.Twelve = string(packet[2:14])
	p.Single = packet[14]
	p.Eight = string(packet[15:23])
	p.Short2 = int16(binary.BigEndian.Uint16(packet[23:25]))
	p.Fifteen = string(packet[25:40])
	p.Long = int32(binary.BigEndian.Uint32(packet[40:]))

	return p, nil
}

func main() {
	packet := []byte{
		0x04, 0xD2,
		0x6B, 0x65, 0x65, 0x70, 0x64, 0x65, 0x63, 0x6F, 0x64, 0x69, 0x6E, 0x67,
		0x38,
		0x64, 0x6F, 0x6E, 0x74, 0x73, 0x74, 0x6F, 0x70,
		0x03, 0x15,
		0x63, 0x6F, 0x6E, 0x67, 0x72, 0x61, 0x74, 0x75, 0x6C, 0x61, 0x74, 0x69, 0x6F, 0x6E, 0x73,
		0x07, 0x5B, 0xCD, 0x15,
		// 0xFF, 0xFF, 0xFB, 0x2E, // -1234
	}

	message, err := decodePacket(packet)
	if err != nil {
		log.Fatalf("decoding failed: %s", err)
	}

	log.Printf("%v\n", message)
}
