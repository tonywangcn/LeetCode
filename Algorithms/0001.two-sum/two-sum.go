package main

import "fmt"

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:

// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

func twoSum(nums []int, target int) []int {
	res := make(map[int]int, len(nums))
	for k, v := range nums {
		// fmt.Println(target - v)
		if j, ok := res[target-v]; ok {
			return []int{j, k}
		}
		res[v] = k
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 16}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}
