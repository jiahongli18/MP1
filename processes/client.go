package main

import (
  "net"
  "fmt"
  "bufio"
  "os"
  "strings"
  "log"
  "time"
  "math/rand"
  "strconv"
)

func delay(min int, max int){
	num := rand.Intn(max-min) +min
	time.Sleep(time.Duration(num) * time.Millisecond)
  fmt.Println("done") 
}

func fetchdelay()(int, int){
  f, err := os.Open("../config.txt")
  if err != nil {
    log.Fatal(err)
  }

  scanner := bufio.NewScanner(f)
  scanner.Scan()
  delays := strings.Fields(scanner.Text())
  min_delay, _ := strconv.Atoi(delays[0])
  max_delay, _ := strconv.Atoi(delays[1])
  f.Close()
  return min_delay, max_delay
}

func main() {
  
  for { 
    // read in input from stdin
    fmt.Print("Please enter a command: ")

    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')

    //Check errors for the user command
    if len(text) < 3 {
      fmt.Println("Please include destination and message properly.")
      return
    }

    //get the process # from the user's input
    destination := strings.Split(text, " ")[1]
    destination = strings.Replace(destination, " ", "", -1)

    //get the message from the user's input
    message := strings.Split(text, " ")[2]
    message = strings.Replace(message, " ", "", -1)
    
    //find the associating host/port according to the user's desired destination #
    ip, port := fetchHostPort(destination)
    min_delay, max_delay := fetchdelay()
    
    // go delay(3000,5000)
    unicast_send(destination, ip + ":" + port, message)

    // listen for reply
    // message, _ := bufio.NewReader(conn).ReadString('\n')
    // fmt.Print("Message from server: "+message)
  }
}

//parses config.txt and returns ip and host
func fetchHostPort(destination string) (string, string){
	line := 0
	f, err := os.Open("../config.txt")
	// arguments := os.Args

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

	scanner := bufio.NewScanner(f)
  
  for scanner.Scan() {
		if(line != 0) {
			process := strings.Split(scanner.Text(), " ")[0]
			ip := strings.Split(scanner.Text(), " ")[1]
			port := strings.Split(scanner.Text(), " ")[2]

			if(process == destination) {
        // fmt.Println(ip,port)
				return ip,port

			}
		}
		
		line = line + 1
	}
	
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

  return "nil","nil"
}

//Sends message to the destination process
func unicast_send(process string, destination string, message string) {
  //dial to the TCP server using net library
  conn, err := net.Dial("tcp", destination)
  if err != nil {
    fmt.Println(err)
    return
  }
  // send to socket
  fmt.Fprintf(conn, process + " " + message)
	fmt.Printf("Sent '%s' to process %s , system time is %s\n", message, process, time.Now().Format("Jan _2 15:04:05.000"))
}

//Simulate network delay by adding an extra layer before sending the message via the TCP channel
func Delay(min int, max int) {
	num := rand.Intn(max-min) + min
	time.Sleep(time.Duration(num) * time.Millisecond)
	return
}

