package syntax

import (
	"fmt"
)

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
		case '(':
			l.tokens = append(l.tokens, *NewToken(TokenLParen, "(", l.line, l.column))
		case ')':
			l.tokens = append(l.tokens, *NewToken(TokenRParen, ")", l.line, l.column))
		case '[':
			l.tokens = append(l.tokens, *NewToken(TokenLBracket, "[", l.line, l.column))
		case ']':
			l.tokens = append(l.tokens, *NewToken(TokenRBracket, "]", l.line, l.column))
		case '{':
			l.tokens = append(l.tokens, *NewToken(TokenLCBracket, "{", l.line, l.column))
		case '}':
			l.tokens = append(l.tokens, *NewToken(TokenRCBracket, "}", l.line, l.column))
		case '<':
			l.tokens = append(l.tokens, *NewToken(TokenLess, "<", l.line, l.column))
		case '>':
			l.tokens = append(l.tokens, *NewToken(TokenMore, ">", l.line, l.column))
		case ';':
			l.tokens = append(l.tokens, *NewToken(TokenSemicolon, ";", l.line, l.column))
		case '=':
			l.tokens = append(l.tokens, *NewToken(TokenAssign, "=", l.line, l.column))
		case '@':
			tempToken := *NewToken(TokenIllegal, "@", l.line, l.column)
			if l.isChar(l.readChar()) {
				symbolToken := l.readSymbol()
				tempToken.Value += symbolToken.Value
				if tp, ok := keywords[tempToken.Value]; ok {
					tempToken.Type = tp
					l.tokens = append(l.tokens, tempToken)
					continue
				}
				l.errors = append(l.errors, fmt.Errorf("unexpected keyword: %s", tempToken))
				continue
			}
			l.errors = append(l.errors, fmt.Errorf("expected keyword, got: %c at %d:%d", l.peekChar(-1), l.line, l.column))
		case '/':
			l.tokens = append(l.tokens, *NewToken(TokenSlash, "/", l.line, l.column))
		case ':':
			if l.peekChar(0) == ':' {
				l.tokens = append(l.tokens, *NewToken(TokenDColon, "::", l.line, l.column))
				l.readChar()
			} else {
				l.tokens = append(l.tokens, *NewToken(TokenColon, ":", l.line, l.column))
			}
		case '"':
			result := l.readString()
			if result != nil {
				l.tokens = append(l.tokens, *result)
			}
		default:
			if l.isDigit(char) {
				result := l.readNumber()
				if result != nil {
					l.tokens = append(l.tokens, *result)
				}
			} else if l.isChar(char) {
				l.tokens = append(l.tokens, *l.readSymbol())
			} else {
				l.errors = append(l.errors, fmt.Errorf("unknown character: %c at %d:%d", char, l.line, l.column))
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

func (l *Lexer) readString() *Token {
	start := l.pos
	startLine := l.line
	startColumn := l.column
	for l.peekChar(0) != '"' && !l.isEOF(0) && l.peekChar(0) != '\n' {
		l.readChar()
	}
	if l.isEOF(0) || l.peekChar(0) == '\n' {
		l.errors = append(l.errors, fmt.Errorf("unterminated string at %d:%d", startLine, startColumn))
		return nil
	}
	value := l.source[start:l.pos]
	l.readChar()
	return NewToken(TokenString, value, startLine, startColumn)
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
			l.errors = append(l.errors, fmt.Errorf("invalid number: %s at %d:%d", l.source[start:l.pos], startLine, startColumn))
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
