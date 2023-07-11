package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	dialedConnection, err := net.Dial("tcp4", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer dialedConnection.Close()
	strBits, err := io.ReadAll(dialedConnection)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", string(strBits))
}

// Just connects once and reads all that is there in the connection.
