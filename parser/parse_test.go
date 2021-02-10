package parser

import (
	"lyz-lang/ast"
	"lyz-lang/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
		`
	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()
	if prog == nil {
		t.Fatalf("ParseProgram returned nil")
	}

	if len(prog.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(prog.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for k, tt := range tests {
		stm := prog.Statements[k]
		if !testLetStatement(t, stm, tt.expectedIdentifier) {
			t.Fatalf("the let statement parsed failed. expected: %s, got: %s", tt.expectedIdentifier, stm.TokenLiteral())
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", name, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}
