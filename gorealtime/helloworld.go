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
	fmt.Println(Solution2(3, "1A 3C 2B 2G 5A"))
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

func Solution2(n int, s string) int {
	reserved := make(map[int]map[string]bool)

	if len(s) > 0 {
		seats := strings.Split(s, " ")
		fmt.Println(seats)
		for _, seat := range seats {
			// fmt.Println("1")
			row := seat[:len(seat)-1]
			col := seat[len(seat)-1:]
			rowint, _ := strconv.Atoi(row)
			// fmt.Println(row, " ", col, " ", rowint)

			if _, exist := reserved[rowint]; !exist {
				reserved[rowint] = make(map[string]bool)
			}

			reserved[rowint][col] = true
		}
	}

	for key, value := range reserved {
		fmt.Println(key)
		// fmt.Println(value)
		for innerkey, innerval := range value {
			fmt.Println(innerkey)
			fmt.Println(innerval)
		}
	}

	totalEmptySeats := 0

	for row := 1; row <= n; row++ {

		//loop each row to check if the row got any reserved seats
		if _, exist := reserved[row]; !exist {
			totalEmptySeats += 2
			continue
		}

		block1 := !reserved[row]["A"] && !reserved[row]["B"] && !reserved[row]["C"] && !reserved[row]["D"]
		block2 := !reserved[row]["C"] && !reserved[row]["D"] && !reserved[row]["F"] && !reserved[row]["G"]
		block3 := !reserved[row]["F"] && !reserved[row]["G"] && !reserved[row]["H"] && !reserved[row]["J"]

		if block1 && block3 {
			totalEmptySeats += 2
		} else if block1 || block2 || block3 {
			totalEmptySeats++
		}

	}

	return totalEmptySeats
}
