package main

import "fmt"

func main() {
	board := compileBoardFromSeed()
	result := isValidSudoku(board)
	fmt.Println(result)
}

func isValidSudoku(board [][]byte) bool {
	columnCheck := true
	rowCheck := true
	for i := 0; i < 9; i++ {
		columnCheck = checkColumns(board, i)
		rowCheck = checkRows(board, i)
		if !columnCheck || !rowCheck {
			return false
		}
	}
	return check3to3(board)
}

func checkRows(board [][]byte, index int) bool {
	myMap := map[byte]byte{}
	for i := 0; i < 9; i++ {
		val := board[index][i]
		if val != 0 {
			_, ok := myMap[val]
			if ok {
				return false
			}
			myMap[val] = board[index][i]
		}
	}
	return true
}

func checkColumns(board [][]byte, index int) bool {
	myMap := map[byte]byte{}
	for i := 0; i < 9; i++ {
		val := board[i][index]
		if val != 0 {
			_, ok := myMap[val]
			if ok {
				return false
			}
			myMap[val] = board[i][index]
		}
	}
	return true
}

func check3to3(board [][]byte) bool {
	for k := 0; k < 3; k++ {
		myMap := map[byte]byte{}
		for i := 0; i < 9; i++ {
			if i%3 == 0 {
				myMap = map[byte]byte{}
			}
			for j := k * 3; j < 3+k*3; j++ {
				fmt.Print(board[i][j])
				val := board[i][j]
				if val != 0 {
					_, ok := myMap[val]
					if ok {
						return false
					}
					myMap[val] = board[i][j]
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	return true
}

func compileBoardFromSeed() [][]byte {
	seed := []rune{
		'.', '.', '.', '.', '5', '.', '.', '1', '.',
		'.', '4', '.', '3', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '3', '.', '.', '1',
		'8', '.', '.', '.', '.', '.', '.', '2', '.',
		'.', '.', '2', '.', '7', '.', '.', '.', '.',
		'.', '1', '5', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '2', '.', '.', '.',
		'.', '2', '.', '9', '.', '.', '.', '.', '.',
		'.', '.', '4', '.', '.', '.', '.', '.', '.'}

	var board [][]byte
	for i := 0; i < 9; i++ {
		row := make([]byte, 9)
		for j := 0; j < 9; j++ {
			val := seed[i*9+j]
			if val != '.' {
				row[j] = byte(val) - '0'
			}
		}
		board = append(board, row)
	}
	return board
}
