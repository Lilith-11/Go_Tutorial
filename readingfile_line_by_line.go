package main
import ("os" 
"fmt" 
"bufio"
"log")
func main(){
	file,err:=os.Open("file.txt")
	if err!=nil{
		log.Fatal(err)
		return
	}
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if err:= scanner.Err(); err!=nil{
		log.Fatal(err)
	}
}