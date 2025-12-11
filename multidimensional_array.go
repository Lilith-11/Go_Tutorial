package main
import "fmt"
func main(){
	//declaring multidimensional array : var arr [2][3]int
	//initializing multidimensional array
	var arr = [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println("multidimensional array",arr)
	//accessing multidimensional array
	fmt.Println("2 row 3 element is",arr[1][2])
	//modifying an element in multidimensional array
	arr[1][1]=8
	fmt.Println("after modified an element in multidimensional array",arr)
    //iterating over elements
	fmt.Println("all elements are")
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
           fmt.Print(arr[i][j]," ")
	    }
		fmt.Println()
}
}