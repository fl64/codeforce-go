package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack struct {
	Size int
	Root *Elem
}

type Elem struct {
	Value string
	Next  *Elem
}

func (s *Stack) GetSize() int {
	return s.Size
}

func (s *Stack) Push(v string) {
	s.Root = &Elem{
		Value: v,
		Next:  s.Root,
	}
	s.Size++
}

func (s *Stack) Pop() (res string, err error) {
	if s.Root == nil {
		return "", fmt.Errorf("Stack is empty")
	}
	res = s.Root.Value
	s.Root = s.Root.Next
	s.Size--
	return
}

func (s *Stack) GetLast() (res string, err error) {
	if s.Root == nil {
		return "", fmt.Errorf("Stack is empty")
	}
	res = s.Root.Value
	return
}

func (s *Stack) Clear() {
	for {
		_, err := s.Pop()
		if err != nil {
			break
		}
	}
}

func main() {

	brackets := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	var line string

	stack := new(Stack)

	const maxCapacity = 100002
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	for scanner.Scan() {
		line = scanner.Text()
		break
	}
	//fmt.Println(line)
	if len(line) > 100000 {
		fmt.Println("no")
		os.Exit(0)
	}
	for _, i := range line {
		s := string(i)
		if strings.Contains("([{", s) {
			stack.Push(s)
		}
		if _, ok := brackets[s]; ok {
			s0, err := stack.Pop()
			if err != nil {
				fmt.Println("no")
				os.Exit(0)
			}
			//fmt.Println(s, s0, brackets[s], stack.GetSize())
			if s1, ok := brackets[s]; !ok || s1 != s0 {
				fmt.Println("no")
				os.Exit(0)
			}
		}
	}
	if stack.GetSize() != 0 {
		fmt.Println("no")
		os.Exit(0)
	}
	fmt.Println("yes")
}

// echo "[({}xxas)]" | go run .
