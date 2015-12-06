package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Grid[1000][1000]int

func turn_on(x1 int, y1 int, x2 int, y2 int, lights *Grid) {
    for i:=x1; i <= x2; i++ {
        for j:=y1; j <= y2; j++ {
            (*lights)[i][j] += 1
        }
    }
}

func turn_off(x1 int, y1 int, x2 int, y2 int, lights *Grid) {
    for i:=x1; i <= x2; i++ {
        for j:=y1; j <= y2; j++ {
            if ((*lights)[i][j] > 0) {
                (*lights)[i][j] -= 1
            }
        }
    }
}

func toggles(x1 int, y1 int, x2 int, y2 int, lights *Grid) {
    turn_on(x1, y1, x2, y2, lights)
    turn_on(x1, y1, x2, y2, lights)
}

func count(lights *Grid) int {
    var counter int
    for i:=0; i < len((*lights)); i++ {
        for j:=0; j < len((*lights)[i]); j++ {
            counter += (*lights)[i][j]
        }
    }
    return counter
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var on = "turn on"
    var off = "turn off"
    var toggle = "toggle"
    var x1,y1,x2,y2 int
    var lights Grid

    for scanner.Scan() {
        fmt.Println(scanner.Text())
        if strings.Contains(scanner.Text(), on) {
            fmt.Sscanf(scanner.Text(),on + " %d,%d through %d,%d",&x1,&y1,&x2,&y2)
            turn_on(x1,y1,x2,y2,&lights)
        } else if strings.Contains(scanner.Text(), off) {
            fmt.Sscanf(scanner.Text(),off + " %d,%d through %d,%d",&x1,&y1,&x2,&y2)
            turn_off(x1,y1,x2,y2,&lights)
        } else if strings.Contains(scanner.Text(), toggle) {
            fmt.Sscanf(scanner.Text(),toggle + " %d,%d through %d,%d",&x1,&y1,&x2,&y2)
            toggles(x1,y1,x2,y2,&lights)
        } else {
            fmt.Println("err")
        }
    }

    fmt.Println("Total brightness ",count(&lights))
}
