package main

import (
    "bufio"
    "fmt"
    "log"
	"os"
	"strings"
	"net"
	"io/ioutil"
)

func main() {
	for {
		go dial()
		// go acceptRequests()
	}
	// Read()
}
func Read(){
	content, err := ioutil.ReadFile("config.txt")
	if err != nil {
    //Do something
	}	
	lines := strings.Split(string(content), "\n")
	fmt.Printf("File contents: %s", content[0])
	fmt.Print(lines)
	
}

func dial() {
	ip := "127:0:0:1"
	port := "6001"
	TCPdial(ip,port)

}

func acceptRequests() {
	line := 0
	f, err := os.Open("config.txt")
	arguments := os.Args

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

	scanner := bufio.NewScanner(f)

    for scanner.Scan() {
		if(line != 0) {
			process := strings.Split(scanner.Text(), " ")[0]
			port := strings.Split(scanner.Text(), " ")[2]

			//check and parse all ip/host that aren't the process itself
			if(process != arguments[1]) {
				listener(port)
			}
		}
		line = line + 1
	}
	
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func TCPdial(ip string,port string) {
	//dial a tcp connection with every other ip:port
	//check if connection is alive, else keep trying

		// conn, err := net.Dial("tcp", ip + ":" + port)
		conn, err := net.Dial("tcp", "127.0.0.1:6001")
		if err != nil {
			fmt.Println(err)
			// panic(err)
			// handle error
		} else {
				fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
				status, err := bufio.NewReader(conn).ReadString('\n')

				if err != nil {
				// handle error
				} 
				for { 
    				// read in input from stdin
    				reader := bufio.NewReader(os.Stdin)
    				fmt.Print("Message and destination? ")
    				text, _ := reader.ReadString('\n')
					// send to socket
					fmt.Print(text)
					// unicast_send()
  		}
				fmt.Printf(status)
		}				 
	
}


func listener(port string) {
	ln, err := net.Listen("tcp", ":" + port)
	if err != nil {
		// handle error
		fmt.Println(err)
		panic(err)
		} else {
			for {
				fmt.Printf("listening on port: " + port)
				conn, err := ln.Accept()
				if err == nil {
					break
					// go handleConnection(conn)
					status, err := bufio.NewReader(conn).ReadString('\n')
					fmt.Printf(status)
					fmt.Println(err)
				break
				}
		}
	}
}
func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

	}
	c.Close()
}


//Delivers the message received from the source process
/*func unicast_receive(source string, message string) {
	fmt.Printf("Received %q from process %s , system time is %s\n", message, source, time.Now())
}*/