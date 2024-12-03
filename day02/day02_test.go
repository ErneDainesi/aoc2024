package day02

import (
	"os"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
    expectedValue := 421
    input, error := os.ReadFile("input.txt")
    if error != nil {
        t.Fatalf("error while reading file")
    }
    parsedInput := strings.Split(string(input), "\n")
    result := PartOne(parsedInput)
    if result != expectedValue {
        t.Fatalf("[PART 1] WRONG! got %d expected %d", result, expectedValue)
    }
}


func TestPartTwo(t *testing.T) {
    expectedValue := 4
    input, error := os.ReadFile("input2.txt")
    if error != nil {
        t.Fatalf("error while reading file")
    }
    parsedInput := strings.Split(string(input), "\n")
    result := PartTwo(parsedInput)
    if result != expectedValue {
        t.Fatalf("[PART 2] WRONG! got %d expected %d", result, expectedValue)
    }
}
