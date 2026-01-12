package main
import ("fmt"
"os"
)
func main(){
	_,err:=os.Stat("createfile.go")
	if os.IsExist(err){
		fmt.Println("file doesn't exist")
	}else{
		fmt.Println("file exist")
	}
}