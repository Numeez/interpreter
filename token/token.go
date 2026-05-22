package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true":TRUE,
	"false":FALSE,
	"return":RETURN,
	"if":IF,
	"else":ELSE,

}


func LookUpIdentifiers(identifier string)TokenType{
	if tokenType,ok:= keywords[identifier];ok{
		return  tokenType
	}
	return  IDENT
}


const(
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	IDENT = "IDENT"
	INT = "INT"
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	LT = "<"
	GT = ">"
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE ="}"
	FUNCTION = "FUNCTION"
	LET = "LET"
	IF = "IF"
	ELSE = "ELSE"
	TRUE = "TRUE"
	FALSE = "FALSE"
	RETURN = "RETURN"
	EQ  = "=="
	NOT_EQ = "!="

)