package token

type TokenType string

type Token struct {
	Type	TokenType
	Literal	string
}

const (
	// special
	ILLEGAL 	= "ILLEGAL"
	EOF			= "EOF"

	// identifiers + literals
	IDENT		= "IDENT"	// add, foo, bar, x, ...
	INT			= "INT"		// any integer

	// operators
	ASSIGN		= "="
	PLUS		= "+"
	MINUS		= "-"
	BANG		= "!"
	ASTERISK	= "*"
	SLASH		= "/"

	LT			= "<"
	GT		 	= ">"

	// delimiters
	COMMA		= ","
	SEMICOLON	= ";"

	LPAREN		= "("
	RPAREN		= ")"
	LBRACE		= "{"
	RBRACE		= "}"

	// keywords
	FUNCTION	= "FUNCTION"
	LET			= "LET"
	TRUE		= "TRUE"
	FALSE		= "FALSE"
	IF			= "IF"
	ELSE		= "ELSE"
	RETURN		= "RETURN"
)

var keywords = map[string]TokenType {
	"fn"		: FUNCTION,
	"let"		: LET,
	"true"		: TRUE,
	"false"		: FALSE,
	"if"		: IF,
	"else"		: ELSE,
	"return"	: RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return  IDENT
}
