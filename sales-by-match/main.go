// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func sockMerchant(n int32, ar []int32) int32 {

	var i int32 = 0

	incidencia_cores := make(map[int32]int32)
	for i = 0; i < n; i++ {
		incidencia_cores[ar[i]] = incidencia_cores[ar[i]] + 1
	}

	var pares int32
	for _, cor := range incidencia_cores {
		pares += cor / 2
	}
	return pares
}

func main() {
	n := int32(8)
	ar := []int32{2, 4, 3, 2, 3, 4, 8, 1}
	fmt.Println(sockMerchant(n, ar))
}
