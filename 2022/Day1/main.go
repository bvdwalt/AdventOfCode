package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var elfsCalories []int64
	file, err := os.Open("/Users/bennie/Library/Application Support/JetBrains/GoLand2022.3/scratches/AoC/2022/Day1/data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	elfIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elfIndex++
			continue
		}
		num, err := strconv.ParseInt(line, 10, 32)
		fmt.Println(err)
		if len(elfsCalories) > elfIndex {
			elfsCalories[elfIndex] += num
		} else {
			elfsCalories = append(elfsCalories, num)
		}
	}

	sort.Slice(elfsCalories, func(i, j int) bool {
		return elfsCalories[i] > elfsCalories[j]
	})
	fmt.Println(fmt.Sprintf("%v \nSlice Properties: %v", elfsCalories, printSlice(elfsCalories)))
	fmt.Println(fmt.Sprintf("Highest Calories: %v\nTop3 Calories Added: %v", elfsCalories[0], elfsCalories[0]+elfsCalories[1]+elfsCalories[2]))
}

func printSlice(slice []int64) string {
	return fmt.Sprintf("len=%d cap=%d size=%d", len(slice), cap(slice), binary.Size(slice))
}
