package lexer

import (
	"grimlang/internal/core/frontend/tokens"
	"io"
)

type Lexer struct {
	code string // source code

	// line         int // current line
	// column       int // current column
	position     int // current position
	nextPosition int // next position
}

func NewLexer(src io.Reader) *Lexer {
	code, err := io.ReadAll(src)
	if err != nil {
		panic(err)
	}
	lexer := &Lexer{
		code: string(code),
		// line:     0,
		// column:   0,
	}

	return lexer
}

func (l *Lexer) Run() []tokens.Token {
	var tokList []tokens.Token

	for {
		ch := l.readChar()
		switch ch {
		case '0':
			tokList = append(tokList, *tokens.NewToken(tokens.EOF, ""))
			return tokList
		case '\n':
		case '(':
			tokList = append(tokList, *tokens.NewToken(tokens.LParen, "("))
		case ')':
			tokList = append(tokList, *tokens.NewToken(tokens.RParen, ")"))
		case '{':
			tokList = append(tokList, *tokens.NewToken(tokens.LBrace, "{"))
		case '}':
			tokList = append(tokList, *tokens.NewToken(tokens.RBrace, "}"))
		case '[':
			tokList = append(tokList, *tokens.NewToken(tokens.LBracket, "["))
		case ']':
			tokList = append(tokList, *tokens.NewToken(tokens.RBracket, "]"))
		default:
			if l.isWhitespace(ch) {
			} else if l.isLetter(ch) {
				tok := l.readSymbol()
				tokList = append(tokList, tok)
			} else if l.isDigit(ch) {
				tok := l.readNumber()
				tokList = append(tokList, tok)
			} else {
				tokList = append(tokList, *tokens.NewToken(tokens.Illegal, string(ch)))
			}
		}
	}

}

func (l *Lexer) readChar() byte {
	ch := l.peekChar()
	l.position = l.nextPosition
	l.nextPosition += 1
	return ch
}

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.code) {
		return '0'
	}
	return l.code[l.nextPosition]
}

func (l *Lexer) readSymbol() tokens.Token {
	pos := l.position
	for l.isLetter(l.peekChar()) {
		l.readChar()
	}
	val := l.code[pos : l.position+1]
	return *tokens.NewToken(tokens.LookupSymbolType(val), val)
}

func (l *Lexer) readNumber() tokens.Token {
	pos := l.position
	for l.isDigit(l.peekChar()) {
		l.readChar()
	}
	val := l.code[pos : l.position+1]
	return *tokens.NewToken(tokens.Number, val)
}

func (l *Lexer) isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\r' || ch == '\t'
}
