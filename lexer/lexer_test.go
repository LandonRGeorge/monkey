package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	testCases := []struct {
		Type    token.TokenType
		Literal string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}
	l := New(input)
	for i, tc := range testCases {
		got := l.NextToken()
		if got.Type != tc.Type {
			t.Errorf("wrong type at [i=%d]; got %q want %q", i, got.Type, tc.Type)
		}
		if got.Literal != tc.Literal {
			t.Errorf("wrong literal at [i=%d]; got %q want %q", i, got.Type, tc.Type)
		}
	}
}
