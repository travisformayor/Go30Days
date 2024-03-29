// https://www.hackerrank.com/challenges/30-conditional-statements/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)

	switch {
	case N%2 == 1:
		fmt.Printf("Weird")
	case (N%2 == 0) && (N >= 2 && N <= 5):
		fmt.Printf("Not Weird")
	case (N%2 == 0) && (N >= 6 && N <= 20):
		fmt.Printf("Weird")
	case N%2 == 0 && N > 20:
		fmt.Printf("Not Weird")
	default:
		fmt.Printf("Out of Range")
	}
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
