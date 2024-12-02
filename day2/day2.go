package main

import (
	"fmt"
	"oatmael/aocutils"
	"strconv"
	"strings"
)

func main() {
	var data [][]int;

	parseLine := func(line string) []int {
		row := strings.Split(line, " ");
		return aocutils.ArrayMap(row, func(element string, _ int) int { 
			item, err := strconv.Atoi(element);
			if err != nil {
				return 0;
			}

			return item;
		});
	}

	for line := range aocutils.ReadInput("./input.txt") {
		row := parseLine(line);
		data = append(data, row);
	}

	safeCount := processSafeCount(data);
	fmt.Println("Answer (Part 1): " + strconv.Itoa(safeCount));

}

func processSafeCount(data [][]int) int {

	safeCount := 0;
	for _, row := range data {
		isIncreasing := false;
		isDecreasing := false;

		safe := true;

		var prev int;
		for i, next := range row {
			if i == 0 {
				prev = next;
				continue;
			}

			if (prev > next) {
				isDecreasing = true;
			} else if next > prev {
				isIncreasing = true;
			}

			diff := aocutils.Abs(prev - next);

			safe = safe && !(isIncreasing == isDecreasing) && (diff >= 1 && diff <= 3);
			prev = next;
		}

		if safe {
			safeCount += 1;
		}
	}

	return safeCount;
}