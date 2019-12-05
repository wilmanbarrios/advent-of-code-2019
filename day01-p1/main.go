package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sum := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		sum += CalculateFuel(n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}

func CalculateFuel(mass int) int {
	return int(math.Floor(float64(mass/3)*100)/100 - float64(2))
}
