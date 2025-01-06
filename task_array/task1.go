package main

import "fmt"

func main() {
	var nums = [...]int{12, 21, 36, 71, 24, 85, 12, 30, 25}
	b := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > b {
			b = nums[i]
		}
	}
	fmt.Println(b)
}
