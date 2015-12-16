package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
        "strings")

type Properties struct {
    capacity, durability, flavor, texture, calories int
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var props = make(map[string][]int)
    var all = []string{"children", "cats", "samoyeds", "pomeranians", "akitas", "vizslas", "goldfish", "trees", "cars", "perfumes"}
    var sol = make(map[string]int)
    sol["children"] = 3
    sol["cats"] = 7
    sol["samoyeds"] = 2
    sol["pomeranians"] = 3
    sol["akitas"] = 0
    sol["vizslas"] = 0
    sol["goldfish"] = 5
    sol["trees"] = 3
    sol["cars"] = 2
    sol["perfumes"] = 1
    for i :=range all {
        props[all[i]] = make([]int, 500)
        for j := 0; j < 500; j++ {
            props[all[i]][j] = -1
        }
    }

    var aunt_number int
    for scanner.Scan() {
        fields := strings.Fields(scanner.Text())
        fmt.Println(fields)
        for i := 2; i < len(fields) - 1; i+=2 {
            props[fields[i]][aunt_number], _ = strconv.Atoi(fields[i+1])
        }
        aunt_number++
    }

    fmt.Println(props)

    var match = make([]int, 500)
    for k, v := range props {
        for i := range v {
            if k == "trees" || k == "cates" {
                if props[k][i] > sol[k] || props[k][i] == -1 {
                    match[i]++
                    if match[i] == 10 {
                        fmt.Println("aunt ", i+1)
                    }
                }
            } else if k== "pomeranians" || k == "goldfish" {
                if props[k][i] < sol[k] || props[k][i] == -1 {
                    match[i]++
                    if match[i] == 10 {
                        fmt.Println("aunt ", i+1)
                    }
                }
            } else if props[k][i] == sol[k] || props[k][i] == -1 {
                match[i]++
                if match[i] == 10 {
                    fmt.Println("aunt ", i+1)
                }
            }
        }
    }
}