package main

import "fmt"

func pageCount(n int32, p int32) int32 {
	// Write your code here
	frontPageFlips := p / 2
	if n%2 == 0 {
		n++
	}
	backpageFlip := (n - p) / 2

	if backpageFlip < frontPageFlips {
		return backpageFlip
	}
	return frontPageFlips
}

func main() {
	fmt.Println(pageCount(50, 14))
}
