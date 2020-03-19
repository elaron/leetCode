package main

import "fmt"

func main240()  {
	matrix := [][]int{
		[]int{-1, 3},
	}
	target := 3
	ret := searchMatrix(matrix, target)
	fmt.Println(ret)
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	minL, maxL := len(matrix), len(matrix[0])
	if minL > maxL {
		minL, maxL = maxL, minL
	}
	for i := 0; i <minL ;i++ {
		if matrix[i][i] == target {
			return true
		}else if matrix[i][i] < target {
			continue
		}else{
			for j := 0; j < minL; j++ {
				if matrix[j][i] == target || matrix[i][j] == target{
					return true
				}
			}
		}
	}
	for i := minL; i < maxL; i++ {
		if (minL == len(matrix) && matrix[minL-1][i] >= target){
			if matrix[minL-1][i] == target {
				return true
			}
			for j := 0; j < minL; j++ {
				if matrix[j][i] == target {
					return true
				}
			}
		}else if (minL == len(matrix[0]) && matrix[i][minL-1] >= target) {
			if matrix[i][minL-1]  == target {
				return true
			}
			for j := 0; j < minL; j++ {
				if matrix[i][j] == target {
					return true
				}
			}
		}
	}
	return false
}
