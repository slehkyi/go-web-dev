// Print to standard out (the terminal) the REQUEST method and the REQUEST URI from the REQUEST LINE.
// Add this data to your REPONSE so that this data is displayed in the browser.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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
	body := fmt.Sprintf("<h1>Hello World!</h1><br>METHOD: %v<br>URI: %v", m, u)
	fmt.Println("Code got here.")
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

	conn.Close()
}
