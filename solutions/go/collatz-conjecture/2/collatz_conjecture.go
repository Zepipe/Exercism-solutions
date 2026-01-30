package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    var numOfSteps int // To count computation steps
    
    if n <= 0 {
        return 0, errors.New("Invalid input")
    } else if n == 1 {
        return 0, nil
    }

    for {
        if n % 2 == 0 {
            n /= 2
        } else {
            n = n * 3 + 1
        }
        numOfSteps += 1

        if n == 1 {
            return numOfSteps, nil
        }
    }
    
}
