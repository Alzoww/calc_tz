package main

import (
	"fmt"
	"strconv"
	"strings"
)

const oper = "+-*/"

var num1, op, num2 string

func main() {
	fmt.Scan(&num1, &op, &num2)
	fmt.Println(result(num1, op, num2))
}
func result(num1, op, num2 string) string {
	check_op(op)

	first, err1 := strconv.Atoi(num1)
	second, err2 := strconv.Atoi(num2)
	if err1 == nil && err2 == nil {
		return strconv.Itoa(num_oper(first, second, op))
	} else {
		return RomeNumber(num1, op, num2)
	}
}

func check_op(op string) string {
	if strings.ContainsAny(op, oper) {
		return oper
	} else {
		panic("Вывод ошибки, так как строка не является математической операцией.")
	}
}

func num_oper(num1, num2 int, op string) int {
	var res int
	if num1 <= 10 && num2 <= 10 {
		switch op {
		case "+":
			res = num1 + num2
		case "-":
			res = num1 - num2
		case "*":
			res = num1 * num2
		case "/":
			res = num1 / num2
		}
		return res
	} else {
		panic("Вывод ошибки, так как жидаются только арбские либо только рисмские числа которые не больше 10")
	}
}

func RomeNumber(num1, op, num2 string) string {
	rome_map := map[string]int{
		"I": 1, "II": 2, "III": 3, "Iv": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	rome1, ex1 := rome_map[num1]
	rome2, ex2 := rome_map[num2]

	if ex1 && ex2 {
		res := num_oper(rome1, rome2, op)
		if res < 0 {
			panic("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
		}
		romesym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

		var result strings.Builder

		for i := 0; i < len(romesym); i++ {
			for res >= values[i] {
				result.WriteString(romesym[i])
				res -= values[i]
			}

		}
		return result.String()
	} else {
		panic("Вывод ошибки, так как на вход ожидаются арабские или римские числа не больше 10")
	}

}