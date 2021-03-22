package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SelectionSort(arr *[]int) {
	arrSize := len(*arr)
	//fmt.Println(arrSize)
	for i := 0; i < len(*arr); i++ {
		maxIndex := 0
		for j := 1; j < arrSize; j++ {
			if (*arr)[maxIndex] < (*arr)[j] {
				maxIndex = j
			}
			if j == arrSize-1 {
				(*arr)[j], (*arr)[maxIndex] = (*arr)[maxIndex], (*arr)[j]
			}
		}
		arrSize--
	}
}

func InsertSort(arr *[]int) {
	for i := 1; i < len(*arr); i++ {
		for j := i; j > 0 && (*arr)[j-1] > (*arr)[j]; j-- {
			(*arr)[j-1], (*arr)[j] = (*arr)[j], (*arr)[j-1]
		}
	}

}

func BubbleSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 1; j < len(*arr)-i; j++ {
			if (*arr)[j-1] > (*arr)[j] {
				(*arr)[j-1], (*arr)[j] = (*arr)[j], (*arr)[j-1]
			}
		}
	}
}

func main() {
	var bSize []byte
	var bData []byte
	var r = bufio.NewReader(os.Stdin)
	bSize, _, _ = r.ReadLine()
	bData, _, _ = r.ReadLine()
	_, _ = strconv.Atoi(string(bSize))
	var arr []int
	for _, data := range strings.Fields(string(bData)) {
		val, _ := strconv.Atoi(data)
		arr = append(arr, val)
	}
	BubbleSort(&arr)
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), " "), "[]"))
}

//echo -e "5\n5 4 3 2 1" | go run .
// https://habr.com/ru/post/204600/
