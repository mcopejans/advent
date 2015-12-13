package main

import (
        "bufio"
        "fmt"
        "github.com/cznic/mathutil"
        "math"
        "os"
        "sort")

type CityPair struct {
    city1, city2 string
}

func AddCityCost(cities CityPair, distance int, city_dist* map[CityPair]int) {
    (*city_dist)[cities] = distance
    (*city_dist)[CityPair{cities.city2, cities.city1}] = distance
}

func LengthOfCombination(permutation []int, cities []string, city_dist map[CityPair]int) int {
    var total_cost int
    for i := 0 ; i < len(permutation) - 1 ; i++ {
        // fmt.Println("City 1 ",cities[permutation[i]])
        // fmt.Println("City 2 ",cities[permutation[i+1]])
        // fmt.Println(city_dist[CityPair{cities[permutation[i]], cities[permutation[i+1]]}])
        total_cost += city_dist[CityPair{cities[permutation[i]], cities[permutation[i+1]]}]
    }

    return total_cost
}

func main() {
    var city_dist = make(map[CityPair]int)
    var cities = []string{"AlphaCentauri",
                          "Snowdin",
                          "Tambi",
                          "Faerun",
                          "Norrath",
                          "Straylight",
                          "Tristram",
                          "Arbre"}

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        var city1, city2 string
        var distance int
        fmt.Sscanf(scanner.Text(), "%s to %s = %d",&city1,&city2,&distance)
        AddCityCost(CityPair{city1,city2},distance,&city_dist)
    }

    fmt.Println(city_dist)

    var data = sort.IntSlice{0, 1, 2, 3, 4, 5, 6, 7}
    i := 0
    var shortest_distance = math.MaxInt64
    var longest_distance = 0
    for mathutil.PermutationFirst(data); ; i++ {
        // fmt.Println(data)
        permutation_distance := LengthOfCombination(data, cities, city_dist)
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
