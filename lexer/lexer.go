package lexer

import "github.com/cmcki85/yeet-interpreter/token"

type Lexer struct {
    input   string
    pos     int     // Current position in input
    readPos int     // Current reading position in input (next char after current char)
    char    byte    // Current char being processed
}

func New(input string) *Lexer {
    lexer := &Lexer{ input: input }
    lexer.ReadChar()

    return lexer
}

func (lexer *Lexer) ReadChar() {
    // Extract the first character from the input sting if exists.
    if (lexer.readPos >= len(lexer.input)) {
        lexer.char = 0 // 0 represents EOF or NIL
    } else {
        lexer.char = lexer.input[lexer.readPos]
    }

    // Set the current position of the lexer.
    lexer.pos = lexer.readPos
    // Set the new read position of the lexer.
    lexer.readPos += 1
}

func (lexer *Lexer) NextToken() token.Token {
    var tok token.Token
    
    switch lexer.char {
        case '=':
            tok = newToken(token.ASSIGN, lexer.char)
        case ';':
            tok = newToken(token.SEMICOLON, lexer.char)        
        case '(':
            tok = newToken(token.LPAREN, lexer.char)        
        case ')':
            tok = newToken(token.RPAREN, lexer.char)        
        case ',':
            tok = newToken(token.COMMA, lexer.char)        
        case '+':
            tok = newToken(token.PLUS, lexer.char)        
        case '{':
            tok = newToken(token.LBRACE, lexer.char)        
        case '}':
            tok = newToken(token.RBRACE, lexer.char)        
        case 0:
            tok.Literal = ""
            tok.Type = token.EOF
        }
        
        lexer.ReadChar()
        return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
    return token.Token{ Type: tokenType, Literal: string(char) }
}
