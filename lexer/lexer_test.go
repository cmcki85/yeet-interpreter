package lexer

import (
    "testing"

    "github.com/cmcki85/yeet-interpreter/token"
)

func TestNextToken(t *testing.T) {
    input := `=+(){},;`

    tests := []struct {
        expectedType    token.TokenType
        expectedLiteral string
    }{
        { token.ASSIGN, "=" },
        { token.PLUS, "+" },
        { token.LPAREN, "(" },
        { token.RPAREN, ")" },
        { token.LBRACE, "{" },
        { token.RBRACE, "}" },
        { token.COMMA, "," },
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
