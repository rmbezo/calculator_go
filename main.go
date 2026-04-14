package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// User type and basic checks + Convert into slice
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Type your problem (Enter to exit): ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Error!")
			break
		}

		userText := scanner.Text()
		userSlice := strings.Fields(userText)

		if len(userSlice)%2 != 1 {
			fmt.Println("Exiting.")
			return
		}

		// Check this section :
		// fmt.Println(userText, userSlice)

		// Basic + and -
		// fmt.Println(lOp(userSlice))

		// Real calculator!!! So happy!!
		value, key := hvOp(userSlice)
		// fmt.Println("hvOp:", value, key)
		fmt.Println(lOp(value, key))
	}
}

// Create function to do operations with high valuable operators
func hvOp(l []string) ([]string, error) {
	errParseFloat := errors.New("Error parsing float!")
	hvSlice := []string{}
	for i := 0; i < len(l); i++ {
		if i%2 != 1 {
			hvSlice = append(hvSlice, l[i])
			continue
		}
		// else if l[i] == "*" || l[i] == ":" || l[i] == "/" || l[i] == "^" {
		// }
		switch l[i] {
		case "*":
			leftN, err := strconv.ParseFloat(hvSlice[len(hvSlice)-1], 64)
			if err != nil {
				return []string{"0"}, errParseFloat
			}
			rightN, err := strconv.ParseFloat(l[i+1], 64)
			if err != nil {
				return []string{"0"}, errParseFloat
			}
			newNum := strconv.FormatFloat(leftN*rightN, 'f', 64, 64)
			hvSlice[len(hvSlice)-1] = newNum
			i++
		case "/", ":":
			leftN, err := strconv.ParseFloat(hvSlice[len(hvSlice)-1], 64)
			if err != nil {
				return []string{"0"}, errParseFloat
			}
			rightN, err := strconv.ParseFloat(l[i+1], 64)
			if err != nil {
				return []string{"0"}, errParseFloat
			}
			if rightN == 0 {
				errZero := errors.New("Divide on zero!!!")
				fmt.Println("Error! Cannot divide on zero!")
				return []string{"0"}, errZero
			}
			newNum := strconv.FormatFloat(leftN/rightN, 'f', 64, 64)
			// hvSlice[len(hvSlice)] = newNum
			// hvSlice[i-1] = newNum
			hvSlice[len(hvSlice)-1] = newNum
			i++
		case "^":
			leftN, err := strconv.ParseFloat(hvSlice[len(hvSlice)-1], 64)
			if err != nil {
				return []string{"0"}, errParseFloat
			}
			rightN, err := strconv.ParseFloat(l[i+1], 64)
			if err != nil {
				return []string{"0"}, errParseFloat
			}
			newNum := strconv.FormatFloat(math.Pow(leftN, rightN), 'f', 64, 64)
			// hvSlice[len(hvSlice)] = newNum
			// hvSlice[i-1] = newNum
			hvSlice[len(hvSlice)-1] = newNum
			i++
		case "+", "-":
			hvSlice = append(hvSlice, l[i])
		default:
			errNotOp := errors.New("Not an operator!")
			return []string{"0"}, errNotOp
		}
	}
	return hvSlice, nil
}

// Function with + and -
func lOp(l []string, e error) (float64, error) {
	if e != nil {
		errHvOp := errors.New("Error with high value operator function!")
		return 0, errHvOp
	}
	result, _ := strconv.ParseFloat(l[0], 64)
	for i := 1; i < len(l); i += 2 {
		num, err := strconv.ParseFloat(l[i+1], 64)
		if err != nil {
			fmt.Println("Error, failed at parsing to float")
			errParse := errors.New("Error! Not able to parse number!!")
			return -1, errParse
		}
		switch l[i] {
		case "+":
			result += num
		case "-":
			result -= num
		}
	}

	return result, nil
}
