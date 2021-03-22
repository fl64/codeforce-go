package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	Size  int
	First *Elem
	Last  *Elem
}

type Elem struct {
	Value int
	Prev  *Elem
}

func (q *Queue) GetSize() int {
	return q.Size
}

func (q *Queue) Push(v int) {
	e := &Elem{
		Value: v,
	}
	if q.Size == 0 {
		q.First = e
		q.Last = e
	} else {
		q.Last.Prev = e
		q.Last = e
	}
	q.Size++
}

func (q *Queue) Pop() (res int, err error) {
	if q.First == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	res = q.First.Value
	q.First = q.First.Prev
	if q.First == nil {
		q.Last = nil
	}
	q.Size--
	return
}

func (q *Queue) GetFront() (res int, err error) {
	if q.First == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	res = q.First.Value
	return
}

func (q *Queue) Clear() {
	for {
		_, err := q.Pop()
		if err != nil {
			break
		}
	}
}

func main() {
	//var line string

	queue := new(Queue)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		switch strings.Fields(line)[0] {
		case "size":
			fmt.Println(queue.GetSize())
		case "push":
			x, err := strconv.Atoi(strings.Fields(line)[1])
			if err != nil {
				break
			}
			queue.Push(x)
			fmt.Println("ok")

		case "pop":
			res, err := queue.Pop()
			if err != nil {
				fmt.Println("error")
				break
			}
			fmt.Println(res)
		case "front":
			res, err := queue.GetFront()
			if err != nil {
				fmt.Println("error")
				break
			}
			fmt.Println(res)
		case "clear":
			queue.Clear()
			fmt.Println("ok")
		case "exit":
			fmt.Println("bye")
			os.Exit(0)
		}
	}
}
