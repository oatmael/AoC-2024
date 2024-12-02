package main

import "testing"

func TestPart1(t *testing.T) {
	testGroup1 := []int{3, 4, 2, 1, 3, 3};
	testGroup2 := []int{4, 3, 5, 3, 9, 3};

	testAnswer1, _ := processGroups(testGroup1, testGroup2);
	if (testAnswer1 != 11) {
		t.Fatalf(`Got %d, expected 11`, testAnswer1);
	}
}

func TestPart2(t *testing.T) {
	testGroup1 := []int{3, 4, 2, 1, 3, 3};
	testGroup2 := []int{4, 3, 5, 3, 9, 3};

	_, testAnswer2 := processGroups(testGroup1, testGroup2);
	if (testAnswer2 != 31) {
		t.Fatalf(`Got %d, expected 31`, testAnswer2);
	}
}

func TestBenchmark(t *testing.T) {
	main();
}