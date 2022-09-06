package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/harrisoncramer/go-connect-tcp/client"
)

const p = "13102"

func main() {

	host := flag.String("h", "", "The host (only applicable for clients)")
	port := flag.String("p", "", "The port on which to run the program")

	flag.Parse()

	if *port == "" {
		log.Println("Using default port")
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

		log.Printf("TCP server listening on port %s", PORT)

		for {
			conn, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}

			go handleRequest(conn)
		}
	}

}

func handleRequest(conn net.Conn) {

	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		/* Here we could plausibly perform some other action or work */
		fmt.Print(string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		conn.Write([]byte("Server responded at " + myTime))
	}
}
