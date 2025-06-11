package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

type target struct {
	port int
	host string
}

func main() {
	t := parseFlags()
	if t == nil {
		return
	}

	for i := 0; i < 10; i++ {
		ping(t)
	}
}

func ping(t *target) {
	p := make([]byte, 2048)
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", t.host, t.port))
	if err != nil {
		fmt.Printf("net.Dial: %v", err)
		return
	}
	fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
	_, err = bufio.NewReader(conn).Read(p)
	if err != nil {
		fmt.Printf("bufio.NewReader: %v\n", err)
	}

	fmt.Printf("%s\n", p)
	conn.Close()
}

func parseFlags() *target {
	if len(os.Args) < 1 {
		fmt.Println("Not enough arguments...")
		return nil
	}

	host := os.Args[len(os.Args)-1]
	port := flag.Int("port", 53, "invalid type, should be int")

	flag.Parse()

	fmt.Printf("Ping: %s on port %d \n", host, *port)

	return &target{
		host: host,
		port: *port,
	}
}
