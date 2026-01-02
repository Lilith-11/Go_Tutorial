package main
import ("fmt"
"os")
func main(){
	err:=os.Remove("hello.txt")
	if err!=nil{
		fmt.Println("error",err)
		return
	}
	fmt.Println("file removed/ deleted successfully")
}