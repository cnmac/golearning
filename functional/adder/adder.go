package main

import "fmt"

type iAdder func(int) (int, iAdder)

//func adder() func(int) int {
//	sum := 0
//	return func(v int) int {
//		sum += v
//		return sum
//	}
//}

func adder(base int) iAdder {
	return func(addNum int) (int, iAdder) {
		return base + addNum, adder(base + addNum)
	}
}

func main() {
	var adder = adder(0)
	for i := 1; i < 100; i++ {
		var sum int
		sum, adder = adder(i)
		fmt.Printf("0 + ... + %d = %d\n", i, sum)
	}
}
