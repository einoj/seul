package main

import "net"
import "fmt"
import "bufio"

func main() {

ln, err := net.Dial("tcp", "10.0.0.2:23")
if err != nil {
  //handle error
}
  for {
    message, _ := bufio.NewReader(ln).ReadString('\n')
    fmt.Print("Message Received:", string(message))
  }
}
