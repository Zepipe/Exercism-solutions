package hamming

import "unicode/utf8"
import "errors"

func Distance(a, b string) (int, error) {
    hammingDistance := 0
    
	if (a == "" && b == "") || a == b {
        return 0, nil
    } else if a == "" || b == "" || utf8.RuneCountInString(a) != utf8.RuneCountInString(b){
        return 0, errors.New("Given strings are blank or have different length")
    }

    for idx, val := range a {
        for idx2, val2 := range b {
            if idx == idx2 { // to find equal indexes
                if string(val) != string(val2) {
                    hammingDistance += 1
                } else {
                    break
                }
            }
        }
    }

    return hammingDistance, nil
}
