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

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			r := Opcode(p, noun, verb)
			if r == 19690720 {
				fmt.Println(100*noun + verb)
				break
			}
		}
	}
}

func Opcode(program []int, noun, verb int) int {
	// Copy the program into memory, so that we do not modify the original.
	memory := make([]int, len(program))
	copy(memory, program)

	// Copy inputs into memory.
	memory[1], memory[2] = noun, verb

	for i := 0; i < len(memory); i += 4 {
		opcode := memory[i]

		if opcode == 1 {
			a, b, c := memory[i+1], memory[i+2], memory[i+3]
			memory[c] = memory[a] + memory[b]
		} else if opcode == 2 {
			a, b, c := memory[i+1], memory[i+2], memory[i+3]
			memory[c] = memory[a] * memory[b]
		} else if opcode == 99 {
			break
		}
	}

	return memory[0]
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
