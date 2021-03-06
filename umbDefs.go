package main

import (
	"fmt"
	"time"
)

const UMB_SOH byte = 1   // 01h
const UMB_STX byte = 2   // 02h
const UMB_ETX byte = 3   // 03h
const UMB_EOT byte = 4   // 04h
const UMB_VER byte = 16  // 10h Software version
const UMB_VERC byte = 16 // 10h Command version
const UMB_CLWS byte = 07 // 7h Class ID for weather station
const UMB_CLPC byte = 15 // fh Class ID for pc station
const UMB_PCID byte = 01 // 1h Device ID for pc station

type umbTypeEnum byte

const (
	umbUnsignedChar  umbTypeEnum = 16
	umbSignedChar    umbTypeEnum = 17
	umbUnsignedShort umbTypeEnum = 18
	umbSignedShort   umbTypeEnum = 19
	umbUnsignedLong  umbTypeEnum = 20
	umbSignedLong    umbTypeEnum = 21
	umbFloat         umbTypeEnum = 22
	umbDouble        umbTypeEnum = 23
	umbNull          umbTypeEnum = 0
)

type umbTypeStruc struct {
	umbName string
	code    umbTypeEnum
	bytes   byte
	maxVal  float64
	minVal  float64
}

var umbTypeMap = map[byte]umbTypeStruc{
	0:  umbTypeStruc{umbName: "NULL", code: umbNull, bytes: 0, minVal: 0, maxVal: 0},
	16: umbTypeStruc{umbName: "UNSIGNED_CHAR", code: umbUnsignedChar, bytes: 1, minVal: 0, maxVal: 255},
	17: umbTypeStruc{umbName: "SIGNED_CHAR", code: umbSignedChar, bytes: 1, minVal: -128, maxVal: 127},
	18: umbTypeStruc{umbName: "UNSIGNED_SHORT", code: umbUnsignedShort, bytes: 2, minVal: 0, maxVal: 65.535},
	19: umbTypeStruc{umbName: "SIGNED_SHORT", code: umbSignedShort, bytes: 2, minVal: -32.768, maxVal: 32.767},
	20: umbTypeStruc{umbName: "UNSIGNED_LONG", code: umbUnsignedLong, bytes: 4, minVal: 0, maxVal: 4294967295},
	21: umbTypeStruc{umbName: "SIGNED_LONG", code: umbSignedLong, bytes: 4, minVal: -2147483648, maxVal: 2147483647},
	22: umbTypeStruc{umbName: "FLOAT", code: umbFloat, bytes: 4, minVal: -3.39e+38, maxVal: 3.39e+38},
	23: umbTypeStruc{umbName: "DOUBLE", code: umbDouble, bytes: 8, minVal: -1.79e+308, maxVal: 1.79e+308},
}

type umbMeasurementTypeEnum byte

const (
	mwtCurrent umbMeasurementTypeEnum = 16 //10h - Current measurement valu
	mwtMin     umbMeasurementTypeEnum = 17 //11h - Minimum value
	mwtMax     umbMeasurementTypeEnum = 18 //12h - Maximum value
	mwtAvg     umbMeasurementTypeEnum = 19 //13h - Mean value
	mwtSum     umbMeasurementTypeEnum = 20 //14h - Sum
	mwtVct     umbMeasurementTypeEnum = 21 //15h - Vectorial mean value
)

type umbMeasurementTypeStruct struct {
	code        umbMeasurementTypeEnum
	name        string
	description string
}

var umbMeasurementTypeMap = map[byte]umbMeasurementTypeStruct{
	16: umbMeasurementTypeStruct{code: mwtCurrent, name: "MWT_CURRENT", description: "Current measurement value (10h)"},
	17: umbMeasurementTypeStruct{code: mwtMin, name: "MWT_MIN", description: "Minimum value (11h)"},
	18: umbMeasurementTypeStruct{code: mwtMax, name: "MWT_MAX", description: "Maximum value (12h)"},
	19: umbMeasurementTypeStruct{code: mwtAvg, name: "MWT_AVG", description: "Mean value (13h)"},
	20: umbMeasurementTypeStruct{code: mwtSum, name: "MWT_SUM", description: "Sum value (14h)"},
	21: umbMeasurementTypeStruct{code: mwtVct, name: "MWT_VCT", description: "Vectorial mean value"},
}

type umbStatusStruct struct {
	code        byte
	name        string
	description string
}

var umbStatusMap = map[byte]umbStatusStruct{
	0: umbStatusStruct{code: 0, name: "OK", description: "Command Succesful OK"},

	16: umbStatusStruct{code: 16, name: "UNBEK_CMD", description: "Unknown command; not supported by this device"},
	17: umbStatusStruct{code: 17, name: "UNGLTG_PARAM", description: "Invalid parameter"},
	18: umbStatusStruct{code: 18, name: "UNGLTG_HEADER", description: "Invalid header version"},
	19: umbStatusStruct{code: 19, name: "UNGLTG_VERC", description: "Invalid version of the command"},
	20: umbStatusStruct{code: 20, name: "UNGLTG_PW", description: "Invalid password for command"},

	32: umbStatusStruct{code: 32, name: "LESE_ERR", description: "Read error"},
	33: umbStatusStruct{code: 33, name: "SCHREIB_ERR", description: "Write error"},
	34: umbStatusStruct{code: 34, name: "ZU_LANG", description: "Length too great"},
	35: umbStatusStruct{code: 35, name: "UNGLTG_ADRESS", description: "Invalid address / storage location"},
	36: umbStatusStruct{code: 36, name: "UNGLTG_KANAL", description: "Invalid channel"},
	37: umbStatusStruct{code: 37, name: "UNGLTG_CMD", description: "Command not possible in this mode"},
	38: umbStatusStruct{code: 38, name: "UNBEK_CAL_CMD", description: "Unknown calibration command"},
	39: umbStatusStruct{code: 39, name: "CAL_ERROR", description: "Calibration error"},
	40: umbStatusStruct{code: 40, name: "BUSY", description: "Device not ready; e.g. initialisation / calibration running"},
	41: umbStatusStruct{code: 41, name: "LOW_VOLTAGE", description: "Undervoltage"},
	42: umbStatusStruct{code: 42, name: "HW_ERROR", description: "Hardware error"},
	43: umbStatusStruct{code: 43, name: "MEAS_ERROR", description: "Measurement error"},
	44: umbStatusStruct{code: 44, name: "INIT_ERROR", description: "Error on device initialization"},
	45: umbStatusStruct{code: 45, name: "OS_ERROR", description: "Error in operating system"},

	48: umbStatusStruct{code: 48, name: "E2_DEFAULT_KONF", description: "Configuration error, default configuration was loaded"},
	49: umbStatusStruct{code: 49, name: "E2_CAL_ERROR", description: "Calibration error / cal invalid or measurement not possible"},
	50: umbStatusStruct{code: 50, name: "E2_CRC_KONF_ERR", description: "CRC error on loading configuration; default configuration was loaded"},
	51: umbStatusStruct{code: 51, name: "E2_CRC_KAL_ERR", description: "CRC error on loading calibration; measurement not possible"},
	52: umbStatusStruct{code: 52, name: "ADJ_STEP1", description: "Calibration Step 1"},
	53: umbStatusStruct{code: 53, name: "ADJ_OK", description: "Calibration OK"},
	54: umbStatusStruct{code: 54, name: "KANAL_AUS", description: "Channel deactivated"},

	80: umbStatusStruct{code: 80, name: "VALUE_OVERFLOW", description: "Measurement variable (+offset) outside the range"},
	81: umbStatusStruct{code: 81, name: "VALUE_UNDERFLOW", description: "Measurement variable (+offset) outside the range"},
	82: umbStatusStruct{code: 82, name: "CHANNEL_OVERRANGE", description: "Measurement (physical) outside measurement range or ADC overrange "},
	83: umbStatusStruct{code: 83, name: "CHANNEL_UNDERRANGE", description: "Measurement (physical) outside measurement range or ADC overrange "},
	84: umbStatusStruct{code: 84, name: "DATA_ERROR", description: "Data error in measurement data or no valid data available"},
	85: umbStatusStruct{code: 85, name: "MEAS_UNABLE", description: "Device / sensor is unable to execute valid measurement due to ambient conditions"},

	96: umbStatusStruct{code: 96, name: "FLASH_CRC_ERR", description: "CRC-Fehler in den Flash-Daten"},
	97: umbStatusStruct{code: 97, name: "FLASH_WRITE_ERR", description: "Fehler beim Schreiben ins Flash"},
	98: umbStatusStruct{code: 98, name: "FLASH_FLOAT_ERR", description: "Flash enthält ungültige Float-Werte"},
	99: umbStatusStruct{code: 99, name: "CONV_GO_ERR", description: "Data type convert error while decoding"},

	255: umbStatusStruct{code: 255, name: "UNBEK_ERR", description: "Unknown error"},
}

type umbCmd struct {
	name       string
	code       umbCmdEnum
	cHex       string
	hasPayload bool
	desc       string
}

var cmdVersion = umbCmd{name: "VERSION", code: 32, cHex: "20h", hasPayload: false, desc: "software and hardware version"}
var cmdRead = umbCmd{name: "READ", code: 35, cHex: "23h", hasPayload: true, desc: "online Data Request"}
var cmdReset = umbCmd{name: "RESET", code: 37, cHex: "25h", hasPayload: true, desc: "software reset"}
var cmdStatus = umbCmd{name: "STATUS", code: 38, cHex: "26h", hasPayload: false, desc: "current status and/or error codes"}
var cmdSetDate = umbCmd{name: "SET_DATE", code: 39, cHex: "27h", hasPayload: true, desc: "set date & time"}
var cmdLastError = umbCmd{name: "GET_ERROR", code: 44, cHex: "2ch", hasPayload: false, desc: "last error message"}
var cmdReadMulti = umbCmd{name: "READ_MULTI", code: 47, cHex: "2fh", hasPayload: true, desc: "online Multi var Data Request"}

type umbCmdEnum byte

const (
	umbCmdVersion   umbCmdEnum = 32
	umbCmdRead      umbCmdEnum = 35
	umbCmdReset     umbCmdEnum = 37
	umbCmdStatus    umbCmdEnum = 38
	umbCmdSetDate   umbCmdEnum = 39
	umbCmdLastError umbCmdEnum = 44
	umbCmdReadMulti umbCmdEnum = 47
)

var umbCmdMap = map[byte]umbCmd{
	32: cmdVersion,
	35: cmdRead,
	37: cmdReset,
	38: cmdStatus,
	39: cmdSetDate,
	44: cmdLastError,
	45: cmdReadMulti,
	0:  umbCmd{code: 0, name: "invalid command"},
}

func (myCmd umbCmd) describe() {
	fmt.Print("Command: %s \n %s", myCmd.name, myCmd.desc)
}

type umbTelegram struct {
	cmd             umbCmd
	message         []byte
	lastRawResponse *[]byte
	lastResponse    *umbResponse
}

func (myTelegram *umbTelegram) send(deviceIP string) (umbResponse, error) {
	reply, err := sendTCPTelegram(deviceIP, myTelegram.message)
	myTelegram.lastRawResponse = &reply
	res := decodeUmbReply(reply)
	myTelegram.lastResponse = &res
	return res, err
}

type umbResponse struct {
	fromAddr    uint16
	fromDevType byte
	toAddr      uint16
	toDevType   byte
	hwVersion   string
	resLen      int
	cmdCode     umbCmdEnum
	cmdResponse umbCmd
	rawResponse []byte
	resStatus   umbStatusStruct
	isOk        bool
	readings    []umbReading
}

func (response umbResponse) describe() {
	fmt.Println("UMB Message Response: ", response.isOk)
	fmt.Printf(" Date:         %s \n", time.Now())
	fmt.Printf(" From device:  %d - Address: %d \n", response.fromDevType, response.fromAddr)
	fmt.Printf(" To device:    %d - Address: %d \n", response.toDevType, response.toAddr)
	fmt.Printf(" H/W Version:%s \n", response.hwVersion)
	fmt.Printf(" Command (%x): %s \n", response.cmdResponse.code, response.cmdResponse.name)
	fmt.Println(" Response length (bytes): ", response.resLen)
	fmt.Printf(" Response (raw): % x \n", response.rawResponse)
	fmt.Printf(" Status (%d):   %s \n", response.resStatus.code, response.resStatus.description)
	fmt.Printf(" Readings:     %d \n", len(response.readings))
	if len(response.readings) > 0 {
		fmt.Printf("  id Channel  Value \n")
	}
	for ix, el := range response.readings {
		fmt.Printf("   %d  %6d  %6.3f \n", ix+1, el.channel, el.value)
	}
}

func (response *umbResponse) decodeReadResponse() {
	response.resStatus = umbStatusMap[response.rawResponse[10]]
	subTelegram := response.rawResponse[10:]
	read := decodeReadSubtelegram(subTelegram)
	response.readings = []umbReading{read}
}

func (response *umbResponse) decodeReadMchResponse() {
	response.resStatus = umbStatusMap[response.rawResponse[10]]
	channels := response.rawResponse[11]
	var readings []umbReading
	var ix int = 13 // index of first subTelegram
	for i := 0; i < int(channels); i++ {
		subTelegram := response.rawResponse[ix:]
		read := decodeReadSubtelegram(subTelegram)
		readings = append(readings, read)
		ix = ix + int(response.rawResponse[ix-1]) + 1 //next subTelegram index
	}
	response.readings = readings
}

func decodeReadSubtelegram(subTelegram []byte) umbReading {
	read := umbReading{
		value:    0,
		channel:  uint16FromLitteEndian(subTelegram[1:3]),
		readType: umbTypeMap[subTelegram[3]],
		readTime: time.Now(),
		status:   umbStatusMap[subTelegram[0]],
	}
	switch read.readType.code {
	case umbSignedChar:
		val, _ := int8FromLittleEndian(subTelegram[4:5])
		read.value = float64(val)
	case umbSignedShort:
		val, _ := int16FromLittleEndian(subTelegram[4:6])
		read.value = float64(val)
	case umbSignedLong:
		val, _ := int32FromLittleEndian(subTelegram[4:8])
		read.value = float64(val)
	case umbUnsignedChar:
		read.value = float64(subTelegram[4])
	case umbUnsignedShort:
		read.value = float64(uint16FromLitteEndian(subTelegram[4:6]))
	case umbUnsignedLong:
		read.value = float64(uint32FromLitteEndian(subTelegram[4:8]))
	case umbFloat:
		val, _ := float32FromLittleEndian(subTelegram[4:8])
		read.value = float64(val)
	case umbDouble:
		val, _ := float32FromLittleEndian(subTelegram[4:12])
		read.value = float64(val)
	}
	return read
}

type umbReading struct {
	value    float64
	channel  uint16
	readType umbTypeStruc
	readTime time.Time
	status   umbStatusStruct
}

func buildReadChannelTelegram(deviceID uint16, channel uint16) umbTelegram {
	message := encodeHead(deviceID)

	payload := uint16ToLitteEndian(channel)
	body := encodeFrameRequest(cmdRead, payload...)
	message = append(message, body...)

	tail := encodeTail(message)
	message = append(message, tail...)
	telegram := umbTelegram{cmd: cmdRead, message: message}
	return telegram
}

func buildReadMultichannelTelegram(deviceID uint16, channels ...uint16) umbTelegram {
	message := encodeHead(deviceID)

	payload := make([]byte, 0, 2)
	payload = append(payload, byte(len(channels)))
	for _, el := range channels {
		payload = append(payload, uint16ToLitteEndian(el)...)
	}
	body := encodeFrameRequest(cmdReadMulti, payload...)
	message = append(message, body...)

	tail := encodeTail(message)
	message = append(message, tail...)
	telegram := umbTelegram{cmd: cmdReadMulti, message: message}
	return telegram
}

func buildGetLastErrorTelegram(deviceID uint16) umbTelegram {
	message := encodeHead(deviceID)

	var payload []byte
	body := encodeFrameRequest(cmdLastError, payload...)
	message = append(message, body...)

	tail := encodeTail(message)
	message = append(message, tail...)
	telegram := umbTelegram{cmd: cmdLastError, message: message}
	return telegram
}

func buildGetStatusTelegram(deviceID uint16) umbTelegram {
	message := encodeHead(deviceID)

	var payload []byte
	body := encodeFrameRequest(cmdStatus, payload...)
	message = append(message, body...)

	tail := encodeTail(message)
	message = append(message, tail...)
	telegram := umbTelegram{cmd: cmdStatus, message: message}
	return telegram
}

func buildHwSwVersionTelegram(deviceID uint16) umbTelegram {
	message := encodeHead(deviceID)

	var payload []byte
	body := encodeFrameRequest(cmdVersion, payload...)
	message = append(message, body...)

	tail := encodeTail(message)
	message = append(message, tail...)

	var telegram = umbTelegram{cmd: cmdVersion, message: message}
	return telegram
}

func buildResetTelegram(deviceID uint16) umbTelegram {
	message := encodeHead(deviceID)

	var payload byte = 0x10
	body := encodeFrameRequest(cmdReset, payload)
	message = append(message, body...)

	tail := encodeTail(message)
	message = append(message, tail...)
	telegram := umbTelegram{cmd: cmdReset, message: message}
	return telegram
}

func buildSetDateTelegram(deviceID uint16) umbTelegram {
	message := encodeHead(deviceID)

	payload := uint32ToLitteEndian(uint32(time.Now().Unix()))
	body := encodeFrameRequest(cmdSetDate, payload...)
	message = append(message, body...)

	tail := encodeTail(message)
	message = append(message, tail...)
	telegram := umbTelegram{cmd: cmdSetDate, message: message}
	return telegram
}

func decodeUmbReply(rawResponse []byte) umbResponse {
	if len(rawResponse) < 10 {
		return umbResponse{isOk: false}
	}
	devFrom, addrFrom := decodeUmbAddr(rawResponse[4:6])
	devTo, addrTo := decodeUmbAddr(rawResponse[2:4])
	cmdStr := umbCmdMap[rawResponse[8]]
	res := umbResponse{
		fromAddr:    addrFrom,
		fromDevType: devFrom,
		toAddr:      addrTo,
		toDevType:   devTo,
		hwVersion:   fmt.Sprintf("%6.2f", float32(rawResponse[1])/10),
		cmdCode:     umbCmdEnum(rawResponse[8]),
		cmdResponse: cmdStr,
		resLen:      len(rawResponse),
		rawResponse: rawResponse,
		isOk:        true,
	}
	if len(rawResponse) == 10 {
		return res
	}

	switch res.cmdCode {
	case umbCmdVersion:
		fmt.Println("get version params")
	case umbCmdRead:
		res.decodeReadResponse()
	case umbCmdReset:
		fmt.Println("reset device params")
	case umbCmdReadMulti:
		res.decodeReadMchResponse()
	default:
		fmt.Println("Error: can't complete response decoding command not found")
	}

	return res
}
