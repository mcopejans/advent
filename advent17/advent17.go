package main

import("fmt")

func Btoi(b bool) int {
    if b {
        return 1
    }
    return 0
 }

func CombinationGenerator(liter int, containers []int) [][]int {
	if (liter == 0) {
		fmt.Println("11")
		return make([][]int, 1)
	}

	if (len(containers) == 0 || liter < 0) {
		fmt.Println("22")
		return [][]int{}
	}

	combinations := CombinationGenerator(liter - containers[0], containers[1:len(containers)])

	for combination := range combinations {
		combinations[combination] = append(combinations[combination], containers[0])
	}

	combinations2 :=  CombinationGenerator(liter, containers[1:len(containers)])

	return append(combinations, combinations2...)
}

func CombinationGenerator2(liter int, containers []int) [][]int {
	if (liter == 150) {
		fmt.Println("1")
		return make([][]int, 1)
	}

	if (len(containers) == 0 || liter > 150) {
		fmt.Println("2")
		return [][]int{}
	}

	combinations := CombinationGenerator2(liter + containers[0], containers[1:len(containers)])

	for combination := range combinations {
		combinations[combination] = append(combinations[combination], containers[0])
	}

	combinations2 :=  CombinationGenerator2(liter, containers[1:len(containers)])

	return append(combinations, combinations2...)
}

func main () {
	// var containers = []int{1, 2, 3, 4}
	var containers = []int{43, 3, 4, 10, 21, 44, 4, 6, 47, 41, 34, 17, 17, 44, 36, 31, 46, 9, 27, 38}

    var combinations = CombinationGenerator(150, containers)
    var combinations2 = CombinationGenerator2(0, containers)
    fmt.Println(combinations2)

    var min = 1000000
    for i := range combinations {
    	if len(combinations[i]) < min {
    		min = len(combinations[i])
    	}
    }

    var number_of_min_container_solutions int
    for i := range combinations {
    	if len(combinations[i]) == min {
    		number_of_min_container_solutions++
    	}
    }

	fmt.Println("Number of combinations ", len(combinations))
	fmt.Println("Number of combinations ", len(combinations2))
	fmt.Println("number_of_min_container_solutions", number_of_min_container_solutions)
}