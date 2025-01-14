package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const (
	salt = "go-2024"
)

func main() {
	var intDecimal, intOctal, intHexadecimal int = 42, 052, 0x2A
	var floatNum float64 = 3.14
	var str string = "Golang"
	var boolean bool = true
	var complexNum complex64 = 1 + 2i
	fmt.Printf("%T\n", complexNum)

	fmt.Println(allTypesToString(intDecimal, intOctal, intHexadecimal, floatNum, str, boolean, complexNum))
	combinedString := stringFromAll(intDecimal, intOctal, intHexadecimal, floatNum, str, boolean, complexNum)
	runeString := []rune(combinedString)

	hashRunes := hashWithSaltInMiddle(runeString)
	fmt.Println(hashRunes)
}

func allTypesToString(args ...interface{}) string {
	var resultStr string

	for i, arg := range args {
		resultStr += fmt.Sprintf("%T", arg)
		if i != len(args)-1 {
			resultStr += "\n"
		}
	}
	return resultStr
}

func stringFromAll(args ...interface{}) string {
	var resultStr string
	for _, arg := range args {
		resultStr += fmt.Sprintf("%v", arg)
	}
	return resultStr
}

func hashWithSaltInMiddle(strRune []rune) string {
	mid := len(strRune) / 2
	runesWithSalt := append(strRune[:mid], append([]rune(salt), strRune[mid:]...)...)
	hash := sha256.Sum256([]byte(string(runesWithSalt)))
	return hex.EncodeToString(hash[:])
}
