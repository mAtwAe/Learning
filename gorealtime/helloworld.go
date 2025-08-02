package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	// addr := ":8080"
	// http.HandleFunc("/events", events)
	// http.ListenAndServe(addr, nil)

}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	tokens := []string{"Hello", "World", "from", "Go", "Real-time"}

	for _, token := range tokens {
		content := fmt.Sprintf("data: %s\n\n", string(token))
		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		time.Sleep(time.Millisecond * 420)
	}
}

// multiple customer which is i
// each customer has multiple banks which is presented as j
// need to know the richest customer and their worth
// the input is in 2d array

func richestCustomer(arr [][]int) int {
	maxwealth := 0

	for i := 0; i < len(arr); i++ {
		wealth := 0
		for j := 0; j < len(arr); j++ {
			wealth += arr[i][j]
		}
		if wealth > maxwealth {
			maxwealth = wealth
		}
	}

	return maxwealth
}

func Solution(N int, S string) int {
	// Step 1: Parse reserved seats into a map[int]map[string]bool
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
		}
	}

	total := 0

	// Step 2: Loop through each row
	for row := 1; row <= N; row++ {
		// If no seats are reserved in this row
		if _, exists := reserved[row]; !exists {
			total += 2
			continue
		}

		// Step 3: Check three possible seat blocks
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
