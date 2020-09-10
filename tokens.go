package strm

type TokenType string

const (
	// Placeholders
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Keywords
	STREAM = "STREAM"

	// Operators
	PIPE = "|"
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	SEMICOLON = ";"
	COMMA     = ","
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACK    = "["
	RBRACK    = "]"

	IDENT = "IDENT"
	INT   = "INT"
)

type Token struct {
	Type  TokenType
	Value string
}

var keywords = map[string]TokenType {
	"stream": STREAM,
}

func LookupKeyword(ident string) TokenType {
	if typ, ok := keywords[ident]; ok {
		return typ
	}
	return IDENT
}


