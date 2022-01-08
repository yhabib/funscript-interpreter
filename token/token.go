package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // Identifies an illegal token
	EOF     = "EOF"     // Tells the parser that there is nothing to read

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER" // add, foobar, x, y, ...
	INT        = "INT"        // 1343456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	RETURN   = "RETURN"

	// Types
	NUMBER = "NUMBER"
)

var keywords = map[string]TokenType{
	"var":    VAR,
	"fun":    FUNCTION,
	"return": RETURN,
}

func (tok Token) ResolveType(str string) TokenType {
	if v, ok := keywords[str]; ok {
		return v
	}
	return IDENTIFIER
}
