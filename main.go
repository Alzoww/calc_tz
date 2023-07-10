package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const oper = "+-*/"

func main() {
	var num1, op, num2 string
	var x int
	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(console, " ", "")
		indplus := strings.Index(s, "+")
		indmin := strings.Index(s, "-")
		inddiv := strings.Index(s, "/")
		indtim := strings.Index(s, "*")
		switch {
		case indplus != -1:
			x = indplus
		case indmin != -1:
			x = indmin
		case inddiv != -1:
			x = inddiv
		case indtim != -1:
			x = indtim
		}

		num1 = s[0:x]
		op = string(s[x])
		switch {
		case len(s) == 5 && len(num1) == 1:
			num2 = s[x+1 : x+2]
		case len(s) == 6 && len(num1) == 1:
			num2 = s[x+1 : x+3]
		case len(s) == 7 && len(num1) == 1:
			num2 = s[x+1 : x+4]
		case len(s) == 8 && len(num1) == 1:
			num2 = s[x+1 : x+5]
		case len(s) == 6 && len(num1) == 2:
			num2 = s[x+1 : x+2]
		case len(s) == 7 && len(num1) == 2:
			num2 = s[x+1 : x+3]
		case len(s) == 8 && len(num1) == 2:
			num2 = s[x+1 : x+4]
		case len(s) == 9 && len(num1) == 2:
			num2 = s[x+1 : x+5]
		case len(s) == 7 && len(num1) == 2:
			num2 = s[x+1 : x+2]
		case len(s) == 8 && len(num1) == 3:
			num2 = s[x+1 : x+3]
		case len(s) == 9 && len(num1) == 3:
			num2 = s[x+1 : x+4]
		case len(s) == 10 && len(num1) == 3:
			num2 = s[x+1 : x+5]
		case len(s) == 8 && len(num1) == 4:
			num2 = s[x+1 : x+2]
		case len(s) == 9 && len(num1) == 4:
			num2 = s[x+1 : x+3]
		case len(s) == 10 && len(num1) == 4:
			num2 = s[x+1 : x+4]
		case len(s) == 11 && len(num1) == 4:
			num2 = s[x+1 : x+5]
		}

		fmt.Println(result(num1, op, num2))
	}
}

func result(num1, op, num2 string) string {
	check_op(op)

	first, err1 := strconv.Atoi(num1)
	second, err2 := strconv.Atoi(num2)
	switch {
	case err1 == nil && err2 == nil:
		return strconv.Itoa(num_oper(first, second, op))
	case err1 != nil && err2 == nil:
		panic("Вывод ошибки, так как используются одновременно разные системы счисления.")
	case err1 == nil && err2 != nil:
		panic("Вывод ошибки, так как используются одновременно разные системы счисления.")
	default:
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
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	Rome1, ex1 := rome_map[num1]
	Rome2, ex2 := rome_map[num2]
	if ex1 && ex2 {
		num := num_oper(Rome1, Rome2, op)
		if num < 0 {
			panic("Ошибка!В римской системе нет отрицательных чисел.")
		}
		symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		var result strings.Builder
		for i := 0; i < len(symbols); i++ {
			for num >= values[i] {
				result.WriteString(symbols[i])
				num -= values[i]
			}
		}
		return result.String()
	} else {
		panic("Вывод ошибки, так как на вход ожидаются арабские или римские числа не больше 10")
	}
}
