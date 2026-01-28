package collatzconjecture

type param_error struct{}
 
func (error_object param_error) Error() string{ 
    return "Invalid parameter"
}

var numOfStepsCurrent int
var numOfStepsTemp int

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, &param_error{}
    } else if n == 1 && numOfStepsCurrent == 0 {
        numOfStepsCurrent = 0
        numOfStepsTemp = 0
        return 0, nil
    } else if n == 1 && numOfStepsCurrent != 0 {
        numOfStepsTemp = numOfStepsCurrent
        numOfStepsCurrent = 0
		return numOfStepsTemp, nil
    }

    if n % 2 == 0 {
        n /= 2
    } else {
        n = n * 3 + 1
    }
	numOfStepsCurrent += 1
    return CollatzConjecture(n)
}
