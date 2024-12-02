package day01

import (
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
    expectedValue := 11
    input, error := os.ReadFile("input.txt")
    if error != nil {
        t.Fatalf("error while reading file")
    }
    result, error := PartOne(string(input))
    if error != nil {
        t.Fatalf("WRONG!")
    }

    if result != expectedValue {
        t.Fatalf("WRONG!")
    }
}


func TestPartTwo(t *testing.T) {
    expectedValue := 11
    input, error := os.ReadFile("input2.txt")
    if error != nil {
        t.Fatalf("error while reading file")
    }
    result, error := PartOne(string(input))
    if error != nil {
        t.Fatalf("WRONG!")
    }

    if result != expectedValue {
        t.Fatalf("WRONG!")
    }
}
