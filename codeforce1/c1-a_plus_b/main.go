package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var line string
	var sum int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		for _, x := range strings.Fields(line) {
			x, err := strconv.Atoi(x)
			if err != nil {
				break
			}
			sum = sum + x
		}
	}
	fmt.Println(sum)
}
