package main

import (
	"fmt"
	"os"
)

func checkerRC(board [][]byte, row, col int, num byte) bool {

	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}
	
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if board[r][c] == num {
				return false
			}
		}
	}
	return true
}


func solve(board [][]byte) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				for num := byte('1'); num <= '9'; num++ {
					if checkerRC(board, r, c, num) {
						board[r][c] = num
						if solve(board) {
							return true
						}
						board[r][c] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func countSolutions(board [][]byte, limit int) int {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				count := 0
				for num := byte('1'); num <= '9'; num++ {
					if checkerRC(board, r, c, num) {
						board[r][c] = num
						count += countSolutions(board, limit)
						board[r][c] = '.'
						if count >= limit {
							return count
						}
					}
				}
				return count
			}
		}
	}
	return 1
}


func copyBoard(src [][]byte) [][]byte {
	dest := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		dest[i] = make([]byte, 9)
		copy(dest[i], src[i])
	}
	return dest
}


func printBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", board[i][j])
			if j < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	board := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		arg := os.Args[i+1]
		if len(arg) != 9 {
			fmt.Println("Error")
			return
		}
		row := make([]byte, 9)
		for j := 0; j < 9; j++ {
			ch := arg[j]
			if ch != '.' && (ch < '1' || ch > '9') {
				fmt.Println("Error")
				return
			}
			row[j] = '.'
		}
		board[i] = row
	}

	for i := 0; i < 9; i++ {
		arg := os.Args[i+1]
		for j := 0; j < 9; j++ {
			ch := arg[j]
			if ch != '.' {
				if !checkerRC(board, i, j, ch) {
					fmt.Println("Error")
					return
				}
				board[i][j] = ch
			}
		}
	}

	if countSolutions(copyBoard(board), 2) != 1 {
		fmt.Println("Error")
		return
	}

	if !solve(board) {
		fmt.Println("Error")
		return
	}

	printBoard(board)
}