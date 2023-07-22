package main

import "fmt"

func main() {
	//fmt.Println("Hello, 世界")
	n := 100000000
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i // sum = sum + i
	}
	fmt.Println("sum =", sum)
}
