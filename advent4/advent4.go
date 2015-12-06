package main

import (
        "crypto/md5"
        "fmt"
        "encoding/hex"
        "strconv"
        "strings")

func starts_with_five_zeros(sum string) bool {
    return strings.HasPrefix(sum, "000000")
}

func main() {
    var key string
    fmt.Scan(&key)

    var i int
    var md5_sum string
    for !starts_with_five_zeros(md5_sum) {
        var md5_sum_byte = md5.Sum([]byte(key + strconv.Itoa(i)))
        md5_sum = hex.EncodeToString(md5_sum_byte[:])
        i++;
    }

    fmt.Println("Number found ", i-1, " corresponding hash ", md5_sum, " input was ", key + strconv.Itoa(i-1))
}