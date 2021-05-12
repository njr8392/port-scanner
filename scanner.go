package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"
)

var open int
var closed int

func main() {
	ip := flag.String("ip", "", "IP address to scan")
	portRange := flag.Int("range", 1024, "Scan up to n ports")
	flag.Parse()
	for i := 1; i <= *portRange; i++ {
		scan(*ip, i)
	}
	fmt.Printf("Scanned %d ports.  %d are open and %d are closed\n", *portRange, open, closed)
}

func scan(ip string, port int) {
	portStr := strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", ip+":"+portStr, 60*time.Millisecond)
	if err == nil {
		conn.Close()
		fmt.Printf("Port %s open\n", portStr)
		open++
		return
	}
	closed++

}
