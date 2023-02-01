package lexer

import (
	"grimlang/internal/core/frontend/tokens"
)

type Lexer struct {
	source string
	tokens []tokens.Token

	start   int
	current int
	row     int
}

func NewLexer(src string) *Lexer {
	return &Lexer{
		source: src,
		row:    1,
	}
}

func (l *Lexer) Run() []tokens.Token {
	for !l.isAtEnd() {
		l.start = l.current
		ch := l.advance()
		switch ch {
		case ' ', '\t', '\r': // skip whitespace
			continue
		case '\n':
			l.row += 1
			continue
		case ':':
			l.addToken(tokens.Colon, ":")
		case '(':
			l.addToken(tokens.LParen, "(")
		case ')':
			l.addToken(tokens.RParen, ")")
		case '[':
			l.addToken(tokens.LBracket, "[")
		case ']':
			l.addToken(tokens.RBracket, "]")
		case '"':
			literal := l.readString()
			l.addToken(tokens.String, literal)
		default:
			if isLetter(ch) {
				tt, symbol := l.readSymbol()
				l.addToken(tt, symbol)
			} else if isDigit(ch) {
				tt, digit := l.readDigit()
				l.addToken(tt, digit)
			} else {
				l.addToken(tokens.Illegal, "Illegal character")
			}
		}
	}
	l.addToken(tokens.EOF, "EOF")
	return l.tokens
}

func (l *Lexer) advance() byte {
	ch := l.source[l.current]
	l.current += 1
	return ch
}

func (l *Lexer) peek(offset int) byte {
	if l.current+offset >= len(l.source) {
		return '0'
	}
	ch := l.source[l.current+offset]
	return ch
}

func (l *Lexer) addToken(tt tokens.TokenType, literal string) {
	if literal == "" {
		literal = l.source[l.start:l.current]
	}
	l.tokens = append(l.tokens, *tokens.NewToken(tt, literal, l.row))
}

func (l *Lexer) readSymbol() (tokens.TokenType, string) {
	for (isLetter(l.peek(0)) || isDigit(l.peek(0)) || l.peek(0) == '_') && !l.isAtEnd() {
		l.advance()
	}
	literal := l.source[l.start:l.current]
	tt := tokens.LookupSymbolType(literal)
	return tt, literal
}

func (l *Lexer) readDigit() (tokens.TokenType, string) {
	for isDigit(l.peek(0)) && !l.isAtEnd() {
		l.advance()
	}
	if l.peek(0) == '.' {
		l.advance()
		for isDigit(l.peek(0)) && !l.isAtEnd() {
			l.advance()
		}
		literal := l.source[l.start:l.current]
		return tokens.Float, literal
	}
	literal := l.source[l.start:l.current]
	return tokens.Int, literal
}

func (l *Lexer) readString() string {
	for l.peek(0) != '"' && !l.isAtEnd() {
		l.advance()
	}
	l.advance()
	literal := l.source[l.start+1 : l.current-1]
	return literal
}

func isLetter(ch byte) bool {
	res := 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
	return res
}

func isDigit(ch byte) bool {
	res := '0' <= ch && ch <= '9'
	return res
}

func (l *Lexer) isAtEnd() bool {
	res := l.current >= len(l.source)
	return res
}
