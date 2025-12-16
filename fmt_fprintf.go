package main

import (
    "fmt"
    "os"
)

func main() {
    _,err:=fmt.Fprintf(os.Stdout,"hello,%s!\n","world")
	if err!=nil{
		fmt.Println("hello this is error")
	}
}