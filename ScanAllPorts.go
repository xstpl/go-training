package main

import (
	"net"
	"time"
	"fmt"
)

func main() {

	var ip string

	fmt.Print("Enter IP address to scan: ")
	fmt.Scan(&ip)

	for port:=1;port<65535;port++ {
		hostIP := fmt.Sprintf("%s:%d", ip, port)

		//fmt.Println(hostIP)
		_, err := net.DialTimeout("tcp", hostIP, 10*time.Millisecond)

		if err != nil {
			//fmt.Printf("Port %d is closed\n", port)
		} else {
			fmt.Printf("Port %d is open\n", port)
		}
	}
}
