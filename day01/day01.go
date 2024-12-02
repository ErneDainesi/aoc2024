package day01

import (
	"bytes"
	"sort"
	"strconv"
)

func PartOne(input []byte) int {
    var left []int
    var right []int
    appendLeft := true
    var buffer bytes.Buffer
    for _, char := range input {
        if char != 32 && int64(char) != 10 {
            buffer.WriteString(string(char))
        } else {
            num, _ := strconv.Atoi(buffer.String())
            if appendLeft {
                left = append(left, num)
            } else {
                right = append(right, num)
            }
            appendLeft = !appendLeft
            buffer.Reset()
        }
    }
    sort.Sort(sort.IntSlice(left))
    sort.Sort(sort.IntSlice(right))
    result := 0
    for pos, val := range left {
        partial := val - right[pos]
        if partial < 0 {
            partial *= -1
        }
        result += partial
    }
    return result
}

func PartTwo(input string) int {
    var left []int
    var right []int
    appendLeft := true
    var buffer bytes.Buffer
    for _, char := range input {
        if char != 32 && int64(char) != 10 {
            buffer.WriteString(string(char))
        } else {
            num, _ := strconv.Atoi(buffer.String())
            if (num > 0) {
                if appendLeft {
                    left = append(left, num)
                } else {
                    right = append(right, num)
                }
                appendLeft = !appendLeft
                buffer.Reset()
            }
        }
    }
    m := make(map[int]int)
    for _, num := range left {
        if _, ok := m[num]; !ok {
            m[num] = 0
        }
    }
    for _, num := range right {
        if reps, ok := m[num]; ok {
            m[num] = reps + 1
        }
    }
    result := 0
    for _, num := range left {
        if reps, ok := m[num]; ok {
            result += num * reps
        }
    }
    return result
}
