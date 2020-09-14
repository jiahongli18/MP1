package main

import (
    "bufio"
    "fmt"
    "log"
	"os"
	"strings"
	"net"
)

func main() {

	line := 0
	f, err := os.Open("config.txt")
	arguments := os.Args

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

	scanner := bufio.NewScanner(f)

	//First for loop: for dialing and establishing connection with every other ip:host
    for scanner.Scan() {
		if(line == 0) {
			min_delay := strings.Split(scanner.Text(), " ")[0]
			max_delay := strings.Split(scanner.Text(), " ")[1]
			fmt.Println(min_delay, max_delay)
		}
		if(line != 0) {
			process := strings.Split(scanner.Text(), " ")[0]
			ip := strings.Split(scanner.Text(), " ")[1]
			port := strings.Split(scanner.Text(), " ")[2]

			//check and parse all ip/host that aren't the process itself
			if(process != arguments[1]) {
				fmt.Println(process, ip, port)

				//dial a tcp connection with every other ip:port
				//check if connection is alive, else keep trying

				x := 1
				for x < 1000 {
					conn, err := net.Dial("tcp", ip + ":" + port)
					if err != nil {
						// fmt.Println(err)
						// handle error
					} else {
						fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
						status, err := bufio.NewReader(conn).ReadString('\n')
		
						if err != nil {
							// handle error
						} 

						fmt.Printf(status)
						go acceptRequests(port)
						break
					}				 
				 }
				}
		}
		
		line = line + 1
	}
	
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func acceptRequests(port string) {
	ln, err := net.Listen("tcp", ":" + port)
	if err != nil {
	// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
		// handle error
		}
		// go handleConnection(conn)
		status, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf(status)
	}
}