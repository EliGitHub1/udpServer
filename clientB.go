package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
	}
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please provide a Pout ")
		return
	}

	pOut := ":" + os.Args[1]

	connToServer, err := net.Dial("udp", pOut)
	checkError(err)

	b := make([]byte, 1024)

	for {
		time.Sleep(1 * time.Second)
		_, err = connToServer.Write(b)
		checkError(err)
	}
	connToServer.Close()
}
