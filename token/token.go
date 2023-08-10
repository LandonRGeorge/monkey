package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL" // token we don't know about
	EOF       = "EOF"     // end of file
	IDENT     = "IDENT"   // add, foobar, x, y, ...
	INT       = "INT"     // 123456...
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	IF        = "IF"
	RETURN    = "RETURN"
	TRUE      = "TRUE"
	ELSE      = "ELSE"
	FALSE     = "FALSE"
)

var keywords = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"if":     IF,
	"return": RETURN,
	"true":   TRUE,
	"else":   ELSE,
	"false":  FALSE,
}

func LookupIdent(text string) TokenType {
	ident, ok := keywords[text]
	if !ok {
		return IDENT
	}
	return ident
}
