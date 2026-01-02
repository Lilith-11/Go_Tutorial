package main
import ("fmt"
"os")
func main(){
	file,err:=os.Create("file.txt")
	if err!=nil{
		fmt.Println("error",err)
	}
	defer file.Close()
	fmt.Println("file create sucuessfully")
	_,err=file.WriteString("hello world!!! \n")
	if err!=nil{
		fmt.Println("error is ",err)
	}
	fmt.Println("written to the the file successfully")
}