package main
import "fmt"
func mypanic(){
	panic("this is a panic messege")
}
func main(){
	defer func(){
		if r:=recover();r!=nil{
			fmt.Println("error occured",r)
		}
	}()
	mypanic()
	fmt.Println("this will not run")
}