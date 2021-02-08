package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // foo, bar, x, y...
	INT   = "INT"   // 5, 10...

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Special characters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
)
