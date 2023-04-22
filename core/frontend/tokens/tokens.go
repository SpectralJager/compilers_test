package tokens

type TokenType int

const (
	TokenIllegal TokenType = iota
	TokenEOF

	// delimeters
	TokenLeftParen
	TokenRightParen
	TokenLeftBracket
	TokenRightBracket
	TokenLeftBrace
	TokenRightBrace

	// special symbols
	TokenColon
	TokenDoubleColon
	TokenSemicolon
	TokenAssign
	TokenComma

	// keywords
	TokenConst
	TokenVar
	TokenSet
	TokenFn
	TokenDo
	TokenFor
	TokenWhile
	TokenEach
	TokenIf
	TokenCond

	// reserved symbols
	TokenElif
	TokenElse
	TokenFalse
	TokenTrue

	// literals
)
