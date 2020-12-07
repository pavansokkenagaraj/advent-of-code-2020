package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// https://adventofcode.com/2020/day/6/input
func main() {
	file, err := os.Open("day-06/input.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	c := make(map[byte]int, 30)
	count := 0
	// part 2
	group := 0
	c2 := make(map[byte]int, 30)
	count2:=0
	for scanner.Scan() {

		b := scanner.Bytes()
		if len(b) != 0 {
			for _, bb := range b {
				if _, ok := c[bb]; !ok {
					c[bb] = 1
					count++
				}
			}

			// part 2
			group++
			for _, bb := range b {
				if _, ok := c2[bb]; !ok {
					c2[bb] = 1
				} else {
					c2[bb] = c2[bb] + 1
				}
			}

		} else {
			c = map[byte]int{}

			// part 2
			for _, value:= range c2{
				if value == group{
					count2++
				}
			}
			group = 0
			c2 = map[byte]int{}
			fmt.Println(count2)

		}
	}
	fmt.Printf("solution part 1: sum of total count = %v\n", count)
	fmt.Printf("solution part 2: sum of total count = %v\n", count2)

	err = file.Close()
	check(err)

}
