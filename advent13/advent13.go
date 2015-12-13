package main

import (
        "bufio"
        "fmt"
        "github.com/cznic/mathutil"
        "math"
        "os"
        "sort"
        "strings")

type Pair struct {
    one, two string
}

func AddCost(people Pair, happy int, people_happy* map[Pair]int) {
    (*people_happy)[people] = happy
}

func LengthOfCombination(permutation []int, people []string, people_happy map[Pair]int) int {
    var total_cost int
    for i := 0 ; i < len(permutation) - 1 ; i++ {
        total_cost += people_happy[Pair{people[permutation[i]], people[permutation[i+1]]}]
        total_cost += people_happy[Pair{people[permutation[i+1]], people[permutation[i]]}]
    }
    total_cost += people_happy[Pair{people[permutation[0]], people[permutation[len(permutation) - 1]]}]
    total_cost += people_happy[Pair{people[permutation[len(permutation) - 1]], people[permutation[0]]}]

    return total_cost
}

func main() {
    var people_happy = make(map[Pair]int)
    var people = []string{"Alice",
                          "Bob",
                          "Carol",
                          "David",
                          "Eric",
                          "Frank",
                          "George",
                          "Mallory",
                          "Me"}

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        var person1, person2 string
        var happy_value int
        if strings.Contains(scanner.Text(), "gain") {
            fmt.Sscanf(scanner.Text(), "%s would gain %d happiness units by sitting next to %s",&person1,&happy_value,&person2)
            AddCost(Pair{person1,person2},happy_value,&people_happy)
        } else {
            fmt.Sscanf(scanner.Text(), "%s would lose %d happiness units by sitting next to %s",&person1,&happy_value,&person2)
            AddCost(Pair{person1,person2},-happy_value,&people_happy)
        }
    }

    fmt.Println(len(people_happy))

    var data = sort.IntSlice{0, 1, 2, 3, 4, 5, 6, 7, 8}
    i := 0
    var shortest_distance = math.MaxInt64
    var longest_distance = 0
    for mathutil.PermutationFirst(data); ; i++ {
        permutation_distance := LengthOfCombination(data, people, people_happy)
        if (permutation_distance < shortest_distance) {
            shortest_distance = permutation_distance
        }
        
        if (permutation_distance > longest_distance) {
            longest_distance = permutation_distance
        }
        if !mathutil.PermutationNext(data) {
            break
        }
    }

    fmt.Println("Shortest distance",shortest_distance)
    fmt.Println("Longest distance",longest_distance)
}
