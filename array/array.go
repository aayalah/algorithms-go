package array

import (
	"math"
	"sort"
)

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

type queue [][]int

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	if q[i][0] > q[j][0] {
		return true
	}

	if q[i][0] == q[j][0] {
		return q[i][1] < q[j][1]
	}

	return false
}

func (q queue) Swap(i, j int) {
	tmp := q[i]
	q[i] = q[j]
	q[j] = tmp
}

func reconstructQueue(people [][]int) [][]int {
	q := queue(people)
	sort.Sort(q)

	for i := 1; i < len(people); i += 1 {
		if q[i][1] != i {
			insert(q, i, q[i][1])
		}
	}
	return q
}

func insert(arr [][]int, s, f int) {
	tmp := arr[s]

	for i := s; i > f; i -= 1 {
		arr[i] = arr[i-1]
	}

	arr[f] = tmp
}

func singleNonDuplicate(nums []int) int {
	return recSingleNonDuplicate(nums, 0, len(nums)-1)
}

func recSingleNonDuplicate(nums []int, s, f int) int {
	if s == f {
		return nums[s]
	} else if s > f && f >= 0 {
		return nums[f]
	} else if s > f {
		return nums[s]
	}

	mInd := (f + s) / 2
	mVal := nums[mInd]

	leftInd := mInd - 1
	rightInd := mInd + 1

	lenLeft := 0
	if nums[leftInd] == mVal {
		lenLeft = leftInd - s
	} else {
		lenLeft = mInd - s
	}

	if lenLeft < 0 {
		lenLeft = 0
	}

	lenRight := f - rightInd
	if nums[rightInd] == mVal {
		lenRight = f - rightInd
	} else {
		lenRight = f - mInd
	}

	if lenRight < 0 {
		lenRight = 0
	}

	if nums[leftInd] != mVal && nums[rightInd] != mVal {
		return mVal
	}

	remLeft := math.Mod(float64(lenLeft), 2.0)
	remRight := math.Mod(float64(lenRight), 2.0)

	if remRight > 0 {
		return recSingleNonDuplicate(nums, f-lenRight+1, f)
	} else if remLeft > 0 {
		return recSingleNonDuplicate(nums, s, s+lenLeft-1)
	}
	return -1
}
