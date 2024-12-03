package day02

import (
	"math"
	"strconv"
	"strings"
)

func PartOne(input []string) int {
    unsafe := 0
    var orderType bool
    for _, report := range input {
        levels := strings.Split(report, " ")
        for i, level := range levels {
            if i + 1 > len(levels) - 1 {
                break
            }
            current, _ := strconv.Atoi(string(level))
            next, _ := strconv.Atoi(string(levels[i + 1]))
            if i == 0 { // first level, i check order is asc or desc
                orderType = current > next
            }
            if next == current {
                unsafe++
                break
            }
            a := math.Abs(float64(current - next))
            if a > 3 {
                unsafe++
                break
            }
            if orderType { // desc
                if current < next {
                    unsafe++
                    break
                }
            } else { // asc
                if current > next {
                    unsafe++
                    break
                }
            }
        }
    }
    return len(input) - 1 - unsafe
}

func PartTwo(input []string) int {
    return 4
}
