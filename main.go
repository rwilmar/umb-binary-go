package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	type Grades int
	const ( //enums
		A Grades = iota
		B
		C
		D
	)
	var i int = 120
	//y := 2
	var m *int
	m = &i
	fmt.Printf("Res %d \n", *m)
	fmt.Printf("hello world!\n")

	sli := make([]int, 0, 7)
	sli2 := []int{34, 55, 1}
	sli = append(sli, 12)
	sli = append(sli, sli2...)
	sli = append(sli, 401)

	for ix, el := range sli {
		fmt.Printf("slice [%d] - %d \n", ix, el)
	}

	src := []byte("48656c6c6f20476f7068657221")
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", dst[:n])
	fmt.Printf("map read_cmd:%d\n", cmdRead.code)

	serverAddr := "192.168.1.37:9750"

	fmt.Printf("\nSending Read Online Channel...\n")
	telegram := buildReadChannelTelegram(1, 100)
	fmt.Printf("Request: %x \n", telegram.message)
	ans, _ := telegram.send(serverAddr)
	fmt.Println("reply size:", ans.resLen)
	fmt.Printf("raw response %x \n", *telegram.lastRawResponse)
	fmt.Printf("raw response %x \n", ans.rawResponse)
	ans.describe()

	fmt.Printf("\nSending Read Multihannel...\n")
	telegram2 := buildReadMultichannelTelegram(1, 100, 4061, 10000)
	reply, _ := telegram2.send(serverAddr)
	reply.describe()

	fmt.Printf("\nSending Test ...\n")
	tel2 := buildHwSwVersionTelegram(1)
	reply, _ = tel2.send(serverAddr)
	reply.describe()

	/*for ix, el := range tel2 {
		fmt.Printf(" ix[%d]: %x \n", ix, el)
	}*/

	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	pi, err := float64FromLittleEndian(b)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Println("Pi : ", pi)

}
