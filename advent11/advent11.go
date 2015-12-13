package main

import(
       "fmt"
       "strings"
       )

func IsGoodPassword(alphabet string, password string) bool {
	var doubles, abc, iol bool
	var found = -2
	for i := 0; i < len(password) - 1; i++ {
		if password[i] == password[i+1] {
			if (i - found >= 2) {  // prevent overlapping
				if (found >= 0) { doubles = true }
				found = i
			}
		}
	}

	for i := 0; i < len(alphabet) - 2; i++ {
		var letters = string(alphabet[i]) + string(alphabet[i+1]) + string(alphabet[i+2]) 
		if strings.Contains(password, letters) {
			abc = true
			break
		}
	}

	if strings.IndexAny(password, "iol") == -1 { iol = true }
	
	// fmt.Println("doubles", doubles)
	// fmt.Println("abc", abc)
	// fmt.Println("iol", iol)
	return doubles && abc && iol

}

func IncreasePassword(alphabet string, length int, password *string) {
	if length == 0 { return }

	last := (*password)[length-1]
	if last == 'z' {
		out := []rune(*password)
        out[length-1] = 'a'
        *password = string(out)
		length = length - 1
		IncreasePassword(alphabet, length, password)
	} else {
		for i := range alphabet {
			if last == alphabet[i] {
				out := []rune(*password)
                out[length-1] = rune(alphabet[i+1])
                *password = string(out)

				// *password = (*password)[0:length-1] + string(alphabet[i+1])
				return
			}
		}
	}
}

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// var alphabet = strings.Split(alphabet1, "")


	fmt.Println("hijklmmn", IsGoodPassword(alphabet, "hijklmmn"))
    fmt.Println("abbceffg", IsGoodPassword(alphabet, "abbceffg"))
	fmt.Println("abbcegjk", IsGoodPassword(alphabet, "abbcegjk"))
	fmt.Println("abcdffaa", IsGoodPassword(alphabet, "abcdffaa"))
	fmt.Println("ghjaabcc", IsGoodPassword(alphabet, "ghjaabcc"))

	// input := "vzbxkghb"
	input := "vzbxxyzz"
	// input := "ghijklmn"
	// input := "abcdefgh"

	// input_split := strings.Split(input, "")
	IncreasePassword(alphabet, len(input), &input)
	for !IsGoodPassword(alphabet, input) {
		// fmt.Println("Bad password", input)
		IncreasePassword(alphabet, len(input), &input)
	}

	fmt.Println("Good password ",input)
}