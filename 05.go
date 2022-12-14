package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./inputs/05.txt")
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	lines = lines[10:]

	stacks1 := [][]byte{
		{'F', 'G', 'V', 'R', 'J', 'L', 'D'},
		{'S', 'J', 'H', 'V', 'B', 'M', 'P', 'T'},
		{'C', 'P', 'G', 'D', 'F', 'M', 'H', 'V'},
		{'Q', 'G', 'N', 'P', 'D', 'M'},
		{'F', 'N', 'H', 'L', 'J'},
		{'Z', 'T', 'G', 'D', 'Q', 'V', 'F', 'N'},
		{'L', 'B', 'x', 'F'},
		{'N', 'D', 'V', 'S', 'B', 'J', 'M'},
		{'D', 'L', 'G'},
	}
	stacks2 := [][]byte{
		{'F', 'G', 'V', 'R', 'J', 'L', 'D'},
		{'S', 'J', 'H', 'V', 'B', 'M', 'P', 'T'},
		{'C', 'P', 'G', 'D', 'F', 'M', 'H', 'V'},
		{'Q', 'G', 'N', 'P', 'D', 'M'},
		{'F', 'N', 'H', 'L', 'J'},
		{'Z', 'T', 'G', 'D', 'Q', 'V', 'F', 'N'},
		{'L', 'B', 'D', 'F'},
		{'N', 'D', 'V', 'S', 'B', 'J', 'M'},
		{'D', 'L', 'G'},
	}
	for _, l := range lines {
		words := strings.Split(l, " ")
		amount, _ := strconv.Atoi(words[1])
		source, _ := strconv.Atoi(words[3])
		target, _ := strconv.Atoi(words[5])
		source--
		target--
		for i := 0; i < amount; i++ {
			temp := make([]byte, 1)
			copy(temp, stacks1[source][:1])
			stacks1[target] = append(temp, stacks1[target]...)
			stacks1[source] = stacks1[source][1:]
		}
		temp := make([]byte, amount)
		copy(temp, stacks2[source][:amount])
		stacks2[target] = append(temp, stacks2[target]...)
		stacks2[source] = stacks2[source][amount:]
	}
	for _, s := range stacks1 {
		fmt.Printf("%c", s[0])
	}
	println()
	for _, s := range stacks2 {
		fmt.Printf("%c", s[0])
	}
	println()
}
