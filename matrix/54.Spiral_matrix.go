package matrix

/*
right: x,y=+1, until[x][y] unavailable, than turn down
down: x=+1,y, until[x][y] unavailable, than turn left
left: x,y-=1, until[x][y] unavailable, than turn up
up:  x-=1,y, until[x][y] unavailable, than turn right
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	unavailableX, unavailableY := make([]bool, len(matrix)), make([]bool, len(matrix[0]))
	x, y, stepX, stepY := 0, -1, 0, 1
	result, idx := make([]int, len(matrix)*len(matrix[0])), 0
	for idx < len(result) {
		for isAvailable(x+stepX, y+stepY, unavailableX, unavailableY) {
			x += stepX
			y += stepY
			result[idx] = matrix[x][y]
			idx++
		}
		switch {
		case stepX == 0 && stepY == 1: //right -> down
			unavailableX[x] = true
			stepX, stepY = 1, 0
		case stepX == 1 && stepY == 0: //down -> left
			unavailableY[y] = true
			stepX, stepY = 0, -1
		case stepX == 0 && stepY == -1: //left -> up
			unavailableX[x] = true
			stepX, stepY = -1, 0
		case stepX == -1 && stepY == 0: //up -> right
			unavailableY[y] = true
			stepX, stepY = 0, 1
		}
	}
	return result
}

func isAvailable(x, y int, unavailableX, unavailableY []bool) bool {
	if x < 0 || x >= len(unavailableX) {
		return false
	}
	if y < 0 || y >= len(unavailableY) {
		return false
	}
	if unavailableX[x] || unavailableY[y] {
		return false
	}
	return true
}
