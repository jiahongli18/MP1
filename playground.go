package main 
  
import ( 
    "fmt"
    "time"
    "math/rand"
) 
  
// Here, the value of Sleep function is zero 
// So, this function return immediately. 
func show(min int, max int) { 
        num := rand.Intn(max-min) +min
	    time.Sleep(time.Duration(num) * time.Millisecond)
        fmt.Println("1") 
} 
  
// Main Function 
func main() { 
  
    // for {
        go show(2000,3000) 
    // }

    fmt.Println("2") 
    
} 