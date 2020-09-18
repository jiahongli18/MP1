package main

import (
    "fmt"
    "net"
    "os"
    "sync"
    "./processes"
)


func main() {
    arguments := os.Args

    go processes.StartServer(arguments[1])
    readConfig()
    initialize(arguments[1])
    processes.Unicast()
}

func readConfig() {
   
}

func initialize(processNum string) {
    ip := "127.0.0.1"
    ports := [3]string{"6001","6002","6003"}

    for port := range ports {
        for {
            status := dial(processNum, ports[port], ip)

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

func dial(processNum string, port string, ip string)(status string){
    if(port != processNum) {
        address := ip + ":" + port

        _, err := net.Dial("tcp", address)
        if err != nil {
            return "error"
        }
        // fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    }
    return "success"
}