package main 
import "fmt"
type Person struct{
	name string
	age int 
}
func main(){
	var p Person
	p=Person{name:"lilith",age:12}
	fmt.Println(p.name)
	//accessing and modifying the struct field
	p.name="Kannan"
	fmt.Println(p.name)
}