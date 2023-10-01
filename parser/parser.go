package parser

import (
	"fmt"
	"github.com/iWangMu/Go-Interpreter/ast"
	"github.com/iWangMu/Go-Interpreter/lexer"
	"github.com/iWangMu/Go-Interpreter/token"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // >, <
	SUM         // +
	PRODUCT     // *
	PREFIX      // !x, -x
	CALL        // myFunc(x)
)

// 前缀解析函数
type prefixParseFn func() ast.Expression

// 中缀解析函数
type infixParseFn func(ast.Expression) ast.Expression

type Parser struct {
	// 词法单元
	lex    *lexer.Lexer
	errors []string

	// 当前词法单元
	curToken token.Token
	// 下一个词法单元
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex, errors: []string{}}
	// 注册前缀解析函数
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(token.INTEGER, p.parseIntegerLiteral)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)

	// 读取两个词法单元，设置curToken和peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {

	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) registerPrefix(tt token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tt] = fn
}
func (p *Parser) registerInfix(tt token.TokenType, fn infixParseFn) {
	p.infixParseFns[tt] = fn
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead.", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

/**
 * 解析语句
 * - let
 * - return
 */
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.SEMICOLON:
		return nil
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO: 处理表达式，直到遇到分号，表示语句结束
	if !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	// 下一个词法单元
	p.nextToken()
	// TODO: 处理表达式，直到遇到分号，表示语句结束
	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)

	// 分号(;)表示语句结束
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()
	return leftExp
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	il := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer.", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	il.Value = value
	return il
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

/**
 * 判断下一个词法单元是否是预期类型的词法单元
 * - 是: 向前解析下一个词法单元
 * - 否: 生成一条错误消息
 */
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) currTokenIs(t token.TokenType) bool {

	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {

	return p.peekToken.Type == t
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
