package main

import (
        "bufio"
        "fmt"
        "os")

type Info struct {
    speed, fly, rest int
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var reindeer, fastest_reindeer string
    var speed, fly, rest int
    var reindeers = make(map[string]Info)
    var points = make(map[string]int)
    var distances = make(map[string]int)
    var resets = make(map[string]int)
    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",&reindeer,&speed, &fly, &rest)
        reindeers[reindeer] = Info{speed, fly, rest}
        points[reindeer] = 0
        distances[reindeer] = 0
    }

    // reindeers = make(map[string]Info)
    // points = make(map[string]int)
    // distances = make(map[string]int)
    // resets = make(map[string]int)

    // reindeers["Comet"] = Info{14, 10, 127}
    // reindeers["Dancer"] = Info{16, 11, 162}
    // points["Comet"] = 0
    // points["Dancer"] = 0
    // distances["Comet"] = 0
    // distances["Dancer"] = 0
    // resets["Comet"] = 0
    // resets["Dancer"] = 0

    var longest_distance int
    for i := 0; i < 2503; i++ {
        for k, v := range reindeers {
            if (resets[k] < v.fly) {
                distances[k] += v.speed
                resets[k]++
            } else if (resets[k] == v.rest + v.fly - 1) {
                // do not increase distance
                resets[k] = 0
            } else {
                resets[k]++
            }
        }

        longest_distance = 0
        for _, v := range distances {
            if v > longest_distance {
                longest_distance = v
            }
        }

        for k, _ := range points {
            if distances[k] == longest_distance {
                points[k]++
            }
        }
    }

    fmt.Println("fastest_reindeer",fastest_reindeer, ", longest distance ",longest_distance)
    fmt.Println("points ", points)
    fmt.Println("distances ", distances)
}
