package main

import "fmt"

func migratoryBirds(arr []int32) int32 {
	frequency := map[int32]int32{
		arr[0]: 1,
	}

	bird := arr[0]

	for i := 1; i < len(arr); i++ {
		currentBird := arr[i]
		frequency[currentBird] += 1

		if frequency[currentBird] > frequency[bird] || frequency[currentBird] == frequency[bird] && currentBird < bird {
			bird = currentBird
		}
	}
	return bird
}

func main() {
	fmt.Println(migratoryBirds([]int32{1, 4, 7, 4, 4, 1}))
	fmt.Println(migratoryBirds([]int32{1, 4, 7, 2, 3,2}))
}