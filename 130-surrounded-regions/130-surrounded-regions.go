package main

import "fmt"

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][len(board[0])-1] == 'O' {
			dfs(board, i, len(board[0])-1)
		}
	}

	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' {
			dfs(board, 0, i)
		}
		if board[len(board)-1][i] == 'O' {
			dfs(board, len(board)-1, i)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '1' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}

	board[i][j] = '1'

	dfs(board, i+1, j)
	dfs(board, i-1, j)
	dfs(board, i, j+1)
	dfs(board, i, j-1)
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	solve(board)

	fmt.Println(board)
}
