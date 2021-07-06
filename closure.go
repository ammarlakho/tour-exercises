package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := 0
	var curr int
	var prev int
	return func() int {
		if i == 0 || i == 1 {
			curr = i
			i += 1
			return i
		}
		temp := curr
		curr += prev
		prev = temp
		return curr
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
