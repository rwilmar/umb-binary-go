package umbBinary

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

var umbCommDestinyDeviceClass umbDeviceClassEnum = umbDevWeatherStat
var umbCommOriginDeviceClass umbDeviceClassEnum = umbDeviceMaster
var umbCommOriginDeviceID uint16 = 01 //default 1h Device ID for pc station

func ConfigDevicesClasses(originClass umbDeviceClassEnum, destinyClass umbDeviceClassEnum) {
	umbCommOriginDeviceClass = originClass
	umbCommDestinyDeviceClass = destinyClass
}

func ConfigOriginDevId(deviceId uint16) {
	umbCommOriginDeviceID = deviceId
}

type umbDeviceClassEnum byte

const (
	umbDevBroadCast    umbDeviceClassEnum = 0
	umbDevRoadSensor   umbDeviceClassEnum = 1
	umbDevRainSensor   umbDeviceClassEnum = 2
	umbDevVisibSensor  umbDeviceClassEnum = 3
	umbDevAcRoadSensor umbDeviceClassEnum = 4
	umbDevNIRoadSensor umbDeviceClassEnum = 5
	umbDevUniversalTx  umbDeviceClassEnum = 6
	umbDevWeatherStat  umbDeviceClassEnum = 7
	umbDevWindSensor   umbDeviceClassEnum = 8
	umbDeviceMaster    umbDeviceClassEnum = 15
)

type UmbDeviceClassStruct struct {
	Name    string
	UmbCode umbDeviceClassEnum
}

var umbDeviceMap = map[byte]UmbDeviceClassStruct{
	1:  {Name: "Broadcast", UmbCode: 1},
	2:  {Name: "Road sensor (IRS31-UMB)", UmbCode: 2},
	3:  {Name: "Visibility sensor (VS20-UMB)", UmbCode: 3},
	4:  {Name: "Active road sensor (ARS31-UMB)", UmbCode: 4},
	5:  {Name: "Non invasive road sensor (NIRS31-UMB)", UmbCode: 5},
	6:  {Name: "Universal measurement transmitter (ANACON)", UmbCode: 6},
	7:  {Name: "Compact weather station (WS family)", UmbCode: 7},
	8:  {Name: "Wind sensor class (VENTUS / V200A)", UmbCode: 8},
	15: {Name: "Master / control devices", UmbCode: 15},
}

type UmbTypeEnum byte

const (
	umbUnsignedChar  UmbTypeEnum = 16
	umbSignedChar    UmbTypeEnum = 17
	umbUnsignedShort UmbTypeEnum = 18
	umbSignedShort   UmbTypeEnum = 19
	umbUnsignedLong  UmbTypeEnum = 20
	umbSignedLong    UmbTypeEnum = 21
	umbFloat         UmbTypeEnum = 22
	umbDouble        UmbTypeEnum = 23
	umbNull          UmbTypeEnum = 0
)

type UmbTypeStruc struct {
	UmbName string
	Code    UmbTypeEnum
	Bytes   byte
	MaxVal  float64
	MinVal  float64
}

var umbTypeMap = map[byte]UmbTypeStruc{
	0:  {UmbName: "NULL", Code: umbNull, Bytes: 0, MinVal: 0, MaxVal: 0},
	16: {UmbName: "UNSIGNED_CHAR", Code: umbUnsignedChar, Bytes: 1, MinVal: 0, MaxVal: 255},
	17: {UmbName: "SIGNED_CHAR", Code: umbSignedChar, Bytes: 1, MinVal: -128, MaxVal: 127},
	18: {UmbName: "UNSIGNED_SHORT", Code: umbUnsignedShort, Bytes: 2, MinVal: 0, MaxVal: 65.535},
	19: {UmbName: "SIGNED_SHORT", Code: umbSignedShort, Bytes: 2, MinVal: -32.768, MaxVal: 32.767},
	20: {UmbName: "UNSIGNED_LONG", Code: umbUnsignedLong, Bytes: 4, MinVal: 0, MaxVal: 4294967295},
	21: {UmbName: "SIGNED_LONG", Code: umbSignedLong, Bytes: 4, MinVal: -2147483648, MaxVal: 2147483647},
	22: {UmbName: "FLOAT", Code: umbFloat, Bytes: 4, MinVal: -3.39e+38, MaxVal: 3.39e+38},
	23: {UmbName: "DOUBLE", Code: umbDouble, Bytes: 8, MinVal: -1.79e+308, MaxVal: 1.79e+308},
}

type UmbMeasurementTypeEnum byte

const (
	mwtCurrent UmbMeasurementTypeEnum = 16 //10h - Current measurement valu
	mwtMin     UmbMeasurementTypeEnum = 17 //11h - Minimum value
	mwtMax     UmbMeasurementTypeEnum = 18 //12h - Maximum value
	mwtAvg     UmbMeasurementTypeEnum = 19 //13h - Mean value
	mwtSum     UmbMeasurementTypeEnum = 20 //14h - Sum
	mwtVct     UmbMeasurementTypeEnum = 21 //15h - Vectorial mean value
)

type UmbMeasurementTypeStruct struct {
	Code        UmbMeasurementTypeEnum
	Name        string
	Description string
}

var umbMeasurementTypeMap = map[byte]UmbMeasurementTypeStruct{
	16: {Code: mwtCurrent, Name: "MWT_CURRENT", Description: "Current measurement value (10h)"},
	17: {Code: mwtMin, Name: "MWT_MIN", Description: "Minimum value (11h)"},
	18: {Code: mwtMax, Name: "MWT_MAX", Description: "Maximum value (12h)"},
	19: {Code: mwtAvg, Name: "MWT_AVG", Description: "Mean value (13h)"},
	20: {Code: mwtSum, Name: "MWT_SUM", Description: "Sum value (14h)"},
	21: {Code: mwtVct, Name: "MWT_VCT", Description: "Vectorial mean value"},
}

type UmbStatusStruct struct {
	Code        byte
	Name        string
	Description string
}

var umbStatusMap = map[byte]UmbStatusStruct{
	0: {Code: 0, Name: "OK", Description: "Command Succesful OK"},

	16: {Code: 16, Name: "UNBEK_CMD", Description: "Unknown command; not supported by this device"},
	17: {Code: 17, Name: "UNGLTG_PARAM", Description: "Invalid parameter"},
	18: {Code: 18, Name: "UNGLTG_HEADER", Description: "Invalid header version"},
	19: {Code: 19, Name: "UNGLTG_VERC", Description: "Invalid version of the command"},
	20: {Code: 20, Name: "UNGLTG_PW", Description: "Invalid password for command"},

	32: {Code: 32, Name: "LESE_ERR", Description: "Read error"},
	33: {Code: 33, Name: "SCHREIB_ERR", Description: "Write error"},
	34: {Code: 34, Name: "ZU_LANG", Description: "Length too great"},
	35: {Code: 35, Name: "UNGLTG_ADRESS", Description: "Invalid address / storage location"},
	36: {Code: 36, Name: "UNGLTG_KANAL", Description: "Invalid channel"},
	37: {Code: 37, Name: "UNGLTG_CMD", Description: "Command not possible in this mode"},
	38: {Code: 38, Name: "UNBEK_CAL_CMD", Description: "Unknown calibration command"},
	39: {Code: 39, Name: "CAL_ERROR", Description: "Calibration error"},
	40: {Code: 40, Name: "BUSY", Description: "Device not ready; e.g. initialisation / calibration running"},
	41: {Code: 41, Name: "LOW_VOLTAGE", Description: "Undervoltage"},
	42: {Code: 42, Name: "HW_ERROR", Description: "Hardware error"},
	43: {Code: 43, Name: "MEAS_ERROR", Description: "Measurement error"},
	44: {Code: 44, Name: "INIT_ERROR", Description: "Error on device initialization"},
	45: {Code: 45, Name: "OS_ERROR", Description: "Error in operating system"},

	48: {Code: 48, Name: "E2_DEFAULT_KONF", Description: "Configuration error, default configuration was loaded"},
	49: {Code: 49, Name: "E2_CAL_ERROR", Description: "Calibration error / cal invalid or measurement not possible"},
	50: {Code: 50, Name: "E2_CRC_KONF_ERR", Description: "CRC error on loading configuration; default configuration was loaded"},
	51: {Code: 51, Name: "E2_CRC_KAL_ERR", Description: "CRC error on loading calibration; measurement not possible"},
	52: {Code: 52, Name: "ADJ_STEP1", Description: "Calibration Step 1"},
	53: {Code: 53, Name: "ADJ_OK", Description: "Calibration OK"},
	54: {Code: 54, Name: "KANAL_AUS", Description: "Channel deactivated"},

	80: {Code: 80, Name: "VALUE_OVERFLOW", Description: "Measurement variable (+offset) outside the range"},
	81: {Code: 81, Name: "VALUE_UNDERFLOW", Description: "Measurement variable (+offset) outside the range"},
	82: {Code: 82, Name: "CHANNEL_OVERRANGE", Description: "Measurement (physical) outside measurement range or ADC overrange "},
	83: {Code: 83, Name: "CHANNEL_UNDERRANGE", Description: "Measurement (physical) outside measurement range or ADC overrange "},
	84: {Code: 84, Name: "DATA_ERROR", Description: "Data error in measurement data or no valid data available"},
	85: {Code: 85, Name: "MEAS_UNABLE", Description: "Device / sensor is unable to execute valid measurement due to ambient conditions"},

	96:  {Code: 96, Name: "FLASH_CRC_ERR", Description: "CRC-Fehler in den Flash-Daten"},
	97:  {Code: 97, Name: "FLASH_WRITE_ERR", Description: "Fehler beim Schreiben ins Flash"},
	98:  {Code: 98, Name: "FLASH_FLOAT_ERR", Description: "Flash enthält ungültige Float-Werte"},
	99:  {Code: 99, Name: "CONV_GO_ERR", Description: "Data type convert error while decoding"},
	100: {Code: 99, Name: "COMM_GO_ERR", Description: "Communications error, device offline?"},

	255: {Code: 255, Name: "UNBEK_ERR", Description: "Unknown error"},
}

type UmbCmd struct {
	Name       string
	Code       UmbCmdEnum
	cHex       string
	hasPayload bool
	Desc       string
}

var cmdVersion = UmbCmd{Name: "VERSION", Code: 32, cHex: "20h", hasPayload: false, Desc: "software and hardware version"}
var cmdRead = UmbCmd{Name: "READ", Code: 35, cHex: "23h", hasPayload: true, Desc: "online Data Request"}
var cmdReset = UmbCmd{Name: "RESET", Code: 37, cHex: "25h", hasPayload: true, Desc: "software reset"}
var cmdStatus = UmbCmd{Name: "STATUS", Code: 38, cHex: "26h", hasPayload: false, Desc: "current status and/or error codes"}
var cmdSetDate = UmbCmd{Name: "SET_DATE", Code: 39, cHex: "27h", hasPayload: true, Desc: "set date & time"}
var cmdLastError = UmbCmd{Name: "GET_ERROR", Code: 44, cHex: "2ch", hasPayload: false, Desc: "last error message"}
var cmdReadMulti = UmbCmd{Name: "READ_MULTI", Code: 47, cHex: "2fh", hasPayload: true, Desc: "online Multi var Data Request"}

type UmbCmdEnum byte

const (
	UmbCmdVersion   UmbCmdEnum = 32
	UmbCmdRead      UmbCmdEnum = 35
	UmbCmdReset     UmbCmdEnum = 37
	UmbCmdStatus    UmbCmdEnum = 38
	UmbCmdSetDate   UmbCmdEnum = 39
	UmbCmdLastError UmbCmdEnum = 44
	UmbCmdReadMulti UmbCmdEnum = 47
)

var UmbCmdMap = map[byte]UmbCmd{
	32: cmdVersion,
	35: cmdRead,
	37: cmdReset,
	38: cmdStatus,
	39: cmdSetDate,
	44: cmdLastError,
	45: cmdReadMulti,
	0:  {Code: 0, Name: "invalid command"},
}

func (myCmd UmbCmd) Describe() {
	fmt.Print("Command: %s \n %s", myCmd.Name, myCmd.Desc)
}

type UmbTelegram struct {
	Cmd             UmbCmd
	Message         []byte
	LastRawResponse *[]byte
	LastResponse    *UmbResponse
}

func (myTelegram *UmbTelegram) Send(deviceIP string) (UmbResponse, error) {
	reply, err := sendTCPTelegram(deviceIP, myTelegram.Message)
	myTelegram.LastRawResponse = &reply
	res := decodeUmbReply(reply)
	myTelegram.LastResponse = &res
	return res, err
}

type UmbResponse struct {
	FromAddr    uint16
	FromDevType UmbDeviceClassStruct
	ToAddr      uint16
	ToDevType   UmbDeviceClassStruct
	HwVersion   string
	SwVersion   string
	ResLen      int
	CmdCode     UmbCmdEnum
	CmdResponse UmbCmd
	RawResponse []byte
	ResStatus   UmbStatusStruct
	IsOk        bool
	Readings    []UmbReading
	LastError   UmbStatusStruct
}

func (response UmbResponse) Describe() {
	fmt.Println("UMB Message Response: ", response.IsOk)
	fmt.Printf(" Date:         %s \n", time.Now())
	fmt.Printf(" From device:  [%d] %s - Address: %d \n", response.FromDevType.UmbCode, response.FromDevType.Name, response.FromAddr)
	fmt.Printf(" To device:    [%d] %s - Address: %d \n", response.ToDevType.UmbCode, response.ToDevType.Name, response.ToAddr)
	fmt.Printf(" H/W Version:%s \n", response.HwVersion)
	fmt.Printf(" S/W Version:%s \n", response.SwVersion)
	fmt.Printf(" Command (%x): %s \n", response.CmdResponse.Code, response.CmdResponse.Name)
	fmt.Println(" Response length (bytes): ", response.ResLen)
	fmt.Printf(" Response (raw): % x \n", response.RawResponse)
	fmt.Printf(" Status (%d):   %s \n", response.ResStatus.Code, response.ResStatus.Description)
	fmt.Printf(" Last Error (%d):   %s \n", response.LastError.Code, response.LastError.Description)
	fmt.Printf(" Readings:     %d \n", len(response.Readings))
	if len(response.Readings) > 0 {
		fmt.Printf("  id Channel  Value \n")
	}
	for ix, el := range response.Readings {
		fmt.Printf("   %d  %6d  %6.3f \n", ix+1, el.Channel, el.Value)
	}
}

func (response *UmbResponse) decodeLastErrorResponse() {
	response.ResStatus = umbStatusMap[response.RawResponse[10]]
	if len(response.RawResponse) > 11 { // last error reported
		response.LastError = umbStatusMap[response.RawResponse[12]]
	}
}

func (response *UmbResponse) decodeReadResponse() {
	response.ResStatus = umbStatusMap[response.RawResponse[10]]
	subTelegram := response.RawResponse[10:]
	read := decodeReadSubtelegram(subTelegram)
	response.Readings = []UmbReading{read}
}

func (response *UmbResponse) decodeReadMchResponse() {
	response.ResStatus = umbStatusMap[response.RawResponse[10]]
	channels := response.RawResponse[11]
	var readings []UmbReading
	var ix int = 13 // index of first subTelegram
	for i := 0; i < int(channels); i++ {
		subTelegram := response.RawResponse[ix:]
		read := decodeReadSubtelegram(subTelegram)
		readings = append(readings, read)
		ix = ix + int(response.RawResponse[ix-1]) + 1 //next subTelegram index
	}
	response.Readings = readings
}

func (response *UmbResponse) decodeResetResponse() {
	response.ResStatus = umbStatusMap[response.RawResponse[10]]
	//no additional info
}

func (response *UmbResponse) decodeSetTimeResponse() {
	response.ResStatus = umbStatusMap[response.RawResponse[10]]
	//no additional info
}

func (response *UmbResponse) decodeStatusResponse() {
	response.ResStatus = umbStatusMap[response.RawResponse[10]]
	if len(response.RawResponse) >= 12 { // contains status info
		response.ResStatus = umbStatusMap[response.RawResponse[12]]
	}
}

func (response *UmbResponse) decodeVersionResponse() {
	response.ResStatus = umbStatusMap[response.RawResponse[10]]
	if len(response.RawResponse) >= 13 { // payload contains info
		response.HwVersion = fmt.Sprintf("%6.2f", float32(response.RawResponse[12])/10)
		response.SwVersion = fmt.Sprintf("%6.2f", float32(response.RawResponse[13])/10)
	}
	//no additional info
}

func decodeReadSubtelegram(subTelegram []byte) UmbReading {
	read := UmbReading{
		Value:    0,
		Channel:  uint16FromLitteEndian(subTelegram[1:3]),
		ReadType: umbTypeMap[subTelegram[3]],
		ReadTime: time.Now(),
		Status:   umbStatusMap[subTelegram[0]],
	}
	switch read.ReadType.Code {
	case umbSignedChar:
		val, _ := int8FromLittleEndian(subTelegram[4:5])
		read.Value = float64(val)
	case umbSignedShort:
		val, _ := int16FromLittleEndian(subTelegram[4:6])
		read.Value = float64(val)
	case umbSignedLong:
		val, _ := int32FromLittleEndian(subTelegram[4:8])
		read.Value = float64(val)
	case umbUnsignedChar:
		read.Value = float64(subTelegram[4])
	case umbUnsignedShort:
		read.Value = float64(uint16FromLitteEndian(subTelegram[4:6]))
	case umbUnsignedLong:
		read.Value = float64(uint32FromLitteEndian(subTelegram[4:8]))
	case umbFloat:
		val, _ := float32FromLittleEndian(subTelegram[4:8])
		read.Value = float64(val)
	case umbDouble:
		val, _ := float32FromLittleEndian(subTelegram[4:12])
		read.Value = float64(val)
	}
	return read
}

type UmbReading struct {
	Value    float64
	Channel  uint16
	ReadType UmbTypeStruc
	ReadTime time.Time
	Status   UmbStatusStruct
}

func BuildReadChannelTelegram(deviceID uint16, channel uint16) UmbTelegram {
	message := encodeHead(deviceID)

	payload := uint16ToLitteEndian(channel)
	body := encodeFrameRequest(cmdRead, payload...)
	message = append(message, body...)

	tail := encodeTail(message)
	message = append(message, tail...)
	telegram := UmbTelegram{Cmd: cmdRead, Message: message}
	return telegram
}

func BuildReadMultichannelTelegram(deviceID uint16, channels ...uint16) UmbTelegram {
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
	telegram := UmbTelegram{Cmd: cmdReadMulti, Message: message}
	return telegram
}

func BuildGetLastErrorTelegram(deviceID uint16) UmbTelegram {
	message := encodeHead(deviceID)

	var payload []byte
	body := encodeFrameRequest(cmdLastError, payload...)
	message = append(message, body...)

	tail := encodeTail(message)
	message = append(message, tail...)
	telegram := UmbTelegram{Cmd: cmdLastError, Message: message}
	return telegram
}

func BuildGetStatusTelegram(deviceID uint16) UmbTelegram {
	message := encodeHead(deviceID)

	var payload []byte
	body := encodeFrameRequest(cmdStatus, payload...)
	message = append(message, body...)

	tail := encodeTail(message)
	message = append(message, tail...)
	telegram := UmbTelegram{Cmd: cmdStatus, Message: message}
	return telegram
}

func BuildHwSwVersionTelegram(deviceID uint16) UmbTelegram {
	message := encodeHead(deviceID)

	var payload []byte
	body := encodeFrameRequest(cmdVersion, payload...)
	message = append(message, body...)

	tail := encodeTail(message)
	message = append(message, tail...)

	var telegram = UmbTelegram{Cmd: cmdVersion, Message: message}
	return telegram
}

func BuildResetTelegram(deviceID uint16) UmbTelegram {
	message := encodeHead(deviceID)

	var payload byte = 0x10
	body := encodeFrameRequest(cmdReset, payload)
	message = append(message, body...)

	tail := encodeTail(message)
	message = append(message, tail...)
	telegram := UmbTelegram{Cmd: cmdReset, Message: message}
	return telegram
}

func BuildSetDateTelegram(deviceID uint16) UmbTelegram {
	message := encodeHead(deviceID)

	payload := uint32ToLitteEndian(uint32(time.Now().Unix()))
	body := encodeFrameRequest(cmdSetDate, payload...)
	message = append(message, body...)

	tail := encodeTail(message)
	message = append(message, tail...)
	telegram := UmbTelegram{Cmd: cmdSetDate, Message: message}
	return telegram
}

func decodeUmbReply(rawResponse []byte) UmbResponse {
	if len(rawResponse) < 10 { // no message available
		return UmbResponse{
			IsOk:      false,
			ResStatus: umbStatusMap[100],
		}
	}
	devFrom, addrFrom := decodeUmbAddr(rawResponse[4:6])
	devTo, addrTo := decodeUmbAddr(rawResponse[2:4])
	cmdStr := UmbCmdMap[rawResponse[8]]
	res := UmbResponse{
		FromAddr:    addrFrom,
		FromDevType: devFrom,
		ToAddr:      addrTo,
		ToDevType:   devTo,
		HwVersion:   fmt.Sprintf("%6.2f", float32(rawResponse[1])/10),
		CmdCode:     UmbCmdEnum(rawResponse[8]),
		CmdResponse: cmdStr,
		ResLen:      len(rawResponse),
		RawResponse: rawResponse,
		IsOk:        true,
	}
	if len(rawResponse) == 10 { // status available, but message empty
		res.ResStatus = umbStatusMap[rawResponse[10]]
		return res
	}

	switch res.CmdCode {
	case UmbCmdLastError:
		res.decodeLastErrorResponse()
	case UmbCmdRead:
		res.decodeReadResponse()
	case UmbCmdReadMulti:
		res.decodeReadMchResponse()
	case UmbCmdReset:
		res.decodeResetResponse()
	case UmbCmdSetDate:
		res.decodeSetTimeResponse()
	case UmbCmdStatus:
		res.decodeStatusResponse()
	case UmbCmdVersion:
		res.decodeVersionResponse()
	default:
		fmt.Println("Error: can't complete response decoding command not found")
	}

	return res
}
