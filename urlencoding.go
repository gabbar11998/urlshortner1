package main

import (
	"math"
	"strings"
)

const (
	SYMBOLS = "0123456789abcdefghijklmnopqrsuvwxyzABCDEFGHIJKLMNOPQRSTUVXYZ"
	BASE = int64(len(SYMBOLS))
)
func Encode(number int64) string {
	rest := number % BASE
	result := string(SYMBOLS[rest])
	if number-rest != 0 {
		newnumber := (number - rest) / BASE
		result = Encode(newnumber) + result
	}
	return result
}
func Decode(input string) int64 {
	const floatbase = float64(BASE)
	l := len(input)
	var sum int = 0
	for index := l - 1; index > -1; index -= 1 {
		current := string(input[index])
		pos := strings.Index(SYMBOLS, current)
		sum = sum + (pos * int(math.Pow(floatbase, float64((l-index-1)))))
	}
	return int64(sum)
}
