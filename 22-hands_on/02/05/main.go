// Add code to WRITE to the connection.

package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"io"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			fmt.Println("end of the headers")
			break
		}
		fmt.Println(ln)
	}
	io.WriteString(conn, "I see you connected.")

	conn.Close()
}
