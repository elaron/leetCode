package dp

import "fmt"

func main_canJump() {
	//nums := []int{2,0,2}
	//nums := []int{1,0,1,0}
	nums := []int{3, 2, 1, 0, 4}
	ret := canJump(nums)
	fmt.Println(ret)
}

func canJump(nums []int) bool {
	maxPosition := 0
	for i, num := range nums {
		if i > maxPosition {
			return false
		}
		maxPosition = max(maxPosition, i+num)
	}
	return true
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
