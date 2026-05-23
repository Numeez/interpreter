package parser

import (
	"fmt"

	"github.com/Numeez/interpreter/ast"
	"github.com/Numeez/interpreter/lexer"
	"github.com/Numeez/interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatements()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatements() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil

	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeekToken(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeekToken(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt

}
func (p *Parser) expectPeekToken(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}

}

func (p *Parser) peekTokenIs(tok token.TokenType) bool {
	return p.peekToken.Type == tok
}
func (p *Parser) curTokenIs(tok token.TokenType) bool {
	return p.curToken.Type == tok
}
func (p *Parser) Errors() []string { return p.errors }

func (p *Parser) peekError(token token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", token, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
