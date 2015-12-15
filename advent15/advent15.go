package main

import (
        "bufio"
        "fmt"
        "os")

type Properties struct {
    capacity, durability, flavor, texture, calories int
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var ingredient string
    var c, d, f, t, cal int
    var ingredients = make(map[string]Properties)
    var teaspoons = make(map[string]int)
    for scanner.Scan() {
        _, err := fmt.Sscanf(scanner.Text(), "%s capacity %d, durability %d, flavor %d, texture %d, calories %d",&ingredient,&c, &d, &f, &t, &cal)
        fmt.Println(err)
        ingredients[ingredient] = Properties{c, d, f, t, cal}

        fmt.Println("ingredient", ingredient)
        fmt.Println("c", c)
        fmt.Println("d", d)
        fmt.Println("f", f)
        fmt.Println("t", t)
        fmt.Println("cal", cal)
        teaspoons[ingredient] = 0
    }

    var max_score int
    for i := 0; i <= 100; i++ {  // "Sprinkles"
        for j := 0; j <= 100; j++ {  // "PeanutButter"
            for k := 0; k <= 100; k++ {  // "Frosting"
                for l := 0; l <= 100; l++ {  // "Sugar"
                    if (i + j + k + l) != 100 {
                        continue
                    }
                    total_capacity   := i * ingredients["Sprinkles"].capacity   + j * ingredients["PeanutButter"].capacity   + k * ingredients["Frosting"].capacity   + l * ingredients["Sugar"].capacity
                    total_durability := i * ingredients["Sprinkles"].durability + j * ingredients["PeanutButter"].durability + k * ingredients["Frosting"].durability + l * ingredients["Sugar"].durability
                    total_flavor     := i * ingredients["Sprinkles"].flavor     + j * ingredients["PeanutButter"].flavor     + k * ingredients["Frosting"].flavor     + l * ingredients["Sugar"].flavor
                    total_texture    := i * ingredients["Sprinkles"].texture    + j * ingredients["PeanutButter"].texture    + k * ingredients["Frosting"].texture    + l * ingredients["Sugar"].texture
                    total_calories    := i * ingredients["Sprinkles"].calories    + j * ingredients["PeanutButter"].calories    + k * ingredients["Frosting"].calories    + l * ingredients["Sugar"].calories

                    var score int
                    if total_capacity < 0 || total_durability < 0 || total_flavor < 0 || total_texture < 0 {
                        score = 0
                    } else {
                        score = total_capacity * total_durability * total_flavor * total_texture
                    }

                    // fmt.Println("score ", score, "[", i, ", ", j, ", ", k, ", ", l, "]")

                    if  total_calories == 500 && score > max_score {
                        max_score = score
                    }

                }
            }
        }
    }

    fmt.Println("Max score of cookie ", max_score)
}