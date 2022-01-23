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
	IDENTIFIER = "IDENTIFIER" // add, foobar, x,
	NUMBER     = "NUMBER"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	MULT   = "*"
	DIV    = "/"

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
)

var keywords = map[string]TokenType{
	"var":    VAR,
	"fun":    FUNCTION,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

func (tok Token) ResolveType(str string) TokenType {
	if v, ok := keywords[str]; ok {
		return v
	}
	return IDENTIFIER
}
