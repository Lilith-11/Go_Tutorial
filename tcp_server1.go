package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error:", err.Error())
        return
    }
    defer ln.Close()

    fmt.Println("Server is listening on port 8080...")
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error:", err.Error())
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    fmt.Println("Client connected:", conn.RemoteAddr().String())
    reader := bufio.NewReader(conn)
    message, _ := reader.ReadString('\n')
    fmt.Println("Received message:", message)
    conn.Close()
}