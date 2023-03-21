// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	ki := int(k)
	a = append(a[len(a)-ki:], a[0:len(a)-ki]...)
	var results []int32
	for i := 0; i < len(queries); i++ {
		results = append(results, a[queries[i]])
	}
	return results
}

func main() {
	a := []int32{1, 2, 3, 4, 5}
	k := int32(3)
	queries := []int32{2, 3}
	fmt.Println(circularArrayRotation(a, k, queries))
}
