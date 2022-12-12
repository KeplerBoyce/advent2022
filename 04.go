package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./inputs/04.txt")
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	count := 0
	count2 := 0
	for _, l := range lines {
		halves := strings.Split(l, ",")
		tempnums := append(strings.Split(halves[0], "-"), strings.Split(halves[1], "-")...)
		var nums []int
		for _, n := range tempnums {
			var temp, _ = strconv.Atoi(n)
			nums = append(nums, temp)
		}
		if nums[0] <= nums[2] && nums[1] >= nums[3] || nums[0] >= nums[2] && nums[1] <= nums[3] {
			count++
		}
		if nums[1] >= nums[2] && nums[3] >= nums[0] {
			count2++
		}
	}
	println(count)
	println(count2)
}
