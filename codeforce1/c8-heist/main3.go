package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var bSize []byte
	var bData []byte

	var r = bufio.NewReaderSize(os.Stdin, 1000000)

	bSize, _, _ = r.ReadLine()
	bData, _, _ = r.ReadLine()
	size, _ := strconv.Atoi(string(bSize))
	var arr []int
	for index, data := range strings.Fields(string(bData)) {
		if index >= size {
			break
		}
		val, _ := strconv.Atoi(data)
		arr = append(arr, val)

	}
	max, min := arr[0], arr[0]
	l := len(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	fmt.Println(max - min + 1 - l)
}

//echo -e "5\n5 4 3 2 1" | go run .
