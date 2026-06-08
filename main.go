package main

import (
	"unicode"
)

//So the main logic is i will tokenise the equation then i will build the ast from this to 
//deal with the predence and all ...
//Basically the replication on how interpreter and compilers work on these equations

type TokenType string

const (
	NUMBER TokenType = "NUMBER"

	PLUS  TokenType = "+"
	MINUS TokenType = "-"
	MUL   TokenType = "*"
	DIV   TokenType = "/"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"

	EOF TokenType = "EOF"
)

//This Literal just meant what is the value of the particular token
type Token struct {
	Type    TokenType
	Literal string
}


func tokenize(input string) []Token {
	var tokens []Token

	for i := 0; i < len(input); {
		ch := rune(input[i])

		if unicode.IsSpace(ch) {
			i++
			continue
		}

		if unicode.IsDigit(ch) {
			start := i

			for i < len(input) && unicode.IsDigit(rune(input[i])) {
				i++
			}

			tokens = append(tokens, Token{
				Type:    NUMBER,
				Literal: input[start:i],
			})
			continue
		}
		switch ch {
		case '+':
			tokens = append(tokens, Token{Type: PLUS, Literal: "+"})
		case '-':
			tokens = append(tokens, Token{Type: MINUS, Literal: "-"})
		case '*':
			tokens = append(tokens, Token{Type: MUL, Literal: "*"})
		case '/':
			tokens = append(tokens, Token{Type: DIV, Literal: "/"})
		case '(':
			tokens = append(tokens, Token{Type: LPAREN, Literal: "("})
		case ')':
			tokens = append(tokens, Token{Type: RPAREN, Literal: ")"})
		default:
			panic("invalid character: " + string(ch))
		}

		i++
	}

	tokens = append(tokens, Token{Type: EOF})
	return tokens
}