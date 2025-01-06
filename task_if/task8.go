package main
import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a % 5 == 0 {
		fmt.Println("true")
	}else {
		fmt.Println("false")
	}
}
