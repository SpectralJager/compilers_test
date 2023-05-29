package frontend

import (
	"fmt"
)

type Lexer struct {
	input  string
	tokens []Token
	errors []error

	pos    int
	line   int
	column int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: input,
	}
}

func (l *Lexer) Lex() *[]Token {
	for !l.isEOF() {
		ch := l.peek(0)
		switch ch {
		case '\n':
			l.next()
			l.line++
			l.column = 0
		case '\r':
			l.next()
		case ' ', '\t':
			l.next()
		case '@':
			tok := NewToken(TokenIllegal, "", l.line, l.column)
			startPos := l.pos
			l.next()
			for {
				ch = l.peek(0)
				if !(ch >= 'a' && ch <= 'z') {
					break
				}
				l.next()
			}
			tok.Value += l.input[startPos:l.pos]
			tt, ok := IsKeyword(tok.Value)
			if !ok {
				l.errors = append(l.errors, fmt.Errorf("unknown keyword %s", tok.String()))
			}
			tok.Type = tt
			l.tokens = append(l.tokens, *tok)
		case '(':
			l.tokens = append(l.tokens, *NewToken(TokenLeftParen, string(l.next()), l.line, l.column-1))
		case ')':
			l.tokens = append(l.tokens, *NewToken(TokenRightParen, string(l.next()), l.line, l.column-1))
		case '{':
			l.tokens = append(l.tokens, *NewToken(TokenLeftBrace, string(l.next()), l.line, l.column-1))
		case '}':
			l.tokens = append(l.tokens, *NewToken(TokenRightBrace, string(l.next()), l.line, l.column-1))
		case '=':
			l.tokens = append(l.tokens, *NewToken(TokenAssign, string(l.next()), l.line, l.column-1))
		case '/':
			l.tokens = append(l.tokens, *NewToken(TokenSlash, string(l.next()), l.line, l.column-1))
		case ':':
			if l.peek(1) == ':' {
				tok := NewToken(TokenDoubleColon, string(l.next()), l.line, l.column)
				tok.Value += string(l.next())
				l.tokens = append(l.tokens, *tok)
			} else {
				l.tokens = append(l.tokens, *NewToken(TokenColon, string(l.next()), l.line, l.column-1))
			}
		case ';':
			l.tokens = append(l.tokens, *NewToken(TokenSemicolon, string(l.next()), l.line, l.column-1))
		case '"':
			tok := NewToken(TokenString, "", l.line, l.column)
			startPos := l.pos
			l.next() // eat "
			for l.isPrintable(l.peek(0)) && !l.isEOF() && l.peek(0) != '\n' {
				l.next()
			}
			if l.isEOF() || l.peek(0) == '\n' {
				l.errors = append(l.errors, fmt.Errorf("unterminated string at %d:%d", l.line, l.column))
				continue
			}
			l.next() // eat "
			tok.Value += l.input[startPos:l.pos]
			l.tokens = append(l.tokens, *tok)
		default:
			if l.isChar(ch) {
				tok := NewToken(TokenIllegal, "", l.line, l.column)
				startPos := l.pos
				l.next()
				for l.isChar(l.peek(0)) && !l.isEOF() {
					l.next()
				}
				tok.Value += l.input[startPos:l.pos]
				tt, ok := IsReserved(tok.Value)
				if !ok {
					tok.Type = TokenSymbol
				} else {
					tok.Type = tt
				}
				l.tokens = append(l.tokens, *tok)
			} else if l.isDigit(ch) {
				tok := NewToken(TokenNumber, "", l.line, l.column)
				startPos := l.pos
				l.next()
				for l.isDigit(l.peek(0)) && !l.isEOF() {
					l.next()
				}
				if l.peek(0) == '.' {
					l.next()
					for l.isDigit(l.peek(0)) && !l.isEOF() {
						l.next()
					}
				}
				tok.Value += l.input[startPos:l.pos]
				l.tokens = append(l.tokens, *tok)

			} else {
				l.next()
				l.errors = append(l.errors, fmt.Errorf("unexpected character '%c' at %d:%d", ch, l.line, l.column))
			}
		}
	}
	l.tokens = append(l.tokens, *NewToken(TokenEOF, "", l.line, l.column))
	return &l.tokens
}

func (l *Lexer) peek(n int) byte {
	if l.pos+n >= len(l.input) {
		return 0
	}
	return l.input[l.pos+n]
}

func (l *Lexer) next() byte {
	if l.isEOF() {
		return 0
	}
	c := l.input[l.pos]
	l.pos++
	l.column++
	return c
}

func (l *Lexer) isEOF() bool {
	return l.pos >= len(l.input)
}

func (l *Lexer) isChar(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func (l *Lexer) isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) isPrintable(ch byte) bool {
	return 0x20 <= ch && ch <= 0x7E && ch != '"'
}

func (l *Lexer) Error() []error {
	return l.errors
}
