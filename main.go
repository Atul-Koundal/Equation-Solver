package main



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
