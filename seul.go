package main

import "net"
import "fmt"
import "bufio"

func main() {

ln, err := net.Listen("tcp", ":9001")
if err != nil {
  //handle error
}
  conn, err := ln.Accept()
  if err != nil {
    // handle  error
  }
  for {
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message Received:", string(message))
    conn.Write([]byte(message+ "\n"))
  }
}
