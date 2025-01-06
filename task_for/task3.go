package main
import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	for i := 0; i <= a; i++ {
		if i % 2 != 0 {
			b += i
		}
	}
	fmt.Println(b)
}
