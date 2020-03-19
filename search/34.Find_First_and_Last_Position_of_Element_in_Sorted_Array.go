package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	ret := searchRange(nums, target)
	fmt.Println(ret)
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	pos := 0
	for l <= r {
		pos = (l + r) >> 1
		if nums[pos] == target {
			break
		} else if nums[pos] > target {
			r = pos - 1
		} else {
			l = pos + 1
		}
	}
	if nums[pos] != target {
		return []int{-1, -1}
	}
	result := make([]int, 2)
	for i := pos; i >= 0; i-- {
		if nums[i] == target {
			result[0] = i
		} else {
			break
		}
	}
	for i := pos; i < len(nums); i++ {
		if nums[i] == target {
			result[1] = i
		} else {
			break
		}
	}
	return result
}
