package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	tcpListner, err := net.Listen("tcp4", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer tcpListner.Close()

	for {
		acceptedConnections, err := tcpListner.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer acceptedConnections.Close()
		io.WriteString(acceptedConnections, "\nHello!\n\tMessage from TCP\n")
		fmt.Fprintf(acceptedConnections, "%v", "\nAntar Graha Vasi\n")
	}

}

// This code is capable of accepting only one connection.
//
