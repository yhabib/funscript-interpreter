package lexer

import (
	"fmt"
	"funscript-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `var two = 2;
	var add = fun(x, y) {
		return x + y;
	}

	var result = add(two, 3);

	1 - 2;
	1 * 3;
	5 / 2; 
	
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENTIFIER, "two"},
		{token.ASSIGN, "="},
		{token.NUMBER, "2"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fun"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "two"},
		{token.COMMA, ","},
		{token.NUMBER, "3"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.NUMBER, "1"},
		{token.MINUS, "-"},
		{token.NUMBER, "2"},
		{token.SEMICOLON, ";"},
		{token.NUMBER, "1"},
		{token.MULT, "*"},
		{token.NUMBER, "3"},
		{token.SEMICOLON, ";"},
		{token.NUMBER, "5"},
		{token.DIV, "/"},
		{token.NUMBER, "2"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Println(tok.Literal)
		fmt.Println(tok.Type)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
