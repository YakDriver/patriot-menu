// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gonum/stat/combin"
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

func _runPerms(arr []string) (int, []string) {
	perms := permutations(arr)
	//fmt.Println(len(perms))
	bestScore := 5000
	bestPerm := []string{}
	for i := 0; i < len(perms); i++ {
		score := _scorePerm(perms[i])
		if score < bestScore {
			bestScore = score
			bestPerm = perms[i]
		}
	}
	return bestScore, bestPerm
}

func main() {

	allGroups := []string{
		"516-0",
		"564-0",
		"667-0",
		"643-0",
		"656-0",
		"580-0",
	}

	bs, bp := _runPerms(allGroups)
	fmt.Printf("Best score: %v\n", bs)
	fmt.Printf("Best perm: %v\n", bp)

	group0 := []string{
		"120-26",
		"84-42",
		"148-77",
		"148-48",
		"79-45",
		"172-24",
	}

	/*
		--36
		"126-75",
		"122-82",
		"120-82",
		"120-88",
		"123-90",
		"126-90",
		"126-91",
		"125-92",
		"124-97",
		"79-45",

		--76
		"120-56",
		"121-59",
		"118-63",
		"118-67",
		"123-70",
		"125-73",
		"122-80",
		"124-85",
		"95-83",
		"120-26",

		--80
		91-72
		115-71
		114-69
		125-64
		128-76
		128-79
		126-88
		123-90
		122-89
		"84-42",

		--102
		92-56
		100-79
		115-76
		122-93
		128-92
		127-95
		127-94
		132-94
		138-88
		"172-24",

		--114
		"103-71",
		"103-73",
		"122-72",
		"126-77",
		"123-86",
		"128-87",
		"128-89",
		"128-98",
		"174-90",
		"148-77",

		--142
		72-58
		128-64
		129-65
		125-72
		125-76
		130-85
		153-90
		160-79
		160-76
		"148-48",
	*/

	group1 := []string{

		"128-64",
		"125-72",
		"130-85",
	}

	group2 := []string{

		"153-90",
		"129-65",
		"125-76",
	}

	group3 := []string{

		"72-58",

		"160-79",
		"160-76",
	}
	fmt.Printf("Len of group0: %v\n", len(group0))
	fmt.Printf("Len of group1: %v\n", len(group1))
	fmt.Printf("Len of group2: %v\n", len(group2))
	fmt.Printf("Len of group3: %v\n", len(group3))

	combos := combin.Combinations(len(group1), 3)
	fmt.Printf("Combos: %v\n", len(combos))

	bestOverall := 5000
	bestOverallPerm := []string{}
	for i := 0; i < len(combos); i++ {
		fmt.Printf("\r%.2f", float64(i/815.00))
		for j := 0; j < len(combos); j++ {
			for k := 0; k < len(combos); k++ {
				combo := []string{}
				for x := 0; x < 3; x++ {
					combo = append(combo, group1[combos[i][x]])
				}
				for x := 0; x < 3; x++ {
					combo = append(combo, group2[combos[j][x]])
				}
				for x := 0; x < 3; x++ {
					combo = append(combo, group3[combos[k][x]])
				}
				bestScore, bestPerm := _runPerms(combo)
				if bestScore < bestOverall {
					bestOverall = bestScore
					bestOverallPerm = bestPerm
					fmt.Printf("Best overall score: %v\n", bestOverall)
					fmt.Printf("Best overall perm: %v\n", bestOverallPerm)
				}
			}
		}
	}

	fmt.Printf("Best overall score: %v\n", bestOverall)
	fmt.Printf("Best overall perm: %v\n", bestOverallPerm)

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

	bestScore, bestPerm := _runPerms(arr)
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

}
