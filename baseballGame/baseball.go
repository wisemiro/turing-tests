package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type score []int

// Push adds values to score
func (s score) Push(v int) score {
	return append(s, v)
}

// Pop removes values from score
func (s score) Pop() (score, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}
func calPoints(ops []string) int {
	ans := 0
	s := make(score, 0)
	for _, str := range ops {
		ok := true
		round := 0
		num, err := strconv.Atoi(str)
		if err != nil {
			if str == "+" {
				l := len(s)
				round = s[l-2] + s[l-1]
			} else if str == "C" {
				s, num = s.Pop()
				round = -num
				ok = false
			} else if str == "D" {
				l := len(s)
				round = 2 * s[l-1]
			}
		} else {
			round = num
		}
		if ok {
			s = s.Push(round)
		}
		ans += round
	}

	return ans
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')
	rawInput = strings.Replace(rawInput, "\n", "", -1)

	ops := strings.Split(rawInput, " ")
	fmt.Println(calPoints(ops))
}
