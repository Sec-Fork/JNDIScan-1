package util

import (
	"io/ioutil"
	"net"
	"net/http"
)

func GetLocalIPs() []string {
	var result []string
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		return result
	}
	for _, address := range addrList {
		if ip, ok := address.(*net.IPNet); ok {
			if ip.IP.To4() != nil {
				result = append(result, ip.IP.String())
			}
		}
	}
	return result
}

func GetExternalIP() string {
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(data)
}
