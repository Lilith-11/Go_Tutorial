package main
import ("fmt"
"os")
func main(){
	file,err:=os.Create("name.txt")
	if err!=nil{
		fmt.Println("error",err)
	}
	defer file.Close()
	fmt.Println("file created successfully")
}