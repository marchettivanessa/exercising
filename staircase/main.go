package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	var i int32 = 0

	for i = 1; i <= n; i++ {
		var n_1 int32 = n - i
		var nEsp int32 = 1
		var nHash int32 = 1
		for nEsp <= n_1 {
			fmt.Print(" ")
			nEsp++
		}
		for nHash <= i {
			fmt.Print("#")
			nHash++
		}
		fmt.Println()
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
