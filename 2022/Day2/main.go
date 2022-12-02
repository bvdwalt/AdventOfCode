package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	myTotalScorePt1, myTotalScorePt2 := RockPaperScissorsScore(lines)
	testScorePt1, testScorePt2 := RockPaperScissorsScore([]string{"A Y", "B X", "C Z"})

	fmt.Println(fmt.Sprintf("MyScore:%v \tPart2:%v\nTest Score:%v\tPart2:%v", myTotalScorePt1, myTotalScorePt2, testScorePt1, testScorePt2))
}

func RockPaperScissorsScore(lines []string) (int, int) {
	scores := map[string][]int{
		"A X": {1 + 3, 3 + 0}, "A Y": {2 + 6, 1 + 3}, "A Z": {3 + 0, 2 + 6},
		"B X": {1 + 0, 1 + 0}, "B Y": {2 + 3, 2 + 3}, "B Z": {3 + 6, 3 + 6},
		"C X": {1 + 6, 2 + 0}, "C Y": {2 + 0, 3 + 3}, "C Z": {3 + 3, 1 + 6},
	}

	part1, part2 := 0, 0
	for _, s := range lines {
		part1 += scores[s][0]
		part2 += scores[s][1]
	}

	return part1, part2
}
