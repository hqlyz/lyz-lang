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
	ASSIGN = "="  // 赋值
	PLUS   = "+"  // 加法
	MINUS  = "-"  // 减法
	MULTI  = "*"  // 乘法
	DIV    = "/"  // 除法
	EQUAL  = "==" // 相等
	NEQUAL = "!=" // 不相等
	LT     = "<"  // 小于
	GT     = ">"  // 大于
	NEG    = "!"  // 取反

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if v, ok := keywords[ident]; ok {
		return v
	}
	return IDENT
}
