package ast

import "interpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	StmNode()
}
type Expression interface {
	Node
	ExpNode()
}

// Program is the root of every ast
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // this will be the LET token prouced by the lexer

	Name *Identifier

	Value Expression
}

func (ls *LetStatement) StmNode()             {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Value }

type Identifier struct {
	Token token.Token

	Value string
}

func (i *Identifier) ExpNode()             {}
func (i *Identifier) TokenLiteral() string { return i.Token.Value }

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) StmNode()             {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Value }
