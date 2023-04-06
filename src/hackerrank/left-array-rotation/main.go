package main

import "fmt"

func rotateLeft(d int32, arr []int32) []int32 {
	// Write your code here
	var result []int32
	result = append(arr[d:len(arr)], arr[0:d]...)
	return result
}

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	d := int32(3)
	fmt.Println(rotateLeft(d, arr))
}
