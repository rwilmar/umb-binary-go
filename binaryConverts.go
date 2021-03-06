package main

import (
	"bytes"
	"encoding/binary"
)

func uint16ToLitteEndian(num uint16) []byte {
	var lowByte byte = byte(num & 255)
	num = num >> 8
	var hiByte byte = byte(num & 255)
	return []byte{lowByte, hiByte}
}

func uint32ToLitteEndian(num uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b[0:], num)
	return b
}

func uint32FromLitteEndian(num []byte) uint32 {
	return binary.LittleEndian.Uint32(num[0:])
}
func uint16FromLitteEndian(num []byte) uint16 {
	return binary.LittleEndian.Uint16(num[0:])
}

func float64FromLittleEndian(num []byte) (float64, error) {
	var res float64
	num = num[:8]
	buf := bytes.NewReader(num)
	err := binary.Read(buf, binary.LittleEndian, &res)
	return res, err
}

func float32FromLittleEndian(num []byte) (float32, error) {
	var res float32
	num = num[:4]
	buf := bytes.NewReader(num)
	err := binary.Read(buf, binary.LittleEndian, &res)
	return res, err
}

func int32FromLittleEndian(num []byte) (int32, error) {
	var res int32
	num = num[:4]
	buf := bytes.NewReader(num)
	err := binary.Read(buf, binary.LittleEndian, &res)
	return res, err
}

func int16FromLittleEndian(num []byte) (int16, error) {
	var res int16
	num = num[:2]
	buf := bytes.NewReader(num)
	err := binary.Read(buf, binary.LittleEndian, &res)
	return res, err
}

func int8FromLittleEndian(num []byte) (int8, error) {
	var res int8
	num = num[:1]
	buf := bytes.NewReader(num)
	err := binary.Read(buf, binary.LittleEndian, &res)
	return res, err
}
