package models

import (
	"bufio"
	"io"
)

type Lexer struct {
	position Position
	reader   *bufio.Reader
	tokens   chan Token
}

func NewLexer(reader io.Reader) (*Lexer, chan Token) {
	l := &Lexer{
		tokens:   make(chan Token),
		reader:   bufio.NewReader(reader),
		position: Position{line: 1, column: 0},
	}
	go l.run()
	return l, l.tokens
}

func (l *Lexer) run() {
	defer close(l.tokens)
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				l.emit(EOF, "")
				break
			}
			panic(err)
		}
		l.position.column += 1
		switch r {
		case '\n':
			l.nextLine()
		case '(':
			l.emit(LeftParen, "(")
		case ')':
			l.emit(RightParen, ")")
		case '+':
			l.emit(Plus, "+")
		default:
			if isWhiteSpace(r) {
				continue
			} else if isDigit(r) {
				l.readDigit()
			}
		}
	}
}

func (l *Lexer) readDigit() {
	l.backup()
	var literal string
	for {
		r, _, err := l.reader.ReadRune()
		errOrUnexpectedEOF(err, l)
		l.position.column += 1
		if isDigit(r) {
			literal += string(r)

		} else {
			break
		}
	}
	l.backup()
	l.emit(Int, literal)
}

func (l *Lexer) emit(t TokenType, value string) {
	l.tokens <- Token{Type: t, Literal: value, Position: l.position}
}

func (l *Lexer) nextLine() {
	l.position.line += 1
	l.position.column = 0
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}
	l.position.column--
}

// Utils
func errOrUnexpectedEOF(err error, l *Lexer) {
	if err != nil {
		if err == io.EOF {
			l.emit(EOF, "0")
			panic(io.ErrUnexpectedEOF)
		}
		panic(err)
	}
}

func isWhiteSpace(ch rune) bool {
	return (ch == ' ' || ch == '\t' || ch == '\r')
}

func isDigit(ch rune) bool {
	return ('0' <= ch && ch <= '9')
}
