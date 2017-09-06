package main

import (
	"net"
	"time"
	"fmt"
)

func main() {

	var ip string
	var port string

	fmt.Print("Enter IP address to scan: ")
	fmt.Scan(&ip)

	fmt.Print("Enter port to scan: ")
	fmt.Scan(&port)

	hostIP:= fmt.Sprintf("%s:%s", ip, port)

	fmt.Println(hostIP)

	_ , err := net.DialTimeout("tcp", hostIP, 10 * time.Millisecond)

	if err != nil {
		fmt.Printf("Port %s is closed", port)
	} else {
		fmt.Printf("Port %s is open", port)
	}
}