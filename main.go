/**
 * Silly little program to help learn go and memorize integer
 * indexes of roman alphabet characters
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

// map of alphabet indexes
var m map[int8]string

// http://golang.org/pkg/bufio/#example_Scanner_lines
func main() {
	// why can I not assign this outside of a function
	m = make(map[int8]string)
	m[1] = "a"

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
