package token

type TokenType string

const 
( 
	ILLEGAL = "ILLEGAL"

	EOF = "EOF"

	// identifiers and literals

	IDENTIFIER = "IDENTIFIER" // add, foobar, x, y

	INT = "INT"

	//operators

	ASSIGN = "="

	PLUS = "+"

	//delimiters

	COMMA = ","

	SEMICOLON = ";"

	LPAREN = "("

	RPAREN = ")"

	LBRACE = "{"

	RBRACE = "}"

	// keywords

	FUNCTION = "FUNCTION"

	LET = "LET"


)

type Token struct {
	Type TokenType

	Literal string
}
