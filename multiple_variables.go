package main
import "fmt"
func main(){
	var a,b int=1,2
	c,d:="hello","world"
	var (
		x int
		y string="go developer"
		z bool
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	//naming rules
	fmt.Println("camel case")
	fmt.Println("camel case:each word start with a capital letter except the first word")
	camelCaseWord:="camelCase"
	fmt.Println(camelCaseWord)
	fmt.Println("pascal case")//eachword start with a capital letter
	fmt.Println("eachword start with a capital letter")
	PascalCaseWord:="PascalCase"
	fmt.Println(PascalCaseWord)
    fmt.Println("snake case") //each word seperated by the underscore
	fmt.Println("each word seperated by the underscore")
	snake_case_word:="snake_case"
	fmt.Println(snake_case_word)
}