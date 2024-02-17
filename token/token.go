package token

type TokenType string

type Token struct {
    Type        TokenType
    Literal     string
}

var keywords = map[string]TokenType{
    "skrrt":        FUNCTION,
    "its_giving":   LET,
    "vibe_check":   IF,
    "i_oop":        ELSE,
    "yeet":         RETURN,
    "fax":          TRUE,
    "cap":          FALSE,
}

func LookupIdentity(identity string) TokenType {
    if token, ok := keywords[identity]; ok {
        return token
    }
    return IDENT
}

const (
    ILLEGAL     = "ILLEGAL"
    EOF         = "EOF"

    // Identifiers & Literals
    IDENT       = "IDENT" // varaible names, foo, bar, x etc.
    INT         = "INT"  // integers, 1, 54 etc.

    // Operators
    ASSIGN      = "="
    PLUS        = "+"
    MINUS       = "-"
    BANG        = "!"
    ASTERISK    = "*"
    SLASH       = "/"

    GT          = ">"
    LT          = "<"
    EQUAL       = "=="
    NOT_EQUAL   = "!="

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
    IF          = "IF"
    ELSE        = "ELSE"
    RETURN      = "RETURN"
    TRUE        = "TRUE"
    FALSE       = "FALSE"
)

