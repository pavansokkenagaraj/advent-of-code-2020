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

func main() {
	file, err := os.Open("day-03/input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var f []string
	for scanner.Scan() {
		text := scanner.Text()
		f = append(f, text)
	}
	err = file.Close()
	check(err)

	part1(f)
	part2(f)
}

func part1(f []string) {
	trees := 0
	for i, j := 0, 0; i < len(f); i++ {
		if f[i][j] == '#' {
			trees++
		}
		j = (j + 3) % len(f[i])
	}
	fmt.Printf("solution part 1: number of trees = %v\n", trees)
}

func part2(f []string) {
	treeTotal :=
		traverseRule(f, 1, 1) *
			traverseRule(f, 3, 1) *
			traverseRule(f, 5, 1) *
			traverseRule(f, 7, 1) *
			traverseRule(f, 1, 2)
	fmt.Printf("solution part 2: number of trees = %v\n", treeTotal)
}

func traverseRule(f []string, right, down int) int {
	i, j, trees := 0, 0, 0
	for i < len(f) {
		if f[i][j] == '#' {
			trees++
		}
		j = (j + right) % len(f[i])
		i += down
	}
	return trees
}
