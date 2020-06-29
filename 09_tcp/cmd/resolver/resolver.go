package main

import (
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveIPAddr("ip", "tinkoff.ru")
	if err != nil {
		return
	}
	ip := addr.IP
	log.Println(ip)
}


