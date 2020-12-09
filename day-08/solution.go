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

type Instruction struct {
	Line  int
	ins   string
	val   int
	visit bool
}

func main() {
	file, err := os.Open("day-08/input.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	var ins []Instruction
	i := 0
	for scanner.Scan() {
		s := scanner.Text()
		//fmt.Println(s)
		v, _ := strconv.Atoi(s[4:])
		in := Instruction{
			Line:  i,
			ins:   s[0:3],
			val:   v,
			visit: false,
		}
		ins = append(ins, in)
		i++
	}
	//fmt.Println(len(ins))
	accumlator := 0
	end := false
	prev := new(Instruction)
	i = 0
	for end != true {
		switch ins[i].ins {
		case "acc":
			if ins[i].visit {
				end = true
			} else {
				accumlator = accumlator + ins[i].val
				ins[i].visit = true
				i++
			}
			break
		case "jmp":
			if ins[i].visit {
				fmt.Println(accumlator)
				end = true
			} else {
				ins[i].visit = true
				i = i+ ins[i].val
			}
			prev=  &ins[i]
			break
		case "nop":
			if ins[i].visit {
				fmt.Println(accumlator)
				end = true
			} else {
				ins[i].visit = true
				i++
			}
			prev=  &ins[i]
			break
		}
	}
	fmt.Println(accumlator)
	fmt.Println(len(ins))
	fmt.Println(prev)

}
