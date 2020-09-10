package strm

import "unicode"

type Lexer struct {
	input string

	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func LexInput(input string) []Token {
	l := NewLexer(input)
	toks := make([]Token, 0)
	for {
		tok := l.NextToken()
		toks = append(toks, tok)
		if tok.Type == EOF {
			break
		}
	}
	return toks
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()
	switch l.ch {
	case '=':
		tok = newToken(ASSIGN, l.ch)
	case '|':
		tok = newToken(PIPE, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case '+':
		tok = newToken(PLUS, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	case '[':
		tok = newToken(LBRACK, l.ch)
	case ']':
		tok = newToken(RBRACK, l.ch)
	case 0:
		tok.Value = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			tok.Value = l.readWhile(isLetter)
			tok.Type = LookupKeyword(tok.Value)
			return tok // readIdentifier consumes all chars, don't read again
		} else if isDigit(l.ch) {
			tok.Value = l.readWhile(isDigit)
			tok.Type = INT
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tk TokenType, c byte) Token {
	return Token{Type: tk, Value: string(c)}
}

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || c == '_'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func (l *Lexer) readWhile(meetsCondition func(byte) bool) string {
	start := l.position
	for meetsCondition(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(l.ch)) {
		l.readChar()
	}
}
