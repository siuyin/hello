package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/siuyin/dflt"
)

func main() {
	fmt.Printf("text protocol server demo.\nThis is just a regular server serving on a TCP port\n")
	port := dflt.EnvString("PORT", ":3999")
	fmt.Printf("server listening on port %q\n", port)

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			fmt.Fprintf(c, "100 textproto server accepted connection from %v %s\n", c.RemoteAddr(), time.Now().Format("15:04:05.000"))
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
			fmt.Printf("closed connection from %v\n", c.RemoteAddr())
		}(conn)
	}
}
