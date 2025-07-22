package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution(N int, S string) int {
	// Implement your solution here
	// n row
	// reserve = 4n or 2n+2n
	// insert all of the seat in to a list
	reserved := make(map[int]map[string]bool)

	if len(S) > 0 {
		seats := strings.Split(S, " ")

		for _, seat := range seats {
			rowStr := seat[:len(seat)-1]
			col := seat[len(seat)-1:]
			row, _ := strconv.Atoi(rowStr)

			if _, exists := reserved[row]; !exists {
				reserved[row] = make(map[string]bool)
			}
			reserved[row][col] = true
			fmt.Println(reserved[row][col])
		}
	}

	total := 0

	for row := 1; row <= N; row++ {
		if _, exists := reserved[row]; !exists {
			total += 2
			continue
		}

		block1 := !reserved[row]["B"] && !reserved[row]["C"] && !reserved[row]["D"] && !reserved[row]["E"]
		block2 := !reserved[row]["D"] && !reserved[row]["E"] && !reserved[row]["F"] && !reserved[row]["G"]
		block3 := !reserved[row]["F"] && !reserved[row]["G"] && !reserved[row]["H"] && !reserved[row]["J"]

		if block1 && block3 {
			total += 2
		} else if block1 || block2 || block3 {
			total += 1
		}
	}

	return total

}

func Solution2(S string) int {
	// Implement your solution here
	//n letters
	//"" is not palindrom
	n := len(S)

	isPalindrom := func(s string) bool {
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
		return true
	}

	if !isPalindrom(S) {
		return 0
	}

	allsame := true
	for i := 1; i < n; i++ {
		if S[i] != S[0] {
			allsame = false
			break
		}
	}
	if allsame {
		return n
	}

	return 1

}
