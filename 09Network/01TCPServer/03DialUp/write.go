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
	io.WriteString(dialedConnection, "\nGitmo\n")
	fmt.Fprintf(dialedConnection, "You got no Idea, Who am I.\n")
	fmt.Fprintf(dialedConnection, "lbh-tbg-ab-vqrn9-jub-nz-v.\n")
	fmt.Printf("Closing-Client\n")
}

// Just connects once and reads all that is there in the connection.
