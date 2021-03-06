// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Perhaps the most basic file reading task is
	// slurping a file's entire contents into memory.
	dat, err := ioutil.ReadFile("./input.go")

	check(err)
	s := string(dat)

	r := regexp.MustCompile(`(?s)type A struct {(.*?)}`)
	result := r.FindStringSubmatch(s)

	fmt.Printf("%v\n", result[1])

}
