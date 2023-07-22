package main

import "fmt"

type Counter struct {
	hits int
}

func increments(counter *Counter) {
	counter.hits += 1
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increments(counter)
}

func main() {
	counter := Counter{}

	hello := "Hello"
	world := "World!"

	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)

	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)

	replace(&hello, "Hello", &counter)
	fmt.Println(hello, world)
	fmt.Println("phrase", phrase)
}
