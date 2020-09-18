package main

import (
     "fmt"
    "net"
    "os"
    "sync"
    "./processes"
    "./Utils"
    // "encoding/gob"
)


func main() {
    arguments := os.Args

    go processes.StartServer(arguments[1])
    ip, port, min_delay, max_delay, ports := Utils.ReadConfig(arguments[1])
    fmt.Println(ip, port, min_delay, max_delay, ports)

    ip = arguments[1]
    initialize(ip, arguments[1], ports)
    processes.Unicast(ip, port, min_delay, max_delay, ports)
}

// func readConfig(port string) {

//     port = ":" + port
//     //listen for TCP requests
// 	l, err := net.Listen("tcp", port)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	} else {
// 		fmt.Println("Listening on port" + port)
// 	}
// 	defer l.Close()

//     go Utils.ParseAll("1")

// 	for {
// 		//The server accepts and begins to interact with TCP client
// 		c, err := l.Accept()
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
//         decoder := gob.NewDecoder(c) //initialize gob decoder

//         //Decode message struct and print it
//         Data := new(processes.Data)
//         data := decoder.Decode(Data)
//         fmt.Println(data)
//     }
// }

func initialize(processIP string, processPort string, ports []string) {
    // ip := "127.0.0.1"
    portss := [3]string{"6001","6002","6003"}

    for port := range portss {
        for {
            status := dial(processPort, portss[port], processIP)

            if(status == "success") {
                break
            }

            fmt.Println(status)

            groupTest := new(sync.WaitGroup)
            go processes.Delay(2000,2500,groupTest)
 
            groupTest.Add(1)
            groupTest.Wait()
        }
    }
}

func dial(processPort string, port string, ip string)(status string){
    fmt.Println(processPort)
    if(port != processPort) {
        address := ip + ":" + port
        fmt.Print(address)
        _, err := net.Dial("tcp", address)
        if err != nil {
            return "error"
        }
        // fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    }
    return "success"
}