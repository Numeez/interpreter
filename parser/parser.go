package parser

import (
	"fmt"
	"strconv"

	"github.com/Numeez/interpreter/ast"
	"github.com/Numeez/interpreter/lexer"
	"github.com/Numeez/interpreter/token"
)

const(
	_int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type(
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression)ast.Expression

)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefixFn(token.IDENT, p.parseIdentifier)
	p.registerPrefixFn(token.INT,p.parseIntegerLiteral)
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
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()

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

func (p *Parser)parseReturnStatement()*ast.ReturnStatement{
	returnStmt:= &ast.ReturnStatement{
		Token: p.curToken,
	}
	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON){
		p.nextToken()
	}
	return returnStmt

}
func (p *Parser)parseExpressionStatement()*ast.ExpressionStatement{
	expressionStmt:= &ast.ExpressionStatement{
		Token: p.curToken,
	}
	expressionStmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON){
		p.nextToken()
	}
	return expressionStmt
}

func (p *Parser)parseIdentifier()ast.Expression{
	return &ast.Identifier{
		Token:p.curToken,
		Value: p.curToken.Literal,
	}
}
func (p *Parser)parseExpression(precedence int)ast.Expression{
	prefix:=p.prefixParseFns[p.curToken.Type]
	if prefix==nil{
		return  nil
	}
	leftExp:= prefix()
	return leftExp
}
func(p *Parser)parseIntegerLiteral()ast.Expression{
	lit:=&ast.IntegerLiteral{
		Token: p.curToken,
	}
	value,err:= strconv.ParseInt(p.curToken.Literal,0,64)
	if err!=nil{
		msg := fmt.Sprintf("could not parse %q as integer",p.curToken.Literal)
		p.errors  = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
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

func (p *Parser)registerPrefixFn(tokenType token.TokenType,fn prefixParseFn){
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser)registerInfixFn(tokenType token.TokenType,fn infixParseFn){
	p.infixParseFns[tokenType] = fn
}