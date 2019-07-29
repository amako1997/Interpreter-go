package parser

import (
	"fmt"
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/token"
)

type Parser struct {
	l            *lexer.Lexer
	errors       []string
	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {

	p := &Parser{l: l}
	p.errors = []string{}
	p.nextToken()
	p.nextToken()
	return p
}
func (p *Parser) Errors() []string {
	return p.errors
}
func (p *Parser) peekErros(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		stm := p.parseStatement()

		if stm != nil {
			program.Statements = append(program.Statements, stm)
		}

		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {

	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}
func (p *Parser) parseLetStatement() *ast.LetStatement {

	stm := &ast.LetStatement{Token: p.currentToken}

	if !p.expectNext(token.IDENT) {
		return nil
	}
	stm.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Value}

	if !p.expectNext(token.ASSIGN) {
		return nil
	}
	// keep reading until semicolon is met
	if !p.curentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stm
}

func (p *Parser) curentTokenIs(t token.TokenType) bool {

	return p.currentToken.Type == t
}
func (p *Parser) nextTokenIs(t token.TokenType) bool {

	return p.peekToken.Type == t
}
func (p *Parser) expectNext(t token.TokenType) bool {

	if p.nextTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekErros(t)
	return false
}
