// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func _scorePerm(arr []string) int {
	k := 0
	for i := 0; i < len(arr)-1; i++ {
		a1 := strings.Split(arr[i], "-")
		a2 := strings.Split(arr[i+1], "-")
		a10, _ := strconv.Atoi(a1[0])
		a11, _ := strconv.Atoi(a2[0])
		a20, _ := strconv.Atoi(a1[1])
		a21, _ := strconv.Atoi(a2[1])
		k += _abs(a10 - a11)
		k += _abs(a20 - a21)
	}
	return k
}

func main() {

	/*arr := []string{
	"100-79",
	"103-71",
	"122-82",
	"123-90",
	"124-85",
	"125-73",
	"126-91",
	"128-92",
	"174-90"}*/
	arr := []string{
		"103-73",
		"91-72",
		"120-82",
		"123-90",
		"128-87",
		"125-92",
		"126-88",
		"128-98",
		"160-79",
	}

	perms := permutations(arr)
	fmt.Println(len(perms))
	bestScore := 5000
	bestPerm := []string{}
	for i := 0; i < len(perms); i++ {
		score := _scorePerm(perms[i])
		if score < bestScore {
			bestScore = score
			bestPerm = perms[i]
		}
	}
	fmt.Printf("Best score: %v\n", bestScore)
	fmt.Printf("Best perm: %v\n", bestPerm)
	//fmt.Println(permutations(arr))

	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it
	// synchronously.
	f("direct")

	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.
	go f("goroutine")

	// You can also start a goroutine for an anonymous
	// function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in
	// separate goroutines now, so execution falls through
	// to here. This `Scanln` requires we press a key
	// before the program exits.
	fmt.Scanln()
	fmt.Println("done")
}
