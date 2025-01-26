package main

type Expression interface {
	isExpression()
	isLiteral() bool
	isBinary() bool
}

type LiteralExpression struct {
	Value []byte
}

// Implement the interface
func (l LiteralExpression) isExpression()   {}
func (l LiteralExpression) isLiteral() bool { return true }
func (l LiteralExpression) isBinary() bool  { return false }

type BinaryExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (b BinaryExpression) isExpression()   {}
func (b BinaryExpression) isLiteral() bool { return false }
func (b BinaryExpression) isBinary() bool  { return true }

func parser(tokens []Token) Expression {
	cursor := 0

	literalExpression := func() Expression {
		if tokens[cursor].Type == TOKEN_NUMBER {
			expr := LiteralExpression{
				Value: tokens[cursor].Value,
			}
			cursor++
			return expr
		}

		panic("Unexpected token: " + string(tokens[cursor].Type))
	}

	expression := func() Expression {
		left := literalExpression()
		for cursor < len(tokens) && tokens[cursor].Type == TOKEN_PLUS {
			cursor++
			left = BinaryExpression{
				Left:     left,
				Operator: "+",
				Right:    literalExpression(),
			}
		}
		return left
	}

	return expression()
}
