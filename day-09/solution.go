package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("day-09/input.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	var tada []int
	for scanner.Scan() {
		s := scanner.Text()
		x, _ := strconv.Atoi(s)
		//fmt.Println(x)
		tada = append(tada, x)

	}
	bad := 0

	for i := 25; i < len(tada); i++ {

		found := false
		for j := i - 25; j < i-1; j++ {
			for k := j + 1; k < i; k++ {
				if tada[j]+tada[k] == tada[i] {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if found == false {
			bad = tada[i]
			break
		}
	}

	weakness := 0
	for i := 0; i < len(tada); i++ {
		sum := 0

		var nums []int
		sum += tada[i]
		nums = append(nums, tada[i])
		found := false
		for j := i + 1; j < len(tada); j++ {
			sum += tada[j]
			nums = append(nums, tada[j])
			if sum == bad {

				sort.Ints(nums)
				weakness = nums[0] + nums[len(nums)-1]
				found = true
				break
			} else if sum > bad {
				break
			}
		}
		if found {
			break
		}
	}

	fmt.Printf("solution part 1: first number that does not have this property = %v\n", bad)
	fmt.Printf("solution part 2: encryption weakness in your XMAS-encrypted list of numbers = %v\n", weakness)
}
