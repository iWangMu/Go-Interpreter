package parser

import (
	"github.com/iWangMu/Go-Interpreter/ast"
	"github.com/iWangMu/Go-Interpreter/lexer"
	"github.com/iWangMu/Go-Interpreter/token"
)

type Parser struct {
	lex       *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex}
	// 读取两个词法单元，设置curToken和peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
