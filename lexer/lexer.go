package lexer

import (
	"unicode"

	"github.com/mahabubulhasibshawon/arithmetic-lexer-parser/domain"
)

type Lexer struct {
	input   string
	pos     int
	readPos int
	ch      byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() domain.Token {
	var tok domain.Token
	l.skipWhiteSpace()

	switch l.ch {
	case '+':
		tok = domain.Token{Type: domain.TokenAdd, Value: "+"}
	case '-':
		tok = domain.Token{Type: domain.TokenSub, Value: "-"}
	case '*':
		tok = domain.Token{Type: domain.TokenMul, Value: "*"}
	case '/':
		tok = domain.Token{Type: domain.TokenDiv, Value: "/"}
	case 0:
		tok = domain.Token{Type: domain.TokenEof, Value: ""}
	default:
		if unicode.IsDigit(rune(l.ch)) {
			tok.Type = domain.TokenNumber
			tok.Value = l.readNumber()
			return tok
		}
		tok = domain.Token{Type: domain.TokenEof, Value: ""}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readNumber() string {
	start := l.pos
	for unicode.IsDigit(rune(l.ch)) || l.ch == '.' {
		l.readChar()
	}
	return l.input[start:l.pos]
}
