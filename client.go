package main

import (
  "log";
  "net";
  "fmt";
  "bufio";
  "os"
)

func handleConnection(c net.Conn) {
  buf := make([]byte, 4096)
  
  for {
    n, err := c.Read(buf)
    if err != nil || n == 0 {
      c.Close()
      break
    }
    n, err := c.Write(buf[0:n])
    if err != nil {
      c.Close()
      break
    }
  }
  log.Printf("Connection to %v closed.", C.RemoteAddr())
}

func main() {
  msgchan := make(chan string)
  go printMessages(msgchan)

  conn, err := net.Dial("tcp", "10.0.0.42:23")
  if err != nil {
    //handle error
    log.Fatal(err)
    os.Exit(1)
  }
  go handleConnection(conn, msgchan)
}
