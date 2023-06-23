// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	arr := []int32{0, -2, -1, 5, 6}
	plusMinus(arr)

}

func plusMinus(arr []int32) {
	countNeg := int32(0)
	countPos := int32(0)
	countZer := int32(0)
	i := 0
	for i = 0; i < len(arr); i++ {
		if arr[i] == 0 {
			countZer++
		}
		if arr[i] < 0 {
			countNeg++
		}
		if arr[i] > 0 {
			countPos++
		}
	}
	divisor := float32(i)
	result := []float32{}
	result = append(result, (float32(countPos) / divisor), (float32(countNeg) / divisor), (float32(countZer) / divisor))

	fmt.Printf("%.6f\n", result[0])
	fmt.Printf("%.6f\n", result[1])
	fmt.Printf("%.6f\n", result[2])
}
