package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func makeRot13Code(word []byte) []byte {
	var r13 = make([]byte, len(word))
	for i, v := range word {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}

func connectionHandler(conn net.Conn) {
	Scanner := bufio.NewScanner(conn)
	defer conn.Close()
	for Scanner.Scan() {
		lowerWord := strings.ToLower(Scanner.Text())
		lw := []byte(lowerWord)
		cypherText := makeRot13Code(lw)
		fmt.Printf("%s\n", cypherText)
	}
}

func main() {
	tcpListner, err := net.Listen("tcp4", ":8081")
	if err != nil {
		log.Fatalln(err)
	}
	defer tcpListner.Close()
	for {
		r13Conn, err := tcpListner.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("\nConnected!\n")
		connectionHandler(r13Conn)
	}
}
