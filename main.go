/**
 * Silly little program to help learn go and memorize integer
 * indexes of roman alphabet characters
 */
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// map of alphabet indexes
var m map[int]string

func randInt(min, max int) int {
	// +1 makes [min, max] <== inclusive
	return min + rand.Intn(max-min+1)
}

// http://golang.org/pkg/bufio/#example_Scanner_lines
func main() {

	// why can I not assign this outside of a function
	m = make(map[int]string)
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

	promptIToC := "Enter index of %s: "
	promptCToI := "Enter character at index %d: \n"

	scanner := bufio.NewScanner(os.Stdin)
	for true {

		// seed random generator
		rand.Seed(time.Now().UTC().UnixNano())
		i := randInt(1, 26)

		if (randInt(0, 1)) > 0 {
			// echo's character, asks for index
			fmt.Printf(promptIToC, m[i])
			scanner.Scan() // gets input/index
			responseString := scanner.Text()
			// convert string to int
			responseInt, _ := strconv.ParseInt(responseString, 10, 0)
			if int(responseInt) == i {
				fmt.Printf("CORRECT. %d is the index of %s \n", i, m[i])
			} else {
				fmt.Printf("WRONG. %d is the index of %s \n", int(responseInt), m[int(responseInt)])
			}
		} else {
			fmt.Printf(promptCToI, i)
			scanner.Scan()
			responseString := scanner.Text()
			if m[i] == responseString {
				fmt.Println("CORRECT. %s is the character with index %d", responseString, i)
			} else {
				fmt.Println("WRONG. %s is the character with index %d", m[i], i)
			}
		}
	}

}
