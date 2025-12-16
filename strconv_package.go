package main
import ("fmt"
        "strconv")
func main(){
	a:=123
	res:=strconv.Itoa(a)
	/*if err!=nil{
		fmt.Println("error",err)
	}*/
	fmt.Println(res)
	fmt.Printf("%t",res)
}