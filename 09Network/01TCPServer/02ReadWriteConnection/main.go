package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func handleConnection(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(20 * time.Second)) //Setting the timeout for connections, no matter what these connections will die.
	strDoYou := "You:\t"
	Scanner := bufio.NewScanner(conn)
	for Scanner.Scan() {
		content := Scanner.Text()
		fmt.Fprintf(conn, strDoYou) // Writes the string to the connection.
		fmt.Printf("%v\n", content) // Scans and Prints the content sent from the client on the server end
	}
	defer conn.Close()
	fmt.Fprintf(conn, "\n\n\n\t..\t...You will Never see This")
	fmt.Printf("\n\n\n\t..\t...Closing Connection.") // This will be printed at the Server end while client connection is terminated
}

func main() {

	tcpListener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer tcpListener.Close()
	io.WriteString(os.Stdout, "\nStarted Server\n") //Will be displayed at the Listener's front only once when started
	for {
		acceptedConnection, err := tcpListener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Commonly displayed to all the connections.
		// acceptedConnection.SetDeadline(time.Now().Add(20 * time.Second)) //Setting the timeout for all connections
		// but this is dangerous as it will terminate outside of handler where handler is calling Close()
		io.WriteString(acceptedConnection, "\nACCEPTING CONNECTIONS")
		io.WriteString(acceptedConnection, "\t....RudeOff1\tTakeOff!\n")
		go handleConnection(acceptedConnection)
	}
}

// if go routines are not used then connections will be queued until the previous connection is closed, then a new connection will be added.
// This connection runns forever until a manual timeout is given
