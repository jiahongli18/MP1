package Utils

import (
    // "fmt"
    "os"
	"bufio"
	"strconv"
    "strings"
	"log"
	// "encoding/gob"
	// "net"
	// "../processes"
)

func ReadConfig(destination string)(string, string, int, int, []string){
    ip, port := fetchHostPort(destination)
    min_delay, max_delay := fetchDelay()

	ports := []string{"6001","6002","6003"}

	return ip, port, min_delay, max_delay, ports
}

// func parseConfig(destination string) (string, string, int, int, []string){
	

    // // create TCP channel
    // CONNECT := destination
	// c, err := net.Dial("tcp", CONNECT)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	
    
	// //create a gob encoder and code the data struct
	// encoder := gob.NewEncoder(c)
	// data := processes.Data{ip, port, min_delay, max_delay, ports}
    // // fmt.Println(data)
	// _ = encoder.Encode(data)
	
// }

//parses config.txt and returns ip and host
func fetchHostPort(destination string) (string, string){
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
func fetchDelay()(int, int){
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