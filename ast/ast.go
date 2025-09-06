package ast

import "leetcode/src/compiler-in-go/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statement []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statement) > 0 {
		return p.Statement[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token *token.Token
	Name  *Identiefier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identiefier struct {
	Token token.Token
	Value string
}

func (i *Identiefier) expressionNode() {}
func (i *Identiefier) TokenLiteral() string {
	return i.Token.Literal
}
