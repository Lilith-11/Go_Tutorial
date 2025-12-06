package main
import "fmt"
func main(){
	day :=1
	switch day{
	case 1:
		fmt.Println("sunday")
	case 2:
		fmt.Println("monday")
	default :
	    fmt.Println("not a week day")
	}
	workday:=2
	switch workday {
		case 1,2,3:
		fmt.Println("working days in home")
	    case 4,5,6:
		fmt.Println("working days in office")
	    default :
	    fmt.Println("not a working days")}


}