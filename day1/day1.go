package main

import (
	"fmt"
	"log"
	"oatmael/aocutils"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var group1 []int;
	var group2 []int;

	parseLine := func(line string) (int, int, error) {
		locationIds := strings.Split(line, "   ");
		
		id1, err := strconv.Atoi(locationIds[0]);
		if err != nil { return 0, 0, err; }

		id2, err := strconv.Atoi(locationIds[1]);
		if err != nil { return 0, 0, err; }

		return id1, id2, nil;
	};

	for line := range aocutils.ReadInput("./input.txt") {
		id1, id2, err := parseLine(line);
		if err != nil {
			log.Fatal(err);
		}

		group1 = append(group1, id1);
		group2 = append(group2, id2);
	}

	slices.SortFunc(group1, func (a, b int) int { return a - b; });
	slices.SortFunc(group2, func (a, b int) int { return a - b; });

	if (len(group1) != len(group2)) {
		panic("Input lists are not of equal length!");
	}

	var distances []int;
	for i, id1 := range group1 {
		id2 := group2[i];
		distances = append(distances, aocutils.Abs(id1 - id2));
	}

	answer1 := aocutils.ArrayReduce(distances, 0, func(element, _, acc int) int { return element + acc });
	fmt.Println("Answer (Part 1): " + strconv.Itoa(answer1));

	var similarities []int;
	for _, id1 := range group1 {
		duplicates := aocutils.ArrayFilter(group2, func(id2, _ int) bool { return id1 == id2; });
		similarities = append(similarities, id1 * len(duplicates));
	}

	answer2 := aocutils.ArrayReduce(similarities, 0, func(element, _, acc int) int { return element + acc });
	fmt.Println("Answer (Part 2): " + strconv.Itoa(answer2));
}