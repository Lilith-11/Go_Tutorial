package main

import (
    "fmt"
    "net"
    "bufio"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()
    
    fmt.Println("Connected to server. Type a message:")
    reader := bufio.NewReader(os.Stdin)
    message, _ := reader.ReadString('\n')
    conn.Write([]byte(message))
}