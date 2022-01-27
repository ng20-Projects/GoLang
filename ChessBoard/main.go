package main

import "fmt"

var chessBoard [8][8]string

func main() {
	chessBoard = [8][8]string{
		{"#", "#", "#", "", "", "", "", "_"},
		{"", "#", "", "#", "", "#", "", "_"},
		{"", "", "#", "", "", "", "", "#"},
		{"#", "", "#", "", "", "#", "", "_"},
		{"#", "", "", "", "", "", "#", ""},
		{"", "", "#", "", "", "#", "_", "#"},
		{"#", "#", "#", "", "", "", "", "_"},
		{"", "", "", "", "#", "", "#", ""},
	}
	for i := 0; i < len(chessBoard); i++ {
		fmt.Print(countRankRow(i), " ")
	}
	fmt.Println()
	for i := 0; i < len(chessBoard); i++ {
		fmt.Print(countRankCol(i), " ")
	}
	fmt.Println()
	fmt.Println(countSq(len(chessBoard)))
	fmt.Println(countOccupied())
}
func countRankRow(rank int) int {
	count := 0
	for _, value := range chessBoard[rank] {
		for _, v1 := range value {
			if string(v1) == "#" {
				count++
			}
		}
	}
	return count
}
func countRankCol(f int) int {
	count := 0
	for i1, _ := range chessBoard {
		if chessBoard[i1][f] == "#" {
			count++
		}
	}
	return count
}
func countSq(n int) int {
	sq := n * n
	return sq
}
func countOccupied() int {
	count := 0
	for i1, _ := range chessBoard {
		for _, value := range chessBoard[i1] {
			if value == "#" {
				count++
			}
		}
	}
	return count
}
