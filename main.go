package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		valid := true
		fmt.Print("Type the problem: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Exiting programm.")
			return
		}

		array := strings.Fields(scanner.Text())
		if len(array) == 0 {
			fmt.Println("Exiting programm.")
			return
		} else if len(array) <= 1 {
			fmt.Println("Error!")
			valid = false
			continue
		} else if len(array)%2 != 1 {
			fmt.Println("Error!")
			valid = false
			continue
		}

		result, err := strconv.ParseFloat(array[0], 64)
		if err != nil {
			fmt.Println("Error!")
			valid = false
			continue
		}
		for i := 1; i < len(array); i += 2 {
			//num1, _ := strconv.ParseFloat(array[i-1], 64)
			num2, err := strconv.ParseFloat(array[i+1], 64)
			if err != nil {
				fmt.Println("Error!")
				valid = false
				break
			} else {
				switch array[i] {
				case "+":
					result += num2
				case "-":
					result -= num2
				case ":", "/":
					result /= num2
				case "*":
					result *= num2
				default:
					fmt.Println("Error! Not an operator, the operators is (+, -)")
					valid = false
					break
				}
			}
		}
		if valid {
			fmt.Println(result)
		} else {
			continue
		}
	}
}
