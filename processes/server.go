package main

import (
"bufio"
"fmt"
"net"
"os"
"strconv"
"strings"
"time"
)

//Delivers the message received from the source process
func unicast_receive(source string, message string) {
	fmt.Printf("Received %q from process %s, system time is %s\n", message, source, time.Now())
}

var count = 0

func handleConnection(c net.Conn) {
  for {
    netData, err := bufio.NewReader(c).ReadString('\n')
    if err != nil {
      fmt.Println(err)
      return
    }
    str := strings.TrimSpace(string(netData))
    
    if str == "STOP" {
        break
    }

    //parse str that source sent to split up source and message
    source := strings.Split(str, " ")[0]
    message := strings.Split(str, " ")[1]
    // message = strings.Replace(destination, " ", "", -1)

    unicast_receive(source, message)
    counter := strconv.Itoa(count) + "\n"
    c.Write([]byte(string(counter)))
    }
    c.Close()
}

func main() {
  arguments := os.Args
  if len(arguments) == 1 {
    fmt.Println("Please provide a port number!")
    return
  }

  PORT := ":" + arguments[1]
  l, err := net.Listen("tcp4", PORT)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer l.Close()

  for {
    c, err := l.Accept()
    if err != nil {
      fmt.Println(err)
      return
    }
    go handleConnection(c)
      count++
    }
}