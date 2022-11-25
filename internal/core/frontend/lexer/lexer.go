package lexer

import (
	"grimlang/internal/core/frontend/tokens"
)

type Lexer struct {
	source string
	tokens []tokens.Token

	start   int
	current int
}

func NewLexer(src string) *Lexer {
	return &Lexer{
		source: src,
	}
}

func (l *Lexer) Run() []tokens.Token {
	for !l.isAtEnd() {
		l.start = l.current
		ch := l.advance()
		switch ch {
		case ' ', '\t', '\r', '\n': // skip whitespace
			continue
		case '(':
			l.addToken(tokens.LParen, "")
		case ')':
			l.addToken(tokens.RParen, "")
		case '[':
			l.addToken(tokens.LBracket, "")
		case ']':
			l.addToken(tokens.RBracket, "")
		case '{':
			l.addToken(tokens.LBrace, "")
		case '}':
			l.addToken(tokens.RBrace, "")
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
	l.tokens = append(l.tokens, *tokens.NewToken(tt, literal))
}

func (l *Lexer) readSymbol() (tokens.TokenType, string) {
	for isLetter(l.peek(0)) && !l.isAtEnd() {
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
	literal := l.source[l.start:l.current]
	return tokens.Number, literal
}

func (l *Lexer) readString() string {
	for l.peek(0) != '"' && !l.isAtEnd() {
		l.advance()
	}
	l.advance()
	literal := l.source[l.start:l.current]
	return literal
}

func isLetter(ch byte) bool {
	res := 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
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
