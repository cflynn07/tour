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
	m[2] = "b"
	m[3] = "c"
	m[4] = "d"
	m[5] = "e"
	m[6] = "f"
	m[7] = "g"
	m[8] = "h"
	m[9] = "i"
	m[10] = "j"
	m[11] = "k"
	m[12] = "l"
	m[13] = "m"
	//---
	m[14] = "n"
	m[15] = "o"
	m[16] = "p"
	m[17] = "q"
	m[18] = "r"
	m[19] = "s"
	m[20] = "t"
	m[21] = "u"
	m[22] = "v"
	m[23] = "w"
	m[24] = "x"
	m[25] = "y"
	m[26] = "z"

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
