package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Bag struct {
	Colour string
	Count int
	Contains []Bag
}

func main() {
	file, err := os.Open("day-07/input.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	var bags []Bag
	bMap := make(map[string][]Bag, 1000)
	for scanner.Scan() {
		s := scanner.Text()
		b:= strings.Split(s, " bags contain ")
		d:= strings.Split(b[1], " bag, ")
		//fmt.Printf("%v\n", a)
		//fmt.Printf("%v\n", b)
		//fmt.Printf("%v\n", c)
		//fmt.Printf("%v\n", d)

		bag:= Bag{
			Colour:   b[0],
			Count:    0,
			Contains: []Bag{},
		}
		for _, ba:= range d{
			count,_:= strconv.Atoi(ba[0:1])
			bg:= Bag{
				Count: count,
				Colour: ba[2:],
			}
			bag.Contains = append(bag.Contains, bg)
		}
		bags = append(bags, bag)
		bMap[bag.Colour] = bag.Contains
	}
	fmt.Println(len(bags))


}
