// https://www.hackerrank.com/challenges/30-data-types/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var i2 uint64
	var d2 float64
	var s2 string

	// Read and save an integer, double, and String to your variables.
	// int
	for scanner.Scan() {
		// ParseUint returns multi-values (result and err)
		// Assignments are single value
		// So use a temp var to catch the value
		temp, _ := strconv.ParseUint(scanner.Text(), 10, 64)
		i2 = temp
		break
	}
	// double
	for scanner.Scan() {
		// ParseUint returns multi-values (result and err)
		// Assignments are single value
		// So use a temp var to catch the value
		temp, _ := strconv.ParseFloat(scanner.Text(), 64)
		d2 = temp
		break
	}
	// string
	for scanner.Scan() {
		s2 = scanner.Text()
		break
	}

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + i2)

	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+d2)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Printf(s + s2)
}
