package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	"io/ioutil"
)

func delay(min int, max int) {
	num := rand.Intn(max-min) + min
	time.Sleep(time.Duration(num) * time.Millisecond)
	return
}

//unicast_receive starts like a TCP server
func unicast_receive(source string, message string) {
	//PORT := ":" + source
	l, err := net.Listen("tcp", source)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		netData, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Println(netData)
		fmt.Println(time.Now())
		return
	}
}

func Application(Sender string, Dest string, message string) {
	var SendIP string
	var DestPort string
	var min_delay int
	var max_delay int
	//Read the config file to get IP addresses and ports
	f, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	delays := strings.Fields(scanner.Text())
	min_delay, _ = strconv.Atoi(delays[0])
	max_delay, _ = strconv.Atoi(delays[1])
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if words[0] == Sender {
			SendIP = words[1]
		}
		if words[0] == Dest {
			DestPort = words[2]
		}
		//fmt.Println(words)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	//fmt.Println("SendIP", SendIP)
	//fmt.Println("DestPort", DestPort)
	//delay(min, max)

	unicast_receive(":"+DestPort, message)
	IpPort := SendIP + ":" + DestPort
	unicast_send(IpPort, message)

	//unicast_receive(":8080", message)
	//unicast_send("golang.org:80", message)

	go delay(min_delay, max_delay)
}

func main() {
	/*var Sender string
	var Dest string

	//channel := make(chan string)
	//Scan user input for process number
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please specify the process you want to run.")
		return
	}
	Sender = arguments[1]

	//scan user input for destination and Message
	var cmd string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("User Command : ")
	cmd, _ = reader.ReadString('\n')
	cmd1 := strings.Split(cmd, " ")
	if len(cmd1) < 3 {
		fmt.Println("Please include destination and message properly.")
		return
	}
	Dest = strings.Split(cmd, " ")[1]
	message := strings.Split(cmd, " ")[2]
	fmt.Print(message)

	Application(Sender, Dest, message)
	*/
	return
}

func Read(){
	content, err := ioutil.ReadFile("config.txt")
	if err != nil {
    //Do something
	}	
	lines := strings.Split(string(content), "\n")
	fmt.Printf("File contents: %s", content[0])
	fmt.Print(lines)
	//var min_delay int
	//var max_delay int
}


//for now for destination arg, put "127:0:0:1:5000"

//start server
/*func startServer(destination string) {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":" + destination.split(":")[1])

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
	  // will listen for message to process ending in newline (\n)
	  message, _ := bufio.NewReader(conn).ReadString('\n')
	  // output message received
	  fmt.Print("Message Received:", string(message))
	  // sample process for string received
	  newmessage := strings.ToUpper(message)
	  // send new string back to client
	  conn.Write([]byte(newmessage + "\n"))
	}
}
*/

//unicast_send acts like a TCP client
func unicast_send(destination string, message string) {

	CONNECT := destination
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Fprintf(c, message)
		return
	}
}
