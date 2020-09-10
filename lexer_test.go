package strm

import (
	"reflect"
	"testing"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []Token
	}{
		{"Operators",
			"=|+(){}[],;",
			[]Token{
				{ASSIGN, "="},
				{PIPE, "|"},
				{PLUS, "+"},
				{LPAREN, "("},
				{RPAREN, ")"},
				{LBRACE, "{"},
				{RBRACE, "}"},
				{LBRACK, "["},
				{RBRACK, "]"},
				{COMMA, ","},
				{SEMICOLON, ";"},
				{EOF, ""}},
		},
		{"Small Program",
			`stream main = [1, 2, 3] | stdout;`,
			[]Token{
				{STREAM, "STREAM"},
				{IDENT, "main"},
				{ASSIGN, "="},
				{LBRACK, "["},
				{INT, "1"},
				{COMMA, ","},
				{INT, "2"},
				{COMMA, ","},
				{INT, "3"},
				{RBRACK, "]"},
				{PIPE, "|"},
				{IDENT, "stdout"},
				{SEMICOLON, ";"},
				{EOF, ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LexInput(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
