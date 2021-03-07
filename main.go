package main

import (
	"fmt"

	umb "github.com/rwilmar/umbGateway/umbBinary"
)

func main() {

	//config device id and classes, default: pc#1, ws10,
	//umb.ConfigDevicesClasses(15, 7)
	//umb.ConfigOriginDevId(1)

	serverAddr := "192.168.1.36:9750"

	fmt.Printf("\nSending Read Online Channel...\n")
	telegram := umb.BuildReadChannelTelegram(1, 100)
	fmt.Printf("Request: %x \n", telegram.Message)
	ans, _ := telegram.Send(serverAddr)
	fmt.Println("reply size:", ans.ResLen)
	fmt.Printf("raw response %x \n", *telegram.LastRawResponse)
	fmt.Printf("raw response %x \n", ans.RawResponse)
	ans.Describe()

	fmt.Printf("\nSending Read Multihannel...\n")
	telegram2 := umb.BuildReadMultichannelTelegram(1, 100, 200, 4061, 4705, 10000)
	reply, _ := telegram2.Send(serverAddr)
	reply.Describe()

	fmt.Printf("\nSending Test ...\n")
	tel2 := umb.BuildHwSwVersionTelegram(1)
	reply, _ = tel2.Send(serverAddr)
	reply.Describe()

	/*for ix, el := range tel2 {
		fmt.Printf(" ix[%d]: %x \n", ix, el)
	}
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Println("Pi : ", pi)
	*/

}
