package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a > b && a > c {
		fmt.Println(a)
	}else if a < b && c < b {
		fmt.Println(b)
	}else if a < c && b < c {
		fmt.Println(c)
	}else {
		fmt.Println("uchchalasi teng")
	}
}
