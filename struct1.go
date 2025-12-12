package main
import "fmt"
type Person struct{
	name string
	age int
	job string
}
func main(){
   var pers1 Person
   //var pers2 Person
   pers1.name="lilith"
   pers1.age=23
   pers1.job="go develop" 
   person1(pers1) 
   var pers2 Person
   pers2.name="Kannan"
   pers2.age=20
   pers2.job="BBA"
   person2(pers2)
}
func person1(pers Person){
   fmt.Println(pers.name)
}
func person2(pers Person){
   fmt.Println(pers.name)
}