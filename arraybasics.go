package main
import "fmt"
func main(){
	//declaring an array : var arrayName [size]type
	//declaring an array with 5 integers : var numbers [5]int
	//initializing array
	var numbers =[5]int{1,2,3,4,5}
	fmt.Println(numbers)
	//accesing array elements
	numbers[1]=8
	fmt.Println(numbers[0])
	fmt.Println(numbers)
	//interating over arrays
	for i:=0;i<len(numbers);i++{
		fmt.Println(numbers[i])
	}
	//iterating by using range keyword
	for index,value:=range numbers{
		fmt.Println(index,value)
	}
	//multidimensional arrays
	// declaring 2x3 array var matrix [2][3]int
	var matrix = [2][3]int{{1,2,3},{1,2,3}}
	fmt.Println(matrix)

	
}