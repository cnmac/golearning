package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type sparseNode struct {
	row int
	col int
	val int
}

func main() {

	// init matrix data
	var matrixArr [13][13]int
	matrixArr[1][2] = 1
	matrixArr[2][3] = 2

	for _, row := range matrixArr {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
	fmt.Println("------------------------------------")

	// convert to sparse array
	var sparseArr []sparseNode
	sparseArr = append(sparseArr, sparseNode{13, 13, 0})
	for i, row := range matrixArr {
		for j, val := range row {
			if val != 0 {
				sparseArr = append(sparseArr, sparseNode{i, j, val})
			}
		}
	}
	// save as a file
	filepath := "datastructures/arr/sparsearr/sparse.in"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file error=%v\n", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	// print
	for i, node := range sparseArr {
		str := fmt.Sprintf("%d %d %d \n", node.row, node.col, node.val)
		writer.WriteString(str)
		fmt.Printf("%d: %d %d %d\n", i, node.row, node.col, node.val)
	}
	writer.Flush()
	fmt.Println("------------------------------------")

	// sparseArr convert to sparseArr initial data
	recoverData(filepath)

}

func recoverData(filepath string) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer file.Close()
	bfrd := bufio.NewReader(file)
	var index = 0
	var arr [][]int
	for {
		line, err := bfrd.ReadBytes('\n')
		if err != nil {
			break
		}
		index++
		temp := strings.Split(string(line), " ")
		row, _ := strconv.Atoi(temp[0])
		col, _ := strconv.Atoi(temp[1])
		val, _ := strconv.Atoi(temp[2])
		if index == 1 {
			for i := 0; i < row; i++ {
				tmparr := []int{}
				for j := 0; j < col; j++ {
					tmparr = append(tmparr, val)
				}
				arr = append(arr, tmparr)
			}
		}
		if index != 1 {
			arr[row][col] = val
		}
	}

	// recover data from file
	fmt.Println("recover data from file:")
	for _, row := range arr {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
