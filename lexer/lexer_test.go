package lexer

import (
    "testing"

    "github.com/cmcki85/yeet-interpreter/token"
)

func TestNextToken(t *testing.T) {
    input := `its_giving five = 5;
    its_giving ten = 10;

    its_giving add = skrrt(x, y) {
        x + y;
    }

    its_giving result = add(five, ten);`

    tests := []struct {
        expectedType    token.TokenType
        expectedLiteral string
    }{
        { token.LET, "its_giving" },
        { token.IDENT, "five" },
        { token.ASSIGN, "=" },
        { token.INT, "5" },
        { token.SEMICOLON, ";" },
        { token.LET, "its_giving" },
        { token.IDENT, "ten" },
        { token.ASSIGN, "=" },
        { token.INT, "10" },
        { token.SEMICOLON, ";" },
        { token.LET, "its_giving" },
        { token.IDENT, "add" },
        { token.ASSIGN, "=" },
        { token.FUNCTION, "skrrt" },
        { token.LPAREN, "(" },
        { token.IDENT, "x" },
        { token.COMMA, "," },
        { token.IDENT, "y" },
        { token.RPAREN, ")" },
        { token.LBRACE, "{" },
        { token.IDENT, "x" },
        { token.PLUS, "+" },
        { token.IDENT, "y" },
        { token.SEMICOLON, ";" },
        { token.RBRACE, "}" },
        { token.LET, "its_giving" },
        { token.IDENT, "result" },
        { token.ASSIGN, "=" },
        { token.IDENT, "add" },
        { token.LPAREN, "(" },
        { token.IDENT, "five" },
        { token.COMMA, "," },
        { token.IDENT, "ten" },
        { token.RPAREN, ")" },
        { token.SEMICOLON, ";" },
        { token.EOF, "" },
    }

    lexer := New(input) // Init the Lexer

    for i, test := range tests {
        token := lexer.NextToken()
        if (token.Type != test.expectedType) {
            t.Fatalf(
                "tests[%d] - tokenType incorrect! expected: %q got %q",
                i,
                test.expectedType,
                token.Type,
            ) 
        }

        if (token.Literal != test.expectedLiteral) {
            t.Fatalf(
                "tests[%d] - literal incorrect! expected: %q got %q",
                i,
                test.expectedLiteral,
                token.Literal,
            )
        }
    }
}

func TestOperatorTokens(t *testing.T) {
    input := `its_giving five = 5;
        its_giving ten = 10;
        !-/*5;
        5 < 10 > 5;

        vibe_check(5 < 10) {
            yeet fax;
        } i_oop {
            yeet cap;
        }
    `

    tests := []struct {
        expectedType    token.TokenType
        expectedLiteral string
    }{
        { token.LET, "its_giving" },
        { token.IDENT, "five" },
        { token.ASSIGN, "=" },
        { token.INT, "5" },
        { token.SEMICOLON, ";" },
        { token.LET, "its_giving" },
        { token.IDENT, "ten" },
        { token.ASSIGN, "=" },
        { token.INT, "10" },
        { token.SEMICOLON, ";" },
        { token.BANG, "!" },
        { token.MINUS, "-" },
        { token.SLASH, "/" },
        { token.ASTERISK, "*"},
        { token.INT, "5" },
        { token.SEMICOLON, ";" },
        { token.INT, "5" },
        { token.LT, "<" },
        { token.INT, "10" },
        { token.GT, ">" },
        { token.INT, "5" },
        { token.SEMICOLON, ";" },
        { token.IF, "vibe_check" },
        { token.LPAREN, "(" },
        { token.INT, "5" },
        { token.LT, "<" },
        { token.INT, "10" },
        { token.RPAREN, ")" },
        { token.LBRACE, "{" },
        { token.RETURN, "yeet" },
        { token.TRUE, "fax" },
        { token.SEMICOLON, ";" },
        { token.RBRACE, "}" },
        { token.ELSE, "i_oop" },
        { token.LBRACE, "{" },
        { token.RETURN, "yeet" },
        { token.FALSE, "cap" },
        { token.SEMICOLON, ";" },
        { token.RBRACE, "}" },
    }

    lexer := New(input) // Init the Lexer

    for i, test := range tests {
        token := lexer.NextToken()
        if (token.Type != test.expectedType) {
            t.Fatalf(
                "tests[%d] - tokenType incorrect! expected: %q got %q",
                i,
                test.expectedType,
                token.Type,
            ) 
        }

        if (token.Literal != test.expectedLiteral) {
            t.Fatalf(
                "tests[%d] - literal incorrect! expected: %q got %q",
                i,
                test.expectedLiteral,
                token.Literal,
            )
        }
    }
}
