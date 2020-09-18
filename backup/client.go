package processes

import (
  "net"
  "fmt"
  "bufio"
  "os"
  "strings"
  // "log"
  "time"
  "math/rand"
  // "strconv"
  "sync"
)

func Unicast(ip string, port string, min_delay int, max_delay int, ports []string) {
  
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
    // ip, port := fetchHostPort(destination)
    //ip := "127.0.0.1"
    //port:= "6002"
    //min_delay, max_delay := fetchdelay()
    //min_delay, max_delay := 2000,3000

    unicast_send(destination, ip + ":" + port, message, min_delay, max_delay)
  }
}

//Simulate network delay by adding an extra layer before sending the message via the TCP channel
func Delay(min int, max int, groupTest *sync.WaitGroup){
	num := rand.Intn(max-min) +min
	time.Sleep(time.Duration(num) * time.Millisecond)

  //decrement value of waitgroup and relay the flow of execution back to main  
  groupTest.Done()    
}

//Sends message to the destination process
func unicast_send(process string, destination string, message string, min_delay int, max_delay int) {
  //dial to the TCP server using net library
  conn, err := net.Dial("tcp", destination)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("Sent '%s' to process %s, system time is %s\n", message, process, time.Now().Format("Jan _2 15:04:05.000"))

   //set delay
    groupTest := new(sync.WaitGroup)
    go Delay(min_delay,max_delay,groupTest)
 
    //Wait group is used to block the execution of code in the main thread until all goroutines are complete and waitgroup counter is decremented to 0
    groupTest.Add(1)
    groupTest.Wait()
  // send to socket
  fmt.Fprintf(conn, process + " " + message)
	
}


