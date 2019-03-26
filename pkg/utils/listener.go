package utils

import (
	"fmt"
	"log"
	"net"
	"time"
)

const (
	defaultMulticastAddress = "239.0.0.0:9999"
	maxDatagramSize         = 8192
)

// func GetServerIP(done chan string) string {
func GetServerIP() string {
	ip := Listen(defaultMulticastAddress)
	if ip != "" {
		return ip
	}

	ipStr := fmt.Sprintf("%s", GetLocalIP())
	return ipStr
}

// Listen binds to the UDP address and port given and writes packets received
// from that address to a buffer which is passed to a hander
func Listen(address string) string {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal(err)
	}

	// Open up a connection
	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		log.Println(err)
	}

	conn.SetReadBuffer(maxDatagramSize)

	ch := make(chan string, 1)
	go checkUDPCast(ch, conn)

	select {
	case ip, _ := <-ch:
		return ip
	case <-time.After(1 * time.Second):
		log.Println("Timed out: no udp-cast available")
		return ""
	}
}

func checkUDPCast(ch chan string, conn *net.UDPConn) {
	buffer := make([]byte, maxDatagramSize)
	numBytes, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Println("ReadFromUDP failed:", err)
	}
	ip := string(buffer[:numBytes])
	ch <- ip
}
