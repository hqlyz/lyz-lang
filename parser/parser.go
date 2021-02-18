package parser

import (
	"fmt"
	"lyz-lang/ast"
	"lyz-lang/lexer"
	"lyz-lang/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errs []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errs: []string{}}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errs() []string {
	return p.errs
}

func (p *Parser) peekError(t token.TokenType) {
	err := fmt.Sprintf(fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type))
	p.errs = append(p.errs, err)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{Statements: []ast.Statement{}}
	for !p.curTokenIs(token.EOF) {
		stm := p.parseStatement()
		if stm != nil {
			prog.Statements = append(prog.Statements, stm)
		}
		p.nextToken()
	}
	return prog
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stm := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stm.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stm
}

func (p *Parser)parseReturnStatement() *ast.ReturnStatement {
	stm := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	
	return stm
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}
