// https://www.hackerrank.com/challenges/30-hello-world/problem
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println("Hello, World.")
		fmt.Println(input.Text())
	}
}
