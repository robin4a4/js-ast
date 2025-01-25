package main

import (
	"fmt"
)

type TokenType int

const (
	TOKEN_NUMBER TokenType = iota
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_MULTIPLY
	TOKEN_EQUAL
	TOKEN_LEFT_PAREN
	TOKEN_RIGHT_PAREN
)

type Token struct {
	Type  TokenType
	Value []byte
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func lexer(str []byte) []Token {
	tokens := []Token{}
	cursor := 0
	fmt.Println(str)

	number := func() {
		numberStartCursor := cursor
		nextCursor := cursor + 1
		for isDigit(str[nextCursor]) {
			cursor++
		}

		if str[cursor] == '.' && isDigit(str[nextCursor+1]) {
			for isDigit(str[nextCursor]) {
				cursor++
			}
		}

		tokens = append(tokens, Token{
			Type:  TOKEN_NUMBER,
			Value: str[numberStartCursor:nextCursor],
		})
	}

	for cursor < len(str) {
		currentChar := str[cursor]
		switch currentChar {
		case '\n', '\r', '\t', ' ':
		case '-':
			tokens = append(tokens, Token{
				Type:  TOKEN_PLUS,
				Value: []byte{currentChar},
			})
		case '+':
			tokens = append(tokens, Token{
				Type:  TOKEN_PLUS,
				Value: []byte{currentChar},
			})
		default:
			if isDigit(currentChar) {
				number()
				break
			}
			panic("Unexpected token:" + string(currentChar))
		}
		cursor++
	}
	return tokens
}
