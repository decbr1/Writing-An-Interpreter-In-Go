// ./cmd/token/token.go

package token

// TokenType is defined as a string to allow us to use many different values as TokenTypes,
// which in turn allows us to distinguish between different types of tokens.
// Using string also has the advantage of being easy to debug without a lot of boilerplate
// and helper functions: we can just print a string.
type TokenType string

// Token is our token data structure. It has a Type field so we can distinguish between
// "integers" and "right bracket" for example, and a Literal field that holds the literal
// value of the token, so we can reuse it later and the information whether a "number"
// token is a 5 or a 10 doesn't get lost.
type Token struct {
	Type    TokenType
	Literal string
}

// Token type constants define the possible TokenTypes in the Monkey language.
const (
	// ILLEGAL signifies a token/character we don't know about
	ILLEGAL = "ILLEGAL"
	// EOF stands for "end of file", which tells our parser later on that it can stop
	EOF = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// Two-character tokens
	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords - these look like identifiers but are part of the language
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// keywords table maps keyword strings to their TokenType constants.
// This allows us to distinguish user-defined identifiers from language keywords.
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent checks the keywords table to see whether the given identifier is in fact a
// keyword. If it is, it returns the keyword's TokenType constant. If it isn't, we just get back
// token.IDENT, which is the TokenType for all user-defined identifiers.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
