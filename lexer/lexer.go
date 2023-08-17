package lexer

import "github.com/iWangMu/Go-Interpreter/token"

type Lexer struct {
	input        string
	position     int  // 所输入字符串中的当前位置, 指向当前字符
	readPosition int  // 所输入字符串中的当前读取位置, 指向当前字符之后的一个字符
	ch           byte // 当前正在查看的字符
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	// 初始化 position, readPosition和ch
	lexer.readChar()
	return lexer
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token
	// 跳过空白字符
	lex.eatWhitespace()
	switch lex.ch {
	case '=':
		if lex.peekChar() == '=' {
			lex.readChar()
			tok.Type = token.EQ
			tok.Literal = "=="
		} else {
			tok = newToken(token.ASSIGN, lex.ch)
		}
	case '(':
		tok = newToken(token.LPAREN, lex.ch)
	case ')':
		tok = newToken(token.RPAREN, lex.ch)
	case '{':
		tok = newToken(token.LBRACE, lex.ch)
	case '}':
		tok = newToken(token.RBRACE, lex.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lex.ch)
	case ',':
		tok = newToken(token.COMMA, lex.ch)
	case '+':
		tok = newToken(token.PLUS, lex.ch)
	case '-':
		tok = newToken(token.MINUS, lex.ch)
	case '*':
		tok = newToken(token.ASTERISK, lex.ch)
	case '/':
		tok = newToken(token.SLASH, lex.ch)
	case '!':
		if lex.peekChar() == '=' {
			lex.readChar()
			tok.Type = token.NEQ
			tok.Literal = "!="
		} else {
			tok = newToken(token.BANG, lex.ch)
		}
	case '>':
		tok = newToken(token.GT, lex.ch)
	case '<':
		tok = newToken(token.LT, lex.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lex.ch) {
			tok.Literal = lex.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(lex.ch) {
			tok.Literal = lex.readNumber()
			tok.Type = token.INTEGER
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lex.ch)
		}
	}

	lex.readChar()
	return tok
}

/**
 * 读取下一个字符，并前移在input中的位置
 * 该词法分析器仅支持 ASCII 字符，不能 支持所有的 Unicode 字符。
 * 要完全支持 Unicode 和 UTF-8，就要将 l.ch 的类型从 byte 改 为 rune，同时还要修改读取下一个字符的方式。
 */
func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readPosition]
	}
	lex.position = lex.readPosition
	lex.readPosition += 1
}

func (lex *Lexer) readIdentifier() string {
	fromPosition := lex.position
	for isLetter(lex.ch) {
		lex.readChar()
	}
	endPosition := lex.position

	return lex.input[fromPosition:endPosition]
}

func (lex *Lexer) readNumber() string {
	fromPosition := lex.position
	for isDigit(lex.ch) {
		lex.readChar()
	}
	endPosition := lex.position

	return lex.input[fromPosition:endPosition]
}

func (lex *Lexer) eatWhitespace() {
	for lex.ch == ' ' || lex.ch == '\n' || lex.ch == '\t' || lex.ch == '\r' {
		lex.readChar()
	}
}

func (lex *Lexer) peekChar() byte {
	if lex.readPosition >= len(lex.input) {
		return 0
	} else {
		return lex.input[lex.readPosition]
	}
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') ||
		('A' <= ch && ch <= 'Z') ||
		(ch == '_')
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
