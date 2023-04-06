// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	records := []int32{}
	greatestScore, lowestScore := int32(scores[0]), int32(scores[0])
	frequencyMax, frequencyMin := int32(0), int32(0)
	for i := 1; i < len(scores); i++ {
		if scores[i] > greatestScore {
			greatestScore = scores[i]
			frequencyMax += 1
		} else if scores[i] < lowestScore {
			lowestScore = scores[i]
			frequencyMin += 1
		}
	}
	records = append(records, frequencyMax, frequencyMin)
	return records
}

func main() {
	fmt.Println(breakingRecords([]int32{22, 12, 33, 14, 5, 67}))
}
