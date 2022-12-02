package day1_test

import (
	"testing"

	day1 "github.com/jalvaydev/advent-of-code/day_1"
)

func TestFindTopCalorieElf(t *testing.T) {
	got := day1.FindTopCalorieElf("./input_test")
	want := 24000
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestFindTopKCalorieElves(t *testing.T) {
	got := day1.FindTopKCalorieElves("./input_test", 3)
	want := 45000
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
