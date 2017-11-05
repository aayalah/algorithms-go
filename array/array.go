package array

func countBattleships(board [][]byte) int {
	numShips := len(board)
	height := len(board)
	width := len(board[0])

	for i := 0; i < height; i += 1 {
		for j := 0; j < width; j += 1 {
			if j > 0 && string(board[i][j-1]) == "X" {
				continue
			}

			if i > 0 && string(board[i-1][j]) == "X" {
				continue
			}

			if string(board[i][j]) == "." {
				continue
			}

			numShips += 1
		}
	}
	return numShips
}

func findOriginalValue(val, convNum, len int) int {
	for val > len {
		val -= convNum
	}
	return val
}

func convertNumber(len, num int) int {
	return len + 1 + num
}

func findDuplicates(nums []int) []int {
	for i := 0; i < len(nums); i += 1 {

		val := convertNumber(len(nums), i)
		index := findOriginalValue(nums[i], val, len(nums))
		nums[index] += convertNumber(len(nums), index)
		nums[i] -= index
	}

	res := make([]int, 0)
	for j := 0; j < len(nums); j += 1 {
		val := convertNumber(len(nums), j+1)

		if nums[j]-val > 0 {
			res = append(res, j+1)
		}
	}
	return res

}
