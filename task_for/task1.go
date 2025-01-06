package main
import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	for i := 0; i <= a; i++ {
		b += i
	}
	fmt.Println(b)
}
