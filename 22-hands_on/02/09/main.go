// Add code to respond to the following METHODS & ROUTES: GET / GET /apply POST /apply

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
			fmt.Println("")
		}
		if ln == "" {
			break
		}
		i++
	}

	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/apply" {
		apply(conn)
	}
	if m == "POST" && u == "/apply" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {
	body := "<h1>HOLY COW THIS IS INDEX LEVEL</h1>"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

	conn.Close()
}

func apply(conn net.Conn) {
	body := `
		<h1>HOLY COW THIS IS LOW LEVEL AND BUTTON</h1>
		<form method="POST" action="/apply">
		<input type="submit" value="applyPost">
		</form>
		`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

	conn.Close()
}

func applyProcess(conn net.Conn) {
	body := "<h1>Processed</h1>"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

	conn.Close()
}
