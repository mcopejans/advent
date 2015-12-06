package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func is_nice(input string) bool {
    if (strings.Index(input, "ab") != -1 ||
        strings.Index(input, "cd") != -1 ||
        strings.Index(input, "pq") != -1 ||
        strings.Index(input, "xy") != -1) {
        return false
    }

    var index = -1
    var test_vowels = input
    for i:=0; i < 3; i++ { 
        index = strings.IndexAny(test_vowels, "aeiou")
        if (index == -1) {
            return false
        }
        test_vowels = test_vowels[index+1:]
    }

    for i:=0; i < len(input)-1; i++ {
        if input[i] == input[i+1] {
            return true
        }
    }

    return false
}

func has_repeating_pair(input string) bool {
    for i:=0; i < len(input) - 2; i++ {
        var first = strings.Index(input, input[i:i+2])
        var last = strings.LastIndex(input, input[i:i+2])
        if (first != -1 && last != -1 && last - first >= 2) {
            return true
        }
    }
    return false
}

func letter_in_between(input string) bool {
    for i:=0; i < len(input) - 2; i++ {
        if (input[i] == input[i+2]) {
            return true
        }
    }
    return false
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var i int
    for scanner.Scan() {
        // if is_nice() {  // Solution 1
        if has_repeating_pair(scanner.Text()) && letter_in_between(scanner.Text()) {
            i++
            fmt.Println(scanner.Text())
        }
    }
    fmt.Println("Number of nice strings 2 ", i)
}
