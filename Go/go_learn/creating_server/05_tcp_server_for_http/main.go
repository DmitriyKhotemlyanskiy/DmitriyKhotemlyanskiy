package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	//read request
	request(conn)

	//write response
	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			//request line
			m := strings.Fields(ln)[0] //method
			fmt.Println("****Method", m)
		}
		if ln == "" {
			//headers are done
			break
		}
		i++
	}
}

func response(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta
	 charset="UTF-8"><title></title></head><body><strong>Hello World</strong>
	 <scrypt><button type="submit" name="formBtn">Submit</button></scrypt></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
