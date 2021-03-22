package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	Size int
	Root *Elem
}

type Elem struct {
	Value int
	Next  *Elem
}

func (s *Stack) GetSize() int {
	return s.Size
}

func (s *Stack) Push(v int) {
	s.Root = &Elem{
		Value: v,
		Next:  s.Root,
	}
	s.Size++
}

func (s *Stack) Pop() (res int, err error) {
	if s.Root == nil {
		return 0, fmt.Errorf("Stack is empty")
	}
	res = s.Root.Value
	s.Root = s.Root.Next
	s.Size--
	return
}

func (s *Stack) GetLast() (res int, err error) {
	if s.Root == nil {
		return 0, fmt.Errorf("Stack is empty")
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
	//var line string

	stack := new(Stack)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		switch strings.Fields(line)[0] {
		case "size":
			fmt.Println(stack.GetSize())
		case "push":
			x, err := strconv.Atoi(strings.Fields(line)[1])
			if err != nil {
				break
			}
			stack.Push(x)
			fmt.Println("ok")

		case "pop":
			res, err := stack.Pop()
			if err != nil {
				fmt.Println("error")
				break
			}
			fmt.Println(res)
		case "back":
			res, err := stack.GetLast()
			if err != nil {
				fmt.Println("error")
				break
			}
			fmt.Println(res)
		case "clear":
			stack.Clear()
			fmt.Println("ok")
		case "exit":
			fmt.Println("bye")
			os.Exit(0)
		}
	}
}
