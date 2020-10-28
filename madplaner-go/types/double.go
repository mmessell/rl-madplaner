package types

import (
	"fmt"
	"math"
)

type Double float64

func (d Double) ToString() string {
	s := fmt.Sprintf("%f", d)
	s = removeSuffix(s, "0")
	s = removeSuffix(s, ".")
	return s
}

func (d Double) RoundUp() Double {
	return Double(math.Ceil(float64(d)))
}

func removeSuffix(s string, suffix string) string {
	if len(s) == 0 {
		return s
	}

	lastChar := s[len(s)-1:]

	if lastChar != suffix {
		return s
	}

	return removeSuffix(s[:len(s)-1], suffix)
}
