package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, k int) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	tmp := board[i][j]
	board[i][j] = '/'
	res := dfs(board, word, i+1, j, k+1) || dfs(board, word, i-1, j, k+1) || dfs(board, word, i, j+1, k+1) || dfs(
		board, word, i, j-1, k+1,
	)
	board[i][j] = tmp
	return res
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}
