package lsproduct

import (
	"errors"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	// Validation
	if span < 0 {
		return 0, errors.New("span must be greater than zero")
	}
	
	if span == 0 {
		return 1, nil
	}
	
	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}
	
	if len(digits) == 0 {
		return 0, errors.New("digits string must not be empty")
	}

	// Calculate first product to initialize max
	maxProduct, err := calcSpanValue(digits[0:span])
	if err != nil {
		return 0, err
	}

	// Check all other possible spans
	for i := 1; i <= len(digits)-span; i++ {
		product, err := calcSpanValue(digits[i : i+span])
		if err != nil {
			return 0, err
		}
		
		if product > maxProduct {
			maxProduct = product
		}
	}

	return maxProduct, nil
}

func calcSpanValue(spanPart string) (int64, error) {
	var product int64 = 1

	for _, val := range spanPart {
		if !unicode.IsDigit(val) {
			return 0, errors.New("digits input must only contain digits")
		}
		product *= int64(val - '0')
	}

	return product, nil
}