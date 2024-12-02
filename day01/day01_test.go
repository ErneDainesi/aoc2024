package day01

import (
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
    expectedValue := 2815556
    input, error := os.ReadFile("input.txt")
    if error != nil {
        t.Fatalf("error while reading file")
    }
    result := PartOne(input)
    if result != expectedValue {
        t.Fatalf("[PART 1] WRONG! got %d expected %d", result, expectedValue)
    }
}


func TestPartTwo(t *testing.T) {
    expectedValue := 23927637
    input, error := os.ReadFile("input2.txt")
    if error != nil {
        t.Fatalf("error while reading file")
    }
    result := PartTwo(string(input))
    if result != expectedValue {
        t.Fatalf("[PART 2] WRONG! got %d expected %d", result, expectedValue)
    }
}
