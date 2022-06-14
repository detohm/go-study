package main

import (
	"fmt"
	"net"
)

func main() {
	addr := ":8000"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	host, port, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Listening on host: %s, port: %s\n", host, port)

	for {
		// waiting for an incoming connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		// handle connections in a new goroutine
		go func(conn net.Conn) {
			buf := make([]byte, 1024)
			len, err := conn.Read(buf)
			if err != nil {
				fmt.Printf("Error reading %#v\n", err)
				return
			}
			fmt.Printf("Message received: %s\n", string(buf[:len]))
			conn.Write([]byte("Message received.\n"))
			conn.Close()
		}(conn)
	}
}
