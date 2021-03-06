package main

import (
	"fmt"
	"net"
)

/**
Function to calc CRC-CCITT
Shall be applied to all bytes from SOH to ETX inclusive
Polynomial: 1021h (LSB first mode) Start value: FFFFh
**/
func calcCrc(crcBuff uint16, input byte) uint16 {
	var x16 uint16 = 0x0000
	for i := 0; i < 8; i++ {
		if ((crcBuff & 0x0001) ^ uint16(input&0x01)) == 1 {
			x16 = 0x8408
		} else {
			x16 = 0x0000
		}
		crcBuff = crcBuff >> 1
		crcBuff = crcBuff ^ x16
		input = input >> 1
	}
	return crcBuff
}

func sendTCPTelegram(serverAddr string, telegram []byte) ([]byte, error) {
	reply := make([]byte, 0, 1024) // reply buffer
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println(err)
		return reply, err
	}

	_, err = conn.Write(telegram)
	if err != nil {
		fmt.Println(err)
		return reply, err
	}

	buff := make([]byte, 10) // using small buffer
	for {
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		reply = append(reply, buff[:n]...)
		if n < 10 { // bytes read < buffer = message ended
			break
		}
	}
	conn.Close()

	if err != nil {
		fmt.Println(err)
		return reply, err
	}

	return reply, nil
}

func encodeUmbAddr(deviceType byte, deviceID uint16) []byte {
	var toAddr uint16 = uint16(deviceType)
	toAddr = toAddr << 12
	toAddr = toAddr + deviceID
	return uint16ToLitteEndian(toAddr)
}

func decodeUmbAddr(addrLE []byte) (devType byte, devAddress uint16) {
	addr := uint16FromLitteEndian(addrLE)
	dev := (addr & 0xF000)
	dev = dev >> 12
	addr = (addr & 0x0FFF)
	return uint8(dev), addr
}

/**
Build/Encode generic telegram header for message
**/
func encodeHead(deviceID uint16) []byte {
	toAddrLE := encodeUmbAddr(UMB_CLWS, deviceID)
	// var toAddr uint16 = uint16(UMB_CLWS)
	// toAddr = toAddr << 12
	// toAddr = toAddr + deviceID
	// toAddrLE := uint16ToLitteEndian(toAddr)

	var fromAddr uint16 = uint16(UMB_CLPC)
	fromAddr = fromAddr << 12
	fromAddr = fromAddr + uint16(UMB_PCID)
	fromAddrLE := uint16ToLitteEndian(fromAddr)

	header := make([]byte, 0, 6)
	header = append(header, UMB_SOH)
	header = append(header, UMB_VER)
	header = append(header, toAddrLE...)
	header = append(header, fromAddrLE...)

	return header
}

func encodeTail(telegram []byte) []byte {
	var crc uint16 = 0xffff
	for _, el := range telegram {
		crc = calcCrc(crc, el)
	}
	res := uint16ToLitteEndian(crc)
	return append(res, UMB_EOT)
}

/***
Build/Encode generic telegram body for message
***/
func encodeFrameRequest(command umbCmd, payload ...byte) []byte {
	frame := make([]byte, 0, 5)
	frame = append(frame, 0)
	frame = append(frame, UMB_STX)
	frame = append(frame, byte(command.code))
	frame = append(frame, UMB_VERC)
	if len(payload) > 0 && command.hasPayload {
		frame = append(frame, payload...)
	}
	frame = append(frame, UMB_ETX)
	frame[0] = 2 + byte(len(payload))
	return frame
}

func decodeReading(payload []byte) {

}
