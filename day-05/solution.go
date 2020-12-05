package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("day-05/input.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	var seats []string
	for scanner.Scan() {
		text := scanner.Text()
		seats = append(seats, text)
		}
	err = file.Close()
	check(err)

	seatIDs := part1(seats)
	part2(seatIDs)
}

func part1(seats []string)[]int  {

	maxID:=0
	var seatIDs []int
	for _,seat:= range seats{
		row := 0
		rp := 64

		col := 0
		cp := 4
		for _, char := range seat {
			switch char {
			case 'F':
				rp/=2
				break
			case 'B':
				row += rp
				rp /= 2
				break
			case 'R':
				col += cp
				cp /= 2
				break
			case 'L':
				cp /= 2
				break
			}
		}
		seatID := row*8+col
		seatIDs = append(seatIDs, seatID)
		if seatID > maxID{
			maxID = seatID
		}
	}
	fmt.Printf("solution part 1: maximum seatID = %v\n", maxID)
	return seatIDs
}

func part2(seatIDs []int) {
	sort.Ints(seatIDs)
	for i := range seatIDs{
		if seatIDs[i+1] != seatIDs[i]+1 {
			fmt.Printf("solution part 1: my seatID = %v\n", seatIDs[i]+1)
			break
		}
	}
}