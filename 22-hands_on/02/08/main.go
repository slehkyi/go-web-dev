// Change your RESPONSE HEADER "content-type" from "text/plain" to "text/html"
// Change the RESPONSE from "CHECK OUT THE RESPONSE BODY PAYLOAD" (and everything else it contained: request method, request URI) to an HTML PAGE that prints "HOLY COW THIS IS LOW LEVEL" in tags.

package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"io"
	"strings"
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
	var i int
	var m, u string
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			// request line
			m = strings.Fields(ln)[0] // method
			u = strings.Fields(ln)[1] // uri
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)
		}
		if ln == "" {
			break
		}
		i++
	}
	body := "<h1>HOLY COW THIS IS LOW LEVEL</h1>"
	fmt.Println("Code got here.")
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

	conn.Close()
}
