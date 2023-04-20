package syntax

import "fmt"

type Parser struct {
	tokens []Token
	errors []error

	pos int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) Run() *Program {
	var programm Program
	for !p.isEOF() {
		tok := p.readToken()
		switch tok.Type {
		case TokenConst:
			res := p.parseConst()
			if res != nil {
				programm.Body = append(programm.Body, res)
			}
		case TokenFunc:
			res := p.parseFunc()
			if res != nil {
				programm.Body = append(programm.Body, res)
			}
		default:
			p.errors = append(p.errors, fmt.Errorf("unknown token: %s", tok))
		}
	}
	return &programm
}

// const	= '@const' SYMBOL ':' SYMBOL '=' LITTERAL ';' ;
func (p *Parser) parseConst() *ConstNode {
	var constNode ConstNode
	res := p.parseSymbolDef()
	if res == nil {
		return nil
	}
	constNode.Symbol = *res
	if tok := p.readToken(); tok.Type != TokenAssign {
		p.errors = append(p.errors, fmt.Errorf("expected assign, got %s", tok))
		return nil
	}
	if tok := p.readToken(); tok.Type != TokenInteger {
		p.errors = append(p.errors, fmt.Errorf("expected number, got %s", tok))
		return nil
	} else {
		constNode.Value = tok
	}
	if tok := p.readToken(); tok.Type != TokenSemicolon {
		p.errors = append(p.errors, fmt.Errorf("expected semicolon, got %s", tok))
		return nil
	}
	return &constNode
}

// func	= '@fn' SYMBOL '(' symbolDef* ')' SYMBOL? '{' localUnion* '}' ;
func (p *Parser) parseFunc() *FunctionNode {
	var funcNode FunctionNode
	if tok := p.readToken(); tok.Type != TokenSymbol {
		p.errors = append(p.errors, fmt.Errorf("expected symbol, got %s", tok))
		return nil
	} else {
		funcNode.Symbol = tok
	}
	if tok := p.readToken(); tok.Type != TokenLeftParen {
		p.errors = append(p.errors, fmt.Errorf("expected '(', got %s", tok))
		return nil
	}
	for p.peek(0).Type != TokenRightParen {
		funcNode.Args = append(funcNode.Args, *p.parseSymbolDef())
	}
	if tok := p.readToken(); tok.Type != TokenRightParen {
		p.errors = append(p.errors, fmt.Errorf("expected ')', got %s", tok))
		return nil
	}
	if tok := p.peek(0); tok.Type == TokenSymbol {
		funcNode.ReturnType = p.readToken()
	}
	if tok := p.readToken(); tok.Type != TokenLeftBrace {
		p.errors = append(p.errors, fmt.Errorf("expected '{', got %s", tok))
		return nil
	}
	if tok := p.readToken(); tok.Type != TokenRightBrace {
		p.errors = append(p.errors, fmt.Errorf("expected '}', got %s", tok))
		return nil
	}

	return &funcNode
}

// symbolDef = SYMBOL ':' SYMBOL ;
func (p *Parser) parseSymbolDef() *SymbolDef {
	var symbolDef SymbolDef
	if tok := p.readToken(); tok.Type != TokenSymbol {
		p.errors = append(p.errors, fmt.Errorf("expected symbol, got %s", tok))
		return nil
	} else {
		symbolDef.Symbol = tok
	}
	if tok := p.readToken(); tok.Type != TokenColon {
		p.errors = append(p.errors, fmt.Errorf("expected colon, got %s", tok))
		return nil
	}
	if tok := p.readToken(); tok.Type != TokenSymbol {
		p.errors = append(p.errors, fmt.Errorf("expected symbol, got %s", tok))
		return nil
	} else {
		symbolDef.Type = tok
	}
	return &symbolDef
}

func (p *Parser) readToken() Token {
	tok := p.tokens[p.pos]
	p.pos += 1
	return tok
}

func (p *Parser) peek(i int) Token {
	if p.pos+i >= len(p.tokens) {
		return p.tokens[len(p.tokens)-1]
	}
	return p.tokens[p.pos+i]
}

func (p *Parser) isEOF() bool {
	return p.tokens[p.pos].Type == TokenEOF
}

func (p *Parser) Errors() []error {
	return p.errors
}
