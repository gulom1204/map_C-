package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if num == 10 {
		fmt.Println(22)
	} else if num < 10 {
		fmt.Println(num * 2)
	} else {
		fmt.Println(num + 3)
	}
}
