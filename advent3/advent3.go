package main
	
import (
  "fmt"
  "math"
  "strings"
)

type coordinate struct {
    x, y int 
}

func count_visited(instruction string, visited map[coordinate]int, c *coordinate) {
    switch instruction {
        case "^":
            (*c).y++
        case "v":
            (*c).y--
        case "<":
            (*c).x--
        case ">":
            (*c).x++
        default:
            panic("Unknown instruction, doesn't know where to go.")
    }

    if visited[*c]!=0 {
        visited[*c]++
    } else {
        visited[*c] = 1
    }
}

func main() {
	var instructions string
    fmt.Scan(&instructions)
    var instruction = strings.Split(instructions, "")
    
    var visited = make(map[coordinate]int)
    var c1, c2 coordinate

    visited[c1]=1
    for i := range instruction {
        if (math.Mod(float64(i),2.)==0) {
            count_visited(instruction[i], visited, &c1)
        } else {
            count_visited(instruction[i], visited, &c2)
        }
    }

    fmt.Println("total num of presents",len(visited))
}