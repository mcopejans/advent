package main

import(
       "fmt"
       // "strconv"
       )

func main() {
    input := "1113122113"
    for count := 0; count < 50 ; count++ {
        fmt.Println(count)
        var doubles = 1
        var new_input string
        for i := range input {
            if (input[i] == input[i+1]) {
                doubles++
            } else {
                new_input += fmt.Sprintf("%d%c", doubles, input[i]) // strconv.Itoa(doubles) + string(input[i])
                doubles = 1
            }

            if (i == len(input) - 2) {
                new_input += fmt.Sprintf("%d%c", doubles, input[i+1]) // strconv.Itoa(doubles) + string(input[i+1])
                input = new_input
                break
            }
        }
    }

    fmt.Println("length", len(input))
}

// package main

// import(
//        "fmt"
//        "strings"
//        "strconv"
//        )

// func main() {
//  input := "1113122113"
//  for count := 0; count < 50 ; count++ {
//      var doubles = 1
//      var new_input string
//      var input_split = strings.Split(input, "")
//      for i := 0; i < len(input_split) - 1; i++ {
//          if (input_split[i] == input_split[i+1]) {
//              doubles++
//          } else {
//              new_input += strconv.Itoa(doubles) + input_split[i]
//              doubles = 1
//          }

//          if (i == len(input) - 2) {
//              new_input += strconv.Itoa(doubles) + input_split[i+1]
//              input = new_input
//          }
//      }
//  }

//  fmt.Println("length", len(input))
// }