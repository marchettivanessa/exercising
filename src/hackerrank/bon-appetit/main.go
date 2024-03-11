package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	var soma int32 = 0
	for ind, v := range bill {
		if int32(ind) == k {
			continue
		} else {
			soma += v
		}
	}
	individualValue := soma / 2
	if b == individualValue {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - individualValue)
	}
}

func main (){

	myBill := []int32{1, 2, 3, 4, 5}
	bonAppetit(myBill, 2, 6)

}