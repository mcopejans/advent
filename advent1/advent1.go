package main

import (
        "fmt"
        "strings"
)


func main() {
    var lisp string
    _, err := fmt.Scan(&lisp)
    fmt.Println(err)
    fmt.Println("read lisp ",lisp," from stdin")
    strings.Trim(lisp, lisp)
    var lisp_split = strings.Split(lisp, "")
    var floor int = 0
    for i := range lisp_split {
      //fmt.Println(lisp_split[i])
      if lisp_split[i] =="(" {
        floor = floor + 1
      } else {
        floor = floor - 1
      }
      if floor == -1 {
         fmt.Println("First time in basement at index ",i+1)
      }
    }
    fmt.Println("Final floor ",floor)
}
