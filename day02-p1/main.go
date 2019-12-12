package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input.txt")

	c := strings.Split(input, ",")
	p := make([]int, len(c))
	for i, n := range c {
		p[i] = toInt(n)
	}

	p[1] = 12
	p[2] = 2
	r := Opcode(p)
	fmt.Println(r)
}

func Opcode(program []int) string {
	for i := 0; i < len(program); i += 4 {
		if program[i] == 99 {
			break
		} else if program[i] == 1 {
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
		} else if program[i] == 2 {
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
		}
	}

	return SplitToString(program, ",")
}

func readFile(file string) string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(bytes))
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return n
}

func SplitToString(a []int, sep string) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, sep)
}
