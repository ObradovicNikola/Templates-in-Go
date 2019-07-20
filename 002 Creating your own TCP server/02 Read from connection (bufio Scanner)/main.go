package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// how does the above reader know when it's done?
	// listening all the time, never gets to defer conn.Close()
	// code never gets here
	fmt.Println("Code got here.")
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

// in browser: localhost:8080
