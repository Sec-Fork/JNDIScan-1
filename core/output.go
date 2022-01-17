package core

import (
	"fmt"
	"github.com/EmYiQing/JNDIScan/config"
	"github.com/EmYiQing/JNDIScan/util"
	"strconv"
	"time"
)

func PrintPayload() {
	time.Sleep(time.Second * 3)
	ipList := util.GetLocalIPs()
	exIP := util.GetExternalIP()
	if exIP != "" {
		ipList = append(ipList, exIP)
	}
	port := config.Port
	borderLen := 23
	max := 0
	for i := 0; i < len(ipList); i++ {
		if len(ipList[i]) > max {
			max = len(ipList[i])
		}
	}
	borderLen += max + len(strconv.Itoa(port)) + 2
	fmt.Print("|")
	for i := 0; i < borderLen; i++ {
		fmt.Print("-")
	}
	fmt.Println("|")

	for i := 0; i < len(ipList); i++ {
		tempLen := 18
		fmt.Print("|--Payload: ldap://")
		tempLen += len(ipList[i])
		fmt.Print(ipList[i])
		tempLen += 1
		fmt.Print(":")
		tempLen += len(strconv.Itoa(port))
		fmt.Print(strconv.Itoa(port))
		num := borderLen - tempLen
		for j := 0; j < num; j++ {
			fmt.Print("-")
		}
		fmt.Println("|")
	}

	for i := 0; i < len(ipList); i++ {
		tempLen := 17
		fmt.Print("|--Payload: rmi://")
		tempLen += len(ipList[i])
		fmt.Print(ipList[i])
		tempLen += 1
		fmt.Print(":")
		tempLen += len(strconv.Itoa(port))
		fmt.Print(strconv.Itoa(port))
		tempLen += 4
		fmt.Print("/xxx")
		num := borderLen - tempLen
		for j := 0; j < num; j++ {
			fmt.Print("-")
		}
		fmt.Println("|")
	}

	fmt.Print("|")
	for i := 0; i < borderLen; i++ {
		fmt.Print("-")
	}
	fmt.Println("|")
}
