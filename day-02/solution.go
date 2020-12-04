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

type Password struct {
	Min int
	Max int
	Val string
	Pwd string
}

func main() {
	file, err := os.Open("day-02/input.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	var passwords []Password
	for scanner.Scan() {
		text := scanner.Text()
		// 2-4 p: vpkpp
		a := strings.Split(text, "-")
		b := strings.Split(a[1], " ")
		c := strings.Split(b[1], ":")
		min, err := strconv.Atoi(a[0])
		check(err)
		max, err := strconv.Atoi(b[0])
		check(err)

		pwd := Password{
			Min: min,
			Max: max,
			Val: c[0],
			Pwd: b[2],
		}
		passwords = append(passwords, pwd)
	}
	err = file.Close()
	check(err)

	part1(passwords)
	part2(passwords)

}

func part1(passwords []Password) {
	fmt.Println("\nsolution part 1")
	validPwdCount := 0
	for _, pwd := range passwords {
		count := strings.Count(pwd.Pwd, pwd.Val)

		if count < pwd.Min || count > pwd.Max {
			//fmt.Println(pwd.Val)
		} else {
			validPwdCount++
		}
	}
	fmt.Printf(`total number of valid passwords: %v`, validPwdCount)
}

func part2(passwords []Password) {
	fmt.Println("\nsolution part 2")
	validPwdCount := 0
	for _, pwd := range passwords {

		if string(pwd.Pwd[pwd.Min-1]) == pwd.Val && string(pwd.Pwd[pwd.Max-1]) == pwd.Val {
			//fmt.Println(pwd.Val)
		} else if string(pwd.Pwd[pwd.Min-1]) == pwd.Val || string(pwd.Pwd[pwd.Max-1]) == pwd.Val {
			validPwdCount++

		}
	}
	fmt.Printf(`total number of valid passwords: %v`, validPwdCount)
}
