package token

type TokenType string

type Token struct {
    Type        TokenType
    Literal     string
}

const (
    ILLEGAL     = "ILLEGAL"
    EOF         = "EOF"

    // Identifiers & Literals
    IDENT       = "IDEN" // varaible names, foo, bar, x etc.
    INT         = "INT"  // integers, 1, 54 etc.

    // Operators
    ASSIGN      = "="
    PLUS        = "+"

    // Delimiters
    COMMA       = ","
    SEMICOLON   = ";"
    
    LPAREN      = "("
    RPAREN      = ")"
    LBRACE      = "{"
    RBRACE      = "}"

    // Keywords
    FUNCTION    = "FUNCTION"
    LET         = "LET"
)

