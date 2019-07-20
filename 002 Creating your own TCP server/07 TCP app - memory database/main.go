package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func handle(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		input := strings.Fields(ln)
		if len(input) < 1 {
			continue
		}
		switch input[0] {
		case "GET":
			key := input[1]
			if value, ok := data[key]; ok {
				fmt.Fprintf(conn, "%s\n", value)
			} else {
				fmt.Fprintln(conn, "error: UNDEFINED VALUE")
			}
		case "SET":
			if len(input) != 3 {
				fmt.Fprintln(conn, "error: Expected key and value!")
				continue
			}
			key := input[1]
			value := input[2]
			data[key] = value
		case "DEL":
			key := input[1]
			delete(data, key)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND "+input[0])
			continue
		}
	}
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

// run server: go run main.go
// connect to server: telnet localhost 8080
