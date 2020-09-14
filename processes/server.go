package main

import "net"
import "fmt"
import "bufio"
import "strings"
import "os"

func main() {

  arguments := os.Args
  port := arguments[1]
  
  // listen on all interfaces
  fmt.Println(port)

  ln, _ := net.Listen("tcp", ":" + port)
  fmt.Println("Listening on port" + port)

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