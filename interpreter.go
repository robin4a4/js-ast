package main

func evalBinary(expression BinaryExpression) any {
	return
}

func interpreter(expression Expression) any {
	switch expr := expression.(type) {
	case LiteralExpression:
		return expr.Value
	case BinaryExpression:
		return evalBinary()
	}
}
