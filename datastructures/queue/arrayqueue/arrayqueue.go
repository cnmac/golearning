package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type arrayQueue struct {
	maxSize int
	front   int
	rear    int
	arr     []int
}

func NewArrayQueue(maxSize int) *arrayQueue {
	var aq = new(arrayQueue)
	aq.maxSize = maxSize + 2
	aq.front = 0
	aq.rear = maxSize + 1
	aq.arr = make([]int, maxSize+2)
	return aq
}

func (self *arrayQueue) isFull() bool {
	front := self.front + 1
	if front > self.maxSize-1 {
		front = 0
	}
	return front == self.rear
}

func (self *arrayQueue) isEmpty() bool {
	rear := self.rear + 1
	if rear > self.maxSize-1 {
		rear = 0
	}
	return rear == self.front
}

func (self *arrayQueue) add(n int) bool {
	if self.isFull() {
		fmt.Println("queue is full, can not add data!")
		return false
	}
	self.arr[self.front] = n
	if self.front+1 > self.maxSize-1 {
		self.front = 0
	} else {
		self.front = self.front + 1
	}
	return true
}

func (self *arrayQueue) get() (int, error) {
	if self.isEmpty() {
		fmt.Println("queue is no data!")
		return 0, errors.New("queue is no data")
	}
	if self.rear+1 > self.maxSize-1 {
		self.rear = 0
	} else {
		self.rear = self.rear + 1
	}
	return self.arr[self.rear], nil
}

func (self *arrayQueue) show() {
	for i, val := range self.arr {
		fmt.Printf("arr[%d]=%d\n", i, val)
	}
	fmt.Printf("front: %d rear: %d\n", self.front, self.rear)
}

func (self arrayQueue) showHead() int {
	if self.isEmpty() {
		fmt.Println("Can not to show queque head. queue is empty!")
	}
	return self.arr[self.front]
}

func main() {
	queue := NewArrayQueue(3)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("please input operation symbol")
	lastText := ""
	for input.Scan() {
		text := ""
		if lastText == "a" {
			val, err := strconv.Atoi(input.Text())
			if err != nil {
				fmt.Println("input value error!")
				break
			}
			add := queue.add(val)
			if add {
				fmt.Println("add ok~")
			}
			lastText = ""
		} else {
			text = input.Text()
		}

		switch text {
		case "s":
			queue.show()
		case "a":
			fmt.Println("please input a number:")
			lastText = "a"
		case "g":
			getVal, err := queue.get()
			if err == nil {
				fmt.Printf("get data is: %d\n", getVal)
			}
		case "h":
			head := queue.showHead()
			fmt.Printf("show head is: %d\n", head)
		}
	}
}
