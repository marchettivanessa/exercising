// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func aVeryBigSum(ar []int64) int64 {
	sum := int64(0)

	for _, v := range ar {
		sum += v
	}
	return sum
}

func main() {
	arr := []int64{1, 3, 5, 7}
	fmt.Println(aVeryBigSum(arr))
}
