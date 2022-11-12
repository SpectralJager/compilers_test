package lexer

import (
	"bufio"
	"grimlang/internal/core/frontend/tokens"
	"grimlang/internal/core/frontend/utils"
	symboltable "grimlang/internal/core/symbol_table"
	"io"
)

type Lexer struct {
	position utils.Position
	reader   *bufio.Reader
	tokens   chan tokens.Token
}

func NewLexer(reader io.Reader, tokenChan chan tokens.Token) (*Lexer, chan tokens.Token) {
	l := &Lexer{
		tokens:   tokenChan,
		reader:   bufio.NewReader(reader),
		position: utils.Position{Line: 0, Column: 0},
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
				l.emit(tokens.EOF, "")
				break
			}
			panic(err)
		}
		l.position.Column += 1
		switch r {
		// special cases
		case '\n':
			l.nextLine()
		// separators
		case '(':
			l.emit(tokens.LeftParen, "(")
		case ')':
			l.emit(tokens.RightParen, ")")
		case '[':
			l.emit(tokens.LeftSBracet, "[")
		case ']':
			l.emit(tokens.RightSBracet, "]")
		case '{':
			l.emit(tokens.LeftCBracet, "{")
		case '}':
			l.emit(tokens.RightCBracet, "}")
		// math operators
		case '+':
			l.emit(tokens.Plus, "+")
		case '-':
			l.emit(tokens.Minus, "-")
		case '*':
			l.emit(tokens.Multimply, "*")
		case '/':
			l.emit(tokens.Divide, "/")
		// operators
		case '\'':
			l.emit(tokens.Quote, "'")
		default:
			if isWhiteSpace(r) {
				continue
			} else if isDigit(r) {
				l.readDigit()
			} else if isLetter(r) {
				l.readIdentifier()
			} else {
				l.emit(tokens.Illegal, string(r))
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
		l.position.Column += 1
		if isDigit(r) {
			literal += string(r)

		} else {
			l.backup()
			break
		}
	}
	l.emit(tokens.Int, literal)
}

func (l *Lexer) readIdentifier() {
	l.backup()
	var literal string
	for {
		r, _, err := l.reader.ReadRune()
		errOrUnexpectedEOF(err, l)
		l.position.Column += 1
		if isLetter(r) {
			literal += string(r)

		} else {
			l.backup()
			break
		}
	}
	kwType, ok := tokens.Keywords[literal]
	if ok {
		l.emit(kwType, literal)
		return
	}
	st := symboltable.GetSymbolTableEntity()
	st.Save(literal, nil)
	l.emit(tokens.Identifier, literal)
}

func (l *Lexer) emit(t tokens.TokenType, value string) {
	l.tokens <- tokens.Token{Type: t, Literal: value, Position: l.position}
}

func (l *Lexer) nextLine() {
	l.position.Line += 1
	l.position.Column = 0
}

func (l *Lexer) peekRune() rune {
	r, _, err := l.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			l.emit(tokens.EOF, "")
		} else {
			panic(err)
		}
	}
	l.backup()
	return r
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}
	l.position.Column--
}

// Utils
func errOrUnexpectedEOF(err error, l *Lexer) {
	if err != nil {
		if err == io.EOF {
			l.emit(tokens.EOF, "0")
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

func isLetter(ch rune) bool {
	return ('a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_')
}
