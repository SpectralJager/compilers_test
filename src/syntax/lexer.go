package syntax

import (
	"fmt"
)

/*
symbol = [a-zA-Z_]+
number = [0-9]+(\.[0-9]+)?
string = \"[[:ascii:]]\"
*/

type Lexer struct {
	source string
	tokens []Token
	pos    int
	line   int
	column int
	errors []error
}

func InitLexer(source string) *Lexer {
	return &Lexer{
		source: source,
	}
}

func (l *Lexer) Run() []Token {
	for !l.isEOF(0) {
		char := l.readChar()
		switch char {
		case '\r':
			continue
		case '\n':
			l.line += 1
			l.column = 0
		case ' ', '\t':
			l.column += 1
		default:
			if l.isDigit(char) {
				result := l.readNumber()
				if result != nil {
					l.tokens = append(l.tokens, *result)
				}
			} else if l.isChar(char) {
				l.tokens = append(l.tokens, *l.readSymbol())
			} else {
				l.errors = append(l.errors, fmt.Errorf("unknown character: %c", char))
			}
		}
	}
	l.tokens = append(l.tokens, *NewToken(TokenEOF, "", l.line, l.column))
	return l.tokens
}

func (l *Lexer) readChar() byte {
	char := l.source[l.pos]
	l.pos += 1
	l.column += 1
	return char
}

func (l *Lexer) unreadChar() {
	l.pos -= 1
	l.column -= 1
}

func (l *Lexer) peekChar(n int) byte {
	if !l.isEOF(n) {
		return l.source[l.pos+n]
	}
	return 0
}

func (l *Lexer) readNumber() *Token {
	l.unreadChar()
	start := l.pos
	startLine := l.line
	startColumn := l.column
	for l.isDigit(l.peekChar(0)) {
		l.readChar()
	}
	if l.peekChar(0) == '.' {
		l.readChar()
		if !l.isDigit(l.peekChar(0)) {
			l.errors = append(l.errors, fmt.Errorf("invalid number: %s", l.source[start:l.pos]))
			return nil
		}
		for l.isDigit(l.peekChar(0)) {
			l.readChar()
		}
	}
	value := l.source[start:l.pos]
	return NewToken(TokenNumber, value, startLine, startColumn)
}

func (l *Lexer) readSymbol() *Token {
	l.unreadChar()
	start := l.pos
	startLine := l.line
	startColumn := l.column
	for l.isChar(l.peekChar(0)) {
		l.readChar()
	}
	value := l.source[start:l.pos]
	return NewToken(TokenSymbol, value, startLine, startColumn)
}

func (l *Lexer) isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func (l *Lexer) isChar(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func (l *Lexer) isEOF(n int) bool {
	return l.pos+n >= len(l.source)
}

func (l *Lexer) Errors() []error {
	return l.errors
}
