package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
		hashNum := make(map[int]int)
		for key, value := range nums {
				com := target - value
				if index, ok := hashNum[com]; ok {
						return []int{index, key}
				}
				hashNum[value] = key
		}
		return []int{}
}

func main() {
		example1 := twoSum([]int{2, 7, 11, 15}, 9)
		fmt.Printf("example1 result: %v\n", example1)

		example2 := twoSum([]int{3, 2, 4}, 6)
		fmt.Printf("example2 result: %v\n", example2)

		example3 := twoSum([]int{3, 3}, 6)
		fmt.Printf("example3 result: %v\n", example3)
}
