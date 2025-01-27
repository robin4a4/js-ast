package main

import (
	"strconv"
)

func evalBinary(expression BinaryExpression) any {
	switch expression.Operator {
	case "+":
		leftNum, _ := strconv.Atoi(string(interpreter(expression.Left).([]byte)))
		rightNum, _ := strconv.Atoi(string(interpreter(expression.Right).([]byte)))
		return []byte(strconv.Itoa(leftNum + rightNum))
	}
	panic("Error")
}

func interpreter(expression Expression) any {
	switch expr := expression.(type) {
	case LiteralExpression:
		return expr.Value
	case BinaryExpression:
		return evalBinary(expr)
	}

	panic("Error")
}
