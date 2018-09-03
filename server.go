package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func checkError(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
		return
	}
}

func listenToB(c *net.UDPConn, m map[*net.UDPAddr]int, q chan struct{}) {
	buffer := make([]byte, 1024)
	_, remoteAddr, err := 0, new(net.UDPAddr), error(nil)
	for err == nil {
		_, remoteAddr, err = c.ReadFromUDP(buffer)
		if string(buffer) == "CONNECT" {
			m[remoteAddr] = 1
		} else if string(buffer) == "DISCONNECT" {
			delete(m, remoteAddr)
		}
		fmt.Println("from", remoteAddr, "buffer was read")
	}

	fmt.Println("listener failed - ", err)
	q <- struct{}{}
}

func listenToA(c *net.UDPConn, q chan struct{}) {
	buffer := make([]byte, 1024)
	_, remoteAddr, err := 0, new(net.UDPAddr), error(nil)
	for err == nil {
		_, remoteAddr, err = c.ReadFromUDP(buffer)
		fmt.Println("from", remoteAddr, "buffer was read")
	}
	fmt.Println("listener failed - ", err)
	q <- struct{}{}
}

func main() {

	if len(os.Args) == 2 {
		fmt.Println("Please provide a Pin and Pout")
		return
	}

	pIn := ":" + os.Args[1]
	pOut := ":" + os.Args[2]

	serverAddrIn, err := net.ResolveUDPAddr("udp", pIn)
	checkError(err)

	connIn, err := net.ListenUDP("udp", serverAddrIn)
	checkError(err)

	serverAddrOut, err := net.ResolveUDPAddr("udp", pOut)
	checkError(err)

	connOut, err := net.ListenUDP("udp", serverAddrOut)
	checkError(err)

	quitPin := make(chan struct{})
	quitPout := make(chan struct{})

	clientsB := make(map[*net.UDPAddr]int)

	for i := 0; i < runtime.NumCPU(); i++ {
		go listenToA(connIn, quitPin)
		go listenToB(connOut, clientsB, quitPout)
		//	go distrubite(clientsB, buffer)
	}
	<-quitPin
	<-quitPout
}
