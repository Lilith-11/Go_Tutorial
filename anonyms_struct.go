package main
import "fmt"

func main(){
	student:=struct{
		name string
		age int
	}{
		name:"lilith",
		age:23,
	}
	fmt.Println(student.name,student.age)
	
}