package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/harrisoncramer/go-serve-tcp/client"
)

func main() {

	host := flag.String("host", "", "The host (only applicable for clients)")
	port := flag.String("port", "", "The port on which to run the program")

	flag.Parse()

	if *port == "" {
		log.Fatal("Must provide port")
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
