package lexer

import (
	"github.com/iWangMu/Go-Interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `	let five = 5;
				let ten = 10;
				let add = fn(x, y) {
					x + y;
				};
				let result = add(five, ten);
				!-/*5;
				5 < 10 > 5;
				if (5 < 10) {
					return true;
				}else {
					return false;
				}
				10 == 10;
				10 != 9;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// let five = 5;
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		// let ten = 10;
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},
		// let add = fn(x, y) {
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		// x + y;
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		// };
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		// let result = add(five, ten);
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		// !-/*5;
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		// 5 < 10 > 5;
		{token.INTEGER, "5"},
		{token.LT, "<"},
		{token.INTEGER, "10"},
		{token.GT, ">"},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		// if (5 < 10) {
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INTEGER, "5"},
		{token.LT, "<"},
		{token.INTEGER, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		// return true;
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		// }else {
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		// return false;
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		// }
		{token.RBRACE, "}"},
		// 10 == 10;
		{token.INTEGER, "10"},
		{token.EQ, "=="},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},
		// 10 != 9;
		{token.INTEGER, "10"},
		{token.NEQ, "!="},
		{token.INTEGER, "9"},
		{token.SEMICOLON, ";"},

		// End Of File
		{token.EOF, ""},
	}

	lex := New(input)

	for i, tt := range tests {
		tok := lex.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
