package lexer

import (
	"github.com/cmcki85/yeet-interpreter/token"
)

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

func (lexer *Lexer) PeekChar() byte {
    // Functionally the same as ReadChar, but doesn't set pos.
    if (lexer.readPos >= len(lexer.input)) {
        return 0
    } else {
        return lexer.input[lexer.readPos]
    }
}

func (lexer *Lexer) ReadNumber() string {
    // Extracts numbers from the the input string if exists.
    pos := lexer.pos
    // Walk over each digit in the word and read each char.
    for isDigit(lexer.char) {
        lexer.ReadChar()
    }
    // Return the parsed number 
    return lexer.input[pos:lexer.pos]
}

func (lexer *Lexer) ReadIdentifier() string {
    pos := lexer.pos
    // Walk over the word & read each value
    for isLetter(lexer.char) {
        lexer.ReadChar()
    }
    return lexer.input[pos:lexer.pos]
}

func (lexer *Lexer) SkipWhitespace() {
    for lexer.char == ' ' || lexer.char == '\t' || lexer.char == '\n' || lexer.char == '\r' {
        lexer.ReadChar()
    }
}

func (lexer *Lexer) NextToken() token.Token {
    var tok token.Token
    // Yeet doesn't acknowledge whitespace as significant.
    lexer.SkipWhitespace()
    switch lexer.char {
        case '=':
            if (lexer.PeekChar() == '=') {
                char := lexer.char
                lexer.ReadChar()
                tok = token.Token{ 
                    Type: token.EQUAL, 
                    Literal: string(char) + string(lexer.char), 
                }
            } else {
                tok = newToken(token.ASSIGN, lexer.char)
            }
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
        case '-':
            tok = newToken(token.MINUS, lexer.char)
        case '!':
            if (lexer.PeekChar() == '=') {
                char := lexer.char
                lexer.ReadChar()
                tok = token.Token{ 
                    Type: token.NOT_EQUAL, 
                    Literal: string(char) + string(lexer.char),
                }
            } else {
                tok = newToken(token.BANG, lexer.char)
            }
        case '/':
            tok = newToken(token.SLASH, lexer.char)
        case '*':
            tok = newToken(token.ASTERISK, lexer.char)
        case '<':
            tok = newToken(token.LT, lexer.char)
        case '>':
            tok = newToken(token.GT, lexer.char)
        case '{':
            tok = newToken(token.LBRACE, lexer.char)        
        case '}':
            tok = newToken(token.RBRACE, lexer.char)        
        case 0:
            tok.Literal = ""
            tok.Type = token.EOF
        default:
            if (isLetter(lexer.char)) {
                tok.Literal = lexer.ReadIdentifier()
                tok.Type = token.LookupIdentity(tok.Literal)
                return tok
            } else if (isDigit(lexer.char)) {
                tok.Literal = lexer.ReadNumber()
                tok.Type = token.INT
                return tok
            } else {
                tok = newToken(token.ILLEGAL, lexer.char)
            }
        }
        
        lexer.ReadChar()
        return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
    return token.Token{ Type: tokenType, Literal: string(char) }
}

func isLetter(char byte) bool {
    // Determine what is a valid identifier character
    return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
    return '0' <= char && char <= '9'
}
