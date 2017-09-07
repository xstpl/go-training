package main

import (
	"net"
	"time"
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func main() {

	var ip string

	fmt.Print("Enter IP address to scan: ")
	fmt.Scan(&ip)
	wg.Add(65535)
	for port:=1;port<65535;port++ {
		go scan(ip, port)
	}
	wg.Wait()
}

func scan(ip string, port int) {

	hostIP := fmt.Sprintf("%s:%d", ip, port)

	//fmt.Println(hostIP)
	_, err := net.DialTimeout("tcp", hostIP, 10*time.Millisecond)

	if err != nil {
		//fmt.Printf("Port %d is closed\n", port)
	} else {

		fmt.Printf("Port %d is open\n", port)
	}

	wg.Done()
}
