package main

import "fmt"

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		sumMap[i] = nums[i]
	}

	fmt.Println(sumMap)

	return []int{}
}
