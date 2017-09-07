package main

import (
	"net"
	"time"
	"fmt"
)

func main() {

	var ip string
	var ports = []int{80, 443, 444,  65,3389}

	fmt.Print("Enter IP address to scan: ")
	fmt.Scan(&ip)

	for _, port:= range ports {
		hostIP := fmt.Sprintf("%s:%d", ip, port)

		//fmt.Println(hostIP)
		_, err := net.DialTimeout("tcp", hostIP, 10*time.Millisecond)

		if err != nil {
			fmt.Printf("Port %d is closed\n", port)
		} else {
			fmt.Printf("Port %d is open\n", port)
		}
	}
}