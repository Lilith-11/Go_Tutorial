package main
import "fmt"
import "os"
func main(){
	fmt.Println(os.Args)
	fmt.Println(os.Args[0])
	//fmt.Println(os.Args[1])
	fmt.Println(len(os.Args))
}