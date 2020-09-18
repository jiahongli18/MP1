package processes

import (
"bufio"
"fmt"
"net"
"strings"
"time"
)

//Delivers the message received from the source process and prints out message, source, and time
func unicast_receive(source string, message string) {
	fmt.Printf("\nReceived %q from process %s, system time is %s\nPlease enter a command: ", message, source, time.Now().Format("Jan _2 15:04:05.000"))
}

// var count = 0

func handleConnection(c net.Conn) {
  for {
    //read message and source from request
    netData, err := bufio.NewReader(c).ReadString('\n')
    if err != nil {
      fmt.Println(err)
      return
    }
    str := strings.TrimSpace(string(netData))

    //parse str that source sent to split up source and message
    source := strings.Split(str, " ")[0]
    message := strings.Split(str, " ")[1]
    // message = strings.Replace(destination, " ", "", -1)

    unicast_receive(source, message)
    }
    c.Close()
}

func StartServer(port string) {

  //get port number from user input and listen in on that port for requests
  PORT := ":" + port
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

    //goroutine for handling requests made to server
    go handleConnection(c)
    }
}