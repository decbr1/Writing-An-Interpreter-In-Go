// ./cmd/lexer/lexer.go

package lexer

import "monkey/cmd/token"

// Lexer tokenizes Monkey source code. It will take source code as input and output
// the tokens that represent the source code. It goes through its input and outputs
// the next token it recognizes.
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// NextToken looks at the current character under examination (l.ch) and returns a token
// depending on which character it is. Before returning the token it advances our pointers
// into the input so when we call NextToken() again the l.ch field is already updated.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Whitespace only acts as a separator of tokens in Monkey and doesn't have meaning,
	// so we skip over it entirely
	l.skipWhitespace()

	switch l.ch {

	case '=': // We need to "peek" ahead to determine if this is = or ==
		if l.peakChar() == '=' {
			// Save l.ch in a local variable before calling l.readChar() again
			// so we don't lose the current character
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case '!': // We need to "peek" ahead to determine if this is ! or !=
		if l.peakChar() == '=' {
			// Save l.ch before advancing so we can build the two-character literal
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}

	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		// 0 is the ASCII code for "NUL" and signifies either "we haven't read anything yet"
		// or "end of file" for us
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// Check whether the current character is a letter to recognize identifiers and keywords
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			// Check if the identifier is actually a keyword
			tok.Type = token.LookupIdent(tok.Literal)
			// Early exit necessary because readIdentifier() calls readChar() repeatedly
			// and advances our positions past the last character of the identifier
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			// Early exit for same reason as identifiers
			return tok
		} else {
			// If we truly don't know how to handle the current character, declare it ILLEGAL
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	// Advance our position in the input string after creating the token
	l.readChar()
	return tok
}

// newToken is a small helper function that initializes tokens with the given type and literal.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// New returns a fully initialized *Lexer with l.ch, l.position and l.readPosition
// already set up by calling readChar() once.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	// Call readChar() to initialize the lexer so it's in a fully working state
	// before anyone calls NextToken()
	l.readChar()
	return l
}

// skipWhitespace advances the lexer past any whitespace characters.
// In Monkey, whitespace only acts as a separator and doesn't have meaning.
// Some languages like Python have significant whitespace, but we skip it entirely.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// readChar gives us the next character and advances our position in the input string.
// It checks whether we have reached the end of input. If that's the case it sets l.ch to 0,
// which is the ASCII code for "NUL". If we haven't reached the end, it sets l.ch to the
// next character by accessing l.input[l.readPosition].
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL" - signals "end of file"
	} else {
		l.ch = l.input[l.readPosition]
	}
	// l.position always points to the position where we last read
	l.position = l.readPosition
	// l.readPosition always points to the next position we're going to read from
	l.readPosition++
}

// peakChar is similar to readChar() except it doesn't increment l.position and l.readPosition.
// We only want to "peek" ahead in the input and not move around in it, so we know what
// a call to readChar() would return. This lets us look ahead to determine two-character tokens.
func (l *Lexer) peakChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// readNumber reads in a number (integer) and advances the lexer's position until it
// encounters a non-digit character. Note: Monkey only supports integers, not floats,
// hex notation, or octal notation.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readIdentifier reads in an identifier and advances the lexer's position until it
// encounters a non-letter character.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isLetter checks whether the given argument is a letter.
// Note that changing this function has a larger impact on the language our interpreter
// will be able to parse than one would expect from such a small function. It includes
// the check ch == '_', which means we treat _ as a letter and allow it in identifiers
// and keywords, enabling variable names like foo_bar.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit returns whether the passed in byte is a Latin digit between 0 and 9.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
