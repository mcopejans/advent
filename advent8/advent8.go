package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv")

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    var total_string_literals, total_interpreted_literals, total_encoded_literals int
    var raw_string_lit string
    for scanner.Scan() {
    	raw_string_lit = scanner.Text()
        fmt.Printf("length ",raw_string_lit,": ",len(raw_string_lit))
        total_string_literals += len(raw_string_lit)

    	interpreted_string, _ := strconv.Unquote(raw_string_lit)
        fmt.Printf("length ",interpreted_string,": ",len(interpreted_string))
        total_interpreted_literals += len(interpreted_string)

        encoded_string := strconv.QuoteToASCII(raw_string_lit)
        total_encoded_literals += len(encoded_string)
    }
    fmt.Println("Difference 1: ", total_string_literals - total_interpreted_literals)
    fmt.Println("Difference 2: ", total_encoded_literals - total_string_literals)
}