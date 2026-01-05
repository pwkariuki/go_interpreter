package ast

import "interpreter/token"

type Node interface {
	TokenLiteral() string // literal value of node's token (debug)
}

// statements do not produce values
type Statement interface {
	Node
	statementNode()
}

// expressions produce values
type Expression interface {
	Node
	expressionNode()
}

// AST's root node
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

// node for a variable binding
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // identifier of the binding
	Value Expression  // expression that produces the value
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
