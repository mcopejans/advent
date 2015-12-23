package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Grid[100][100]int

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var lights Grid
    var row int

    for scanner.Scan() {
        // fmt.Println(scanner.Text())
        split := strings.Split(scanner.Text(), "")

        for column := range scanner.Text() {
            if split[column] == "." {
                lights[row][column] = 0
            } else {
                lights[row][column] = 1
            }
        }
        row++
    }

    fmt.Println(lights)

    lights[0][0] = 1
    lights[0][99] = 1
    lights[99][0] = 1
    lights[99][99] = 1
    var updated_lights = lights
    for i := 0; i < 100; i++ {
        for i := range lights {
            for j := range lights[i] {
                var neighbors = make([][]int, 8)
                neighbors[0] = []int{i - 1, j - 1}
                neighbors[1] = []int{i    , j - 1}
                neighbors[2] = []int{i + 1, j - 1}
                neighbors[3] = []int{i - 1, j + 1}
                neighbors[4] = []int{i    , j + 1}
                neighbors[5] = []int{i + 1, j + 1}
                neighbors[6] = []int{i - 1, j}
                neighbors[7] = []int{i + 1, j}

                // fmt.Println("(i,j) (",i,", ",j,")")
                // fmt.Println(neighbors)

                var number_of_neighbors_that_are_on int
                for i := range neighbors {
                    if neighbors[i][0] >= 0 && neighbors[i][0] < len(lights) &&
                       neighbors[i][1] >= 0 && neighbors[i][1] < len(lights) {
                        number_of_neighbors_that_are_on += lights[neighbors[i][0]][neighbors[i][1]]
                    }
                }
                if lights[i][j] == 0 &&
                   number_of_neighbors_that_are_on == 3 {
                    updated_lights[i][j] = 1
                } else if lights[i][j] == 1 &&
                          (number_of_neighbors_that_are_on == 0 ||
                           number_of_neighbors_that_are_on == 1 ||
                           number_of_neighbors_that_are_on == 4 ||
                           number_of_neighbors_that_are_on == 5 ||
                           number_of_neighbors_that_are_on == 6 ||
                           number_of_neighbors_that_are_on == 7 ||
                           number_of_neighbors_that_are_on == 8) {
                    updated_lights[i][j] = 0   
                }
            }
        }

        lights = updated_lights
        lights[0][0] = 1
        lights[0][99] = 1
        lights[99][0] = 1
        lights[99][99] = 1
    }

    var on_count int
    for i := range lights {
        for j := range lights[i] {
            on_count += lights[i][j]
        }
    }


    fmt.Println("On count ",on_count)
}
