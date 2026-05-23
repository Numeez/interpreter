package parser

import (
	"testing"

	"github.com/Numeez/interpreter/ast"
	"github.com/Numeez/interpreter/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x  = 5;
	let y = 10;
	let foobar = 6969;
	
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseError(t, p)
	if program == nil {
		t.Fatal("program is nil")
	}
	if len(program.Statements) != 3 {
		t.Fatal("program statements are not equal to 3", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 7878;
	
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseError(t, p)
	if program == nil {
		t.Fatal("program is nil")
	}
	if len(program.Statements) != 3 {
		t.Fatal("program statements are not equal to 3", len(program.Statements))
	}

	for  _,stmt := range program.Statements {
		returnStmt,ok:=stmt.(*ast.ReturnStatement)
		if !ok{
			t.Errorf("stmt is not a return statement, got:%T",stmt)
			continue
		}
		if returnStmt.TokenLiteral()!="return"{
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}

}

func testLetStatement(t *testing.T, stmt ast.Statement, expected string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", stmt)
		return false
	}
	if letStmt.Name.Value != expected {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", expected, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != expected {
		t.Errorf("letStmt.Name.TokenLiteral() not equal to %s got %s", expected, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParseError(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parse errror:%q", msg)
	}
	t.FailNow()
}
