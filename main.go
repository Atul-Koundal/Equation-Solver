package main


//So the main logic is i will tokenise the equation

type TokenType int

const (
	TokenType_Unknown TokenType = iota
	TokenType_Number
	TokenType_Operator
	TokenType_OpenBracket
	TokenType_CloseBracket
)

type Token struct {
	Val  string
	Type TokenType
}

