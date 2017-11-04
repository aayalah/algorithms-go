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
