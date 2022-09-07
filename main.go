package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/harrisoncramer/go-connect-tcp/client"
)

const p = "13102"

type Monitor struct {
	*log.Logger
}

func (m *Monitor) Write(p []byte) (int, error) {
	return len(p), m.Output(2, string(p))
}

func main() {

	host := flag.String("h", "", "The host (only applicable for clients)")
	port := flag.String("p", "", "The port on which to run the program")

	monitor := &Monitor{Logger: log.New(os.Stdout, "monitor: ", 0)}

	flag.Parse()

	if *port == "" {
		monitor.Println("Using default port")
		*port = p
	}

	PORT := ":" + *port

	if *host != "" {
		client.Run(*host, PORT)
	} else {

		l, err := net.Listen("tcp", PORT)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer l.Close()

		monitor.Printf("TCP server listening on port %s", PORT)

		for {
			conn, err := l.Accept()
			if err != nil {
				monitor.Fatal(err)
			}

			go handleRequest(conn, *monitor)
		}
	}

}

func handleRequest(conn net.Conn, monitor Monitor) {

	for {
		/* Creates a reader for the data and also writes to the monitor */
		r := io.TeeReader(conn, &monitor)
		buf := make([]byte, 8192)
		_, err := r.Read(buf)

		if err != nil {
			monitor.Fatal(err)
		}
		if strings.TrimSpace(string(buf)) == "STOP" {
			monitor.Println("Exiting TCP server!")
			return
		}

		/* Here we could plausibly perform some other action or work with the buffer */
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		conn.Write([]byte("Server responded at " + myTime))
	}
}
