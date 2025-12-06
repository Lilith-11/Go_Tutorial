package main
import "fmt"
func main(){
	x:=10
	if x<10{
		fmt.Println("x is less than 10")
	}else if x==10{
		fmt.Println("x is equal to 10")
	}else{
		fmt.Println("x is grater than 20")
	}
	var age int
	fmt.Print("enter your age:")
	fmt.Scan(&age)
	if age<10{
		fmt.Println("enter age less than 10 not applicable to vote")
	}else if age<18{
		fmt.Println("not elligible to vote")
	}else{
		fmt.Println("elligible to vote")
	}
	var year int
	fmt.Print("enter your year")
	fmt.Scan(&year)
	if year%400==0{
		fmt.Println("leap year")
	}else  if year%100==0{
		fmt.Println("not a leap year")
	}else if year%4==0{
        fmt.Println("leap year")
	}else{
        fmt.Println("not a leap year")
	}
}