package main

import (
    "bufio"
    "fmt"
	"math"
    "os"
	"strings"
	"strconv"
)

func main() {
    var total_surface = 0.
    var total_length = 0.
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        var dimension_split = strings.Split(scanner.Text(),"x")
        var l,_ = strconv.ParseFloat(dimension_split[0],64)
        var w,_ = strconv.ParseFloat(dimension_split[1],64)
        var h,_ = strconv.ParseFloat(dimension_split[2],64)
        var surface_wrapping_paper = 2*l*w + 2*l*h + 2*w*h + math.Min(l*w, math.Min(l*h, w*h))
        total_surface += surface_wrapping_paper

        var volume = l*w*h;
        var max = math.Max(l, math.Max(w,h))
        if (l==max) {
            total_length += 2*w + 2*h + volume
        } else if (w==max) {
            total_length += 2*l + 2*h + volume
        } else {
            total_length += 2*l + 2*w + volume   
        }
    }
    fmt.Println("Total surface wrapping paper",total_surface)
    fmt.Println("Total length wrapping paper",total_length)
}