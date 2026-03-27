package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
  "math"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Type the problem: ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Exiting program.")
			return
		}

		input := scanner.Text()
		if input == "" {
			fmt.Println("Exiting program.")
			return
		}

		array := strings.Fields(input)

		// Простейшая проверка на корректность (число операндов и операторов)
		if len(array)%2 == 0 {
			fmt.Println("Error! Invalid expression format.")
			continue
		}

		// Обработка умножения и деления (высокий приоритет)
		// Мы создаем новый слайс, куда складываем числа, уже умноженные или деленные
		var intermediate []string
		valid := true

		for i := 0; i < len(array); i++ {
			if i%2 == 0 {
				// Это число, пока просто добавляем
				intermediate = append(intermediate, array[i])
			} else {
				// Это оператор
				op := array[i]
				if op == "*" || op == "/" || op == ":" || op == "^" {
					// Берем последнее число из результата, считаем и заменяем его
					leftVal, _ := strconv.ParseFloat(intermediate[len(intermediate)-1], 64)
					rightVal, err := strconv.ParseFloat(array[i+1], 64)
					if err != nil {
						valid = false
						break
					}

					var res float64
					if op == "*" {
						res = leftVal * rightVal
					} else if op == "^" {
            res = math.Pow(leftVal, rightVal)
          } else {
						if rightVal == 0 {
							fmt.Println("Error: Division by zero!")
							valid = false
							break
						}
						res = leftVal / rightVal
					}
					// Обновляем последнее число в промежуточном слайсе
					intermediate[len(intermediate)-1] = fmt.Sprintf("%f", res)
					i++ // Пропускаем следующее число, так как мы его уже использовали
				} else {
					// Это + или -, просто переносим в промежуточный список
					intermediate = append(intermediate, op)
				}
			}
		}

		if !valid {
			continue
		}

		// ШАГ 2: Обработка сложения и вычитания (низкий приоритет)
		result, _ := strconv.ParseFloat(intermediate[0], 64)
		for i := 1; i < len(intermediate); i += 2 {
			op := intermediate[i]
			num, err := strconv.ParseFloat(intermediate[i+1], 64)
			if err != nil {
				fmt.Println("Error parsing number!")
				valid = false
				break
			}

			switch op {
			case "+":
				result += num
			case "-":
				result -= num
			default:
				fmt.Println("Error! Unknown operator:", op)
				valid = false
				break
			}
		}

		if valid {
			fmt.Printf("Result: %v\n", result)
		}
	}
}
