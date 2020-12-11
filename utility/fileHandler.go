package utility

import (
	"io/ioutil"
	"log"
	"strings"
)

//Given a path and given the path leads to a txt-file this will return the content as a string.
func inputToString(path string) string  {

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	return strings.TrimSuffix(string(content), "\n")
}

//Day 10
func Input2019Day1() string {
	return inputToString("2019/assets/day1.txt")
}

//Day 6
func Input2020Day6() string {
	return inputToString("2020/assets/day6.txt")
}

//Day 7
func Input2020Day7() string {
	return inputToString("2020/assets/day7.txt")
}

//Day 8
func Input2020Day8() string {
	return inputToString("2020/assets/day8.txt")
}

//Day 9
func Input2020Day9() string {
	return inputToString("2020/assets/day9.txt")
}

//Day 10
func Input2020Day10() string {
	return inputToString("2020/assets/day10.txt")
}
