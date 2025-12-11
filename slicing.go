package main
import "fmt"

func main(){
	slice:=[]int{1,2,3,4,5}
	//accessing a slice
	fmt.Println("the slice slice value is",slice[0])
    //modifying a slice 
	slice[1]=10
	fmt.Println("after modification ",slice)
	//appending 
	slice=append(slice,6)
	fmt.Println("after appending 6",slice)
	//slice capacity and length
	fmt.Println("capacity of slice",cap(slice))
	fmt.Println("length of slice",len(slice))
	//slicing the slice with range
	slice1:=slice[2:5]
	fmt.Println("after slicing",slice1)
	//iterating over a slice
	for index,value:= range slice{
		fmt.Printf("index:%v value:%d\n",index,value)
	}
	//creating a slice using make
	slice3:=make([]int,4,4)//here the first 4 denotes length and second 4 denotes capacity
	fmt.Println(slice3)
	slice3[0]=1
	slice3[1]=2
	slice3[2]=3
	slice3[3]=4
	fmt.Println("after modifying values in the slice3",slice3)
	//coyping slices
	src:=[]int{1,2,3}
	des:=make([]int,3)
	fmt.Println("before copying src to des",des)
	copy(des,src)
    fmt.Println("after copying src to des",des)
}