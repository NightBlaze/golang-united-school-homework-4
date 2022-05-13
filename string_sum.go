package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	return calculate(input)
}

func calculate(in string) (string, error) {
	if isOnlyWhitespaces(in) {
		return "", fmt.Errorf("only whitespaces: %w", errorEmptyInput)
	}

	in = strings.ReplaceAll(in, " ", "")

	return calculateWithourWhitespaces(in)
}

func calculateWithourWhitespaces(in string) (string, error) {
	numberOfNumbers := 0
	result := 0
	numberAsString := ""
	for _, r := range in {
		if !isValidCharacter(string(r)) {
			numberAsString += string(r)
			_, err := toInt(string(numberAsString))
			return "", fmt.Errorf("not valid character: %w", err)
		}

		if len(numberAsString) == 0 {
			numberAsString += string(r)
			continue
		}

		if isSign(string(r)) {
			var err error
			numberAsString, result, numberOfNumbers, err = addResult(numberAsString, result, numberOfNumbers, string(r))
			if err != nil {
				return "", err
			}
			continue
		}

		numberAsString += string(r)
	}

	if len(numberAsString) > 0 {
		var err error
		numberAsString, result, numberOfNumbers, err = addResult(numberAsString, result, numberOfNumbers, "")
		if err != nil {
			return "", err
		}
	}
	if numberOfNumbers < 2 {
		return "", fmt.Errorf("too low: %w", errorNotTwoOperands)
	}
	return strconv.Itoa(result), nil
}

func addResult(numberAsString string, result, numberOfNumbers int, r string) (string, int, int, error) {
	i, err := toInt(numberAsString)
	if err != nil {
		return "", 0, 0, fmt.Errorf("can't convert to int: %w", err)
	}
	result += i
	numberOfNumbers += 1
	if numberOfNumbers > 2 {
		return "", 0, 0, fmt.Errorf("too many: %w", errorNotTwoOperands)
	}
	numberAsString = r

	return numberAsString, result, numberOfNumbers, nil
}

func isSign(r string) bool {
	return r == "+" || r == "-"
}

func isNumber(r string) bool {
	return r == "0" || r == "1" || r == "2" || r == "3" || r == "4" || r == "5" || r == "6" || r == "7" || r == "8" || r == "9"
}

func toInt(in string) (int, error) {
	in = strings.ReplaceAll(in, "+", "")
	return strconv.Atoi(in)
}

func isOnlyWhitespaces(in string) bool {
	return strings.TrimSpace(in) == ""
}

func isValidCharacter(in string) bool {
	return isNumber(in) || isSign(in)
}
