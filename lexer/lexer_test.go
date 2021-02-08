package lexer

import (
	"lyz-lang/token"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `=+(){},;`

	expectedTokens := []token.Token{
		token.Token{token.ASSIGN, "="},
		token.Token{token.PLUS, "+"},
		token.Token{token.LPAREN, "("},
		token.Token{token.RPAREN, ")"},
		token.Token{token.LBRACE, "{"},
		token.Token{token.RBRACE, "}"},
		token.Token{token.COMMA, ","},
		token.Token{token.SEMICOLON, ";"},
	}

	l := New(input)

	var nextToken token.Token
	for k, v := range expectedTokens {
		nextToken = l.NextToken()
		if v.Type != nextToken.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", k, v.Type, nextToken.Type)
		}

		if v.Literal != nextToken.Literal {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", k, v.Literal, nextToken.Literal)
		}
	}
}
