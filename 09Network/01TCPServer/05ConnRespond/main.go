package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	OK = 200 + iota
	CREATED
	ACCEPTED
	NonAUTHiNFO
	NoCONTENT
)

const (
	BadREQUEST = 400 + iota
	UnAUTHORIZED
	PaymentREQUIRED
	FORBIDDEN
	NOTFOUND
)

var responseStatus = make(map[int]string, 20)

func responder(conn net.Conn, method, resource string, Status int, message string) {
	var httpResponseHeadder string = "HTTP/1.1 %d %s\r\n" +
		"Content-Length: %d\r\n" +
		"charset: UTF-8\r\n" +
		"Content-Type: text/html\r\n\r\n"

	var BasicHTML string = `
				<!DOCTYPE HTML>
				<HTML lang="en">
				<head>
					<meta charset="UTF-8">
					<title>MetaPrac</title>
				</head>
				<body>
					<h1>%s</h1>
					<p>
					Method: %s                              </p>
					<p>Resource: %s                         </p>
				</body>
				</HTML>`
	fmt.Fprintf(conn, httpResponseHeadder, Status, responseStatus[Status], len(BasicHTML))
	fmt.Fprintf(conn, BasicHTML, message, method, resource)
}

func conRespond(conn net.Conn, requestedDetails []string) {
	var method string = requestedDetails[0]
	var resource string = strings.ToLower(requestedDetails[1])
	switch resource {
	case "/":
		message := "Welcome Home"
		responder(conn, method, resource, OK, message)
	case "/hello":
		message := "Hello, You are Fasaak!"
		responder(conn, method, resource, OK, message)
	default:
		message := "The Page you are looking for is NOT FOUND"
		responder(conn, method, resource, NOTFOUND, message)
	}
}

func conScanRequest(conn net.Conn) []string {
	var checkerI uint8 = 0
	var requestDetails []string
	Scanner := bufio.NewScanner(conn)
	for Scanner.Scan() {
		var data string = Scanner.Text()
		log.Default().Printf("%s", data)
		if checkerI == 0 {
			requestDetails = strings.Fields(data)
		}
		if data == "" {
			break
		}
		checkerI++
	}
	return requestDetails
}

func connectionHandler(conn net.Conn) {
	defer conn.Close()
	var requestedDetails []string = conScanRequest(conn)
	conRespond(conn, requestedDetails)
}

func acceptConnections(tcpListner net.Listener) {
	tcpConnection, err := tcpListner.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	go connectionHandler(tcpConnection)
}

func startTcpServer(addr string) net.Listener {
	tcpListner, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatalln(err)
	}
	return tcpListner
}

func main() {
	responseStatus = map[int]string{
		OK:              "OK",
		CREATED:         "CREATED",
		ACCEPTED:        "ACCEPTED",
		NonAUTHiNFO:     "NonAUTHiNFO",
		NoCONTENT:       "NoCONTENT",
		BadREQUEST:      "BadREQUEST",
		UnAUTHORIZED:    "UnAUTHORIZED",
		PaymentREQUIRED: "PaymentREQUIRED",
		FORBIDDEN:       "FORBIDDEN",
		NOTFOUND:        "NOTFOUND"}
	var tcpListner8081 = startTcpServer(":8081")
	defer tcpListner8081.Close()
	for {
		acceptConnections(tcpListner8081)
	}

}
