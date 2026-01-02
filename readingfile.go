package main
import ("fmt"
"os"
"io/ioutil")
func main(){
    file,err:=os.Open("file.txt")
	if err!=nil{
		fmt.Println("error",err)
	}
	defer file.Close()
	data,err:=ioutil.ReadAll(file)
	if err!=nil{
		fmt.Println("error",err)
	}
	fmt.Println("file contents are")
	fmt.Println(string(data))
}