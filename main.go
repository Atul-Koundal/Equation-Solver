package main

import (
	"fmt"
	"strconv"
	"unicode"
)

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

type Node interface{}

type NumberNode struct {
	Value float64
}

type BinaryNode struct {
	Left  Node
	Op    TokenType
	Right Node
}

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) current() Token {
	return p.tokens[p.pos]
}

func (p *Parser) eat(t TokenType) {
	if p.current().Type != t {
		panic(fmt.Sprintf(
			"expected %s got %s",
			t,
			p.current().Type,
		))
	}
	p.pos++
}

func (p *Parser) parse() Node {
	return p.expression()
}

func (p *Parser) expression() Node {
	node := p.term()

	for p.current().Type == PLUS || p.current().Type == MINUS {
		op := p.current().Type
		p.pos++

		node = &BinaryNode{
			Left:  node,
			Op:    op,
			Right: p.term(),
		}
	}

	return node
}

func (p *Parser) term() Node {
	node := p.factor()

	for p.current().Type == MUL || p.current().Type == DIV {
		op := p.current().Type
		p.pos++

		node = &BinaryNode{
			Left:  node,
			Op:    op,
			Right: p.factor(),
		}
	}

	return node
}

func (p *Parser) factor() Node {
	token := p.current()

	switch token.Type {

	case NUMBER:
		p.eat(NUMBER)

		val, _ := strconv.ParseFloat(token.Literal, 64)

		return &NumberNode{
			Value: val,
		}

	case LPAREN:
		p.eat(LPAREN)

		node := p.expression()

		p.eat(RPAREN)

		return node
	}

	panic("unexpected token")
}

func evaluate(node Node) float64 {
	switch n := node.(type) {

	case *NumberNode:
		return n.Value

	case *BinaryNode:
		left := evaluate(n.Left)
		right := evaluate(n.Right)

		switch n.Op {
		case PLUS:
			return left + right

		case MINUS:
			return left - right

		case MUL:
			return left * right

		case DIV:
			return left / right
		}
	}

	panic("invalid node")
}

func main() {
	input := "(2 + 3) * 4 - 10 / 2"

	tokens := tokenize(input)

	parser := NewParser(tokens)

	ast := parser.parse()

	result := evaluate(ast)

	fmt.Printf("Expression: %s\n", input)
	fmt.Printf("Result: %.2f\n", result)
}