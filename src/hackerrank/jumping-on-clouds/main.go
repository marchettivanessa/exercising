package main

import "fmt"

func main() {
	clouds := []int{0, 0, 0, 1, 0, 0, 0, 1, 0}

	fmt.Println(jumpingOnClouds(clouds))
}

func jumpingOnClouds(c []int) int {
	counter := int(0)
	i := 0
	for {
		if i+2 < len(c) && c[i+2] == 0 {
			i += 2

		} else if i+1 < len(c) {
			i++

		} else {
			break
		}
		counter++
	}
	return counter
}
