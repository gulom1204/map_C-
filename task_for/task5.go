package main
import "fmt"

func main() {
	var a, count int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		if i % 2 == 0 {
			count++
		}
	}
	fmt.Println(count)
}
