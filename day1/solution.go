package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("day1/input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var num []int64

	for scanner.Scan() {
		text := scanner.Text()
		i, err := strconv.ParseInt(text, 10, 64)
		check(err)
		num = append(num, i)
	}
	file.Close()

	//fmt.Print(num)
	part1(num)
	part2(num)

}

func part1(num []int64) {
	found := false
	fmt.Println(`solution part 1`)
	for i := 0; i < len(num); i++ {
		for j := 1; j < len(num); j++ {
			sum := num[i] + num[j]
			if sum == 2020 {
				fmt.Printf(`%v, %v, solution = %v`, num[i], num[j], num[i]*num[j])
				found = true
				break
			}
			if found == true {
				break
			}
		}
	}
}

func part2(num []int64) {
	fmt.Println("\nsolution part 2")
	found := false
	for i := 0; i < len(num); i++ {
		for j := 1; j < len(num); j++ {
			for k := 0; k < len(num); k++ {
				sum := num[i] + num[j] + num[k]
				if sum == 2020 {
					fmt.Printf(`%v, %v, %v, solution = %v`, num[i], num[j], num[k], num[i]*num[j]*num[k])
					found = true
					break
				}
				if found == true {
					break
				}
			}
			if found == true {
				break
			}
		}
	}
}
