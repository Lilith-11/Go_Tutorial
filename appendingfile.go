package main
import ("os"
"fmt"
"io/ioutil")
func main(){
	file,err:=os.OpenFile("file.txt",os.O_APPEND|os.O_WRONLY,0644)
	if err!=nil{
		fmt.Println("error",err)
		return
	}
	defer file.Close()
	_,err=file.WriteString("this is appending text\n")
		if err!=nil{
		fmt.Println("error",err)
		return
	}
	file,err=os.Open("file.txt")
	data,err:=ioutil.ReadAll(file)
	fmt.Println(string(data))
}