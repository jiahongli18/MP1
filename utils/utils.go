package Utils

import (
  "os"
	"bufio"
	"strconv"
  "strings"
	"log"
  "sync"
  "time"
  "math/rand"
)

func ReadConfig(destination string)(string, string, int, int, []string){
    ip, port := FetchHostPort(destination)
    min_delay, max_delay := FetchDelay()

	ports := []string{"6001","6002","6003"}

	return ip, port, min_delay, max_delay, ports
}

//parses config.txt and returns ip and host
func FetchHostPort(destination string) (string, string){
	line := 0
	f, err := os.Open("./config.txt")

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

//parses the min and max delays from the config file
func FetchDelay()(int, int){
  f, err := os.Open("./config.txt")
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

//Simulate network delay by adding an extra layer before sending the message via the TCP channel
func Delay(min int, max int, groupTest *sync.WaitGroup){
	num := rand.Intn(max-min) +min
	time.Sleep(time.Duration(num) * time.Millisecond)

  //decrement value of waitgroup and relay the flow of execution back to main  
  groupTest.Done()    
}
