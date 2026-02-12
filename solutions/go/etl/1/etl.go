package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	if len(in) == 0 {
        return map[string]int{}
    }

    out := make(map[string]int)

    for _, list := range in {
        for _, letter := range list {
        	letter = strings.ToLower(letter)
        	_, exists := out[letter]
        	if !exists {
            	// check score for each letter
                if letter == "a" || letter == "e" || letter == "i" || letter == "o" || letter == "u" || letter == "l" || letter == "n" 				|| letter == "r" || letter == "s" || letter == "t" {
                    out[letter] = 1
                } else if letter == "d" || letter == "g" {
                    out[letter] = 2 
                } else if letter == "b" || letter == "c" || letter == "m" || letter == "p" {
                    out[letter] = 3
                } else if letter == "f" || letter == "h" || letter == "v" || letter == "w" || letter == "y" {
                    out[letter] = 4
                } else if letter == "k" {
                    out[letter] = 5
                } else if letter == "j" || letter == "x" {
                    out[letter] = 8
                } else if letter == "q" || letter == "z" {
                    out[letter] = 10
                }
            }
        }
    }

    return out   
}
