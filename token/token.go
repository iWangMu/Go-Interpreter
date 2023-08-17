package token

type TokenType string // 词法单元类型

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // 未知的词法单元
	EOF     = "EOF"     // 文件结尾(End Of File)

	IDENTIFIER = "IDENTIFIER" // 标识符: x, y
	INTEGER    = "INTEGER"    // 整型字面量: 1,2,3,...

	TRUE  = "TRUE"
	FALSE = "FALSE"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"

	LT  = "<"
	GT  = ">"
	EQ  = "=="
	NEQ = "!="

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var KEYWORDS = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := KEYWORDS[ident]; ok {
		return tok
	}
	return IDENTIFIER
}
