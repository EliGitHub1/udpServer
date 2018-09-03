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
		fmt.Println("Please provide a Pin")
		return
	}

	pIn := ":" + os.Args[1]

	connToServer, err := net.Dial("udp", pIn)
	checkError(err)

	b := make([]byte, 1024)

	for {
		time.Sleep(100 * time.Millisecond)
		_, err = connToServer.Write(b)
		checkError(err)
	}
	connToServer.Close()
}
