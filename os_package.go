package main
import ("fmt"
"os")
func main(){
	dir,err:=os.Getwd()
	if err!=nil{
		fmt.Print("error")
        return
	}
	fmt.Println(dir)
}