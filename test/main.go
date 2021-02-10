package main

import (
	"github.com/OpenIoTHub/service-register/nettool"
	"log"
	"time"
)

var mdnsSServiceBaseInfo = map[string]string{
	"name":                 "OpenIoTHub服务",
	"model":                "com.iotserv.services.web",
	"author":               "Farry",
	"email":                "newfarry@126.com",
	"home-page":            "https://github.com/OpenIoTHub",
	"firmware-respository": "https://github.com/iotdevice",
	"firmware-version":     "1.0",
	//	mac
	//	id
}

func main() {
	_, err := nettool.RegistermDNSService(mdnsSServiceBaseInfo, 80)
	if err != nil {
		log.Println(err)
		return
	}
	time.Sleep(time.Second * 10)
}
