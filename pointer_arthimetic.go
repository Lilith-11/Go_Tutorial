package main 
import "fmt"
func main(){
	arr:=[5]int{1,2,3,4,5}
	arr1 :=&arr[0]
	array1:=*arr1
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(array1)
	slice :=arr[1:3]
	fmt.Println(slice)
	type person struct{
		name string
		age int
	}
	var p *person=&person{name:"lilith",age:18}
	fmt.Println(p.age) 
    //pointer safety in go
	nums:=[]int{1,2,3,4,5}
	nums=append(nums,6)
	fmt.Println(nums)
}