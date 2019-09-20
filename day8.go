// https://www.hackerrank.com/challenges/30-dictionaries-and-maps/problem
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
	phoneBook := make(map[string]string)
	index := 0
	howMany := 0

	// Grab all inputs in loop
	// Handle each input based on what type of input it is.
	// First one is expected # of entries, then the entries, then the lookups
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		switch { // true switch. always run
		case index == 0:
			// first loop. collect number of inputs
			howManyTemp, err := strconv.ParseInt(text, 10, 64)
			checkErr(err)
			howMany = int(howManyTemp)
		case index <= howMany:
			// This is an entry for the phonebook
			// Save input to the phonebook map
			split := strings.Split(text, " ")
			phoneBook[split[0]] = split[1]
		default:
			// All other inputs after the last entry
			// These are the lookup requests
			// Check each one against the phonebook
			number, test := phoneBook[text]
			if test == true {
				fmt.Printf("%s=%s\n", text, number)
			} else {
				fmt.Printf("Not found\n")
			}
		}
		// Increment the index
		index++
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
