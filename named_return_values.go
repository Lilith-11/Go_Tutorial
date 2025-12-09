package main
import "fmt"
func swap(a,b int)(x int ,y int){
	x=b
	y=a
	return 
}
func main(){
	s,p:=swap(1,2)
	fmt.Println("swapped values",s,p)
}