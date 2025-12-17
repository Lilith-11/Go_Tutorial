package main
import ("fmt"
        "time")
func sayhello(){
	fmt.Println("this is go language")

}
func main(){
	go sayhello()
	time.Sleep((1)*time.Second)
	message := make(chan string)
	go func(){
		message<-"hello world"
	}()
	msg:=<-message
	fmt.Println(msg)
	//buffered channels 
	words:=make(chan string,2)
	words<-"hello"
	words<-"world"
	fmt.Println(<-words)
	fmt.Println(<-words)
}