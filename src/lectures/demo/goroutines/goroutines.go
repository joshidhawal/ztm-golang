package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune

	capIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
		fmt.Printf("%c done!\n", r)
	}

	for i := 0; i < len(data); i++ {
		go capIt(data[i])
	}

	time.Sleep(300 * time.Millisecond)
	fmt.Printf("Captialized: %c\n", capitalized)
}
