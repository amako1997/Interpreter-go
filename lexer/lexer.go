package lexer

import "interpreter/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input:        input,
		position:     0,
		readPosition: 0,
	}
	l.readChar()
	return l
}
func (this *Lexer) readChar() {

	if this.readPosition >= len(this.input) {
		this.ch = 0
	} else {
		this.ch = this.input[this.readPosition]
	}
	this.position = this.readPosition
	this.readPosition += 1

}
func (this *Lexer) seeNextChar() byte {
	if this.readPosition >= len(this.input) {

		return 0
	} else {
		ch := this.input[this.readPosition]
		return ch
	}
}
func (this *Lexer) NextToken() token.Token {

	var tok token.Token
	this.skipWitheSpace()

	switch this.ch {

	case '=':
		if this.seeNextChar() == '=' {
			// save current character
			ch := this.ch
			this.readChar()
			tok = token.Token{Type: token.EQ, Value: string(ch) + string(this.ch)}
		} else {
			tok = makeToken(token.ASSIGN, this.ch)
		}
	case ';':
		tok = makeToken(token.SEMICOLON, this.ch)
	case '(':
		tok = makeToken(token.LPAREN, this.ch)
	case ')':
		tok = makeToken(token.RPAREN, this.ch)
	case ',':
		tok = makeToken(token.COMMA, this.ch)
	case '+':
		tok = makeToken(token.PLUS, this.ch)
	case '{':
		tok = makeToken(token.LBRACE, this.ch)
	case '}':
		tok = makeToken(token.RBRACE, this.ch)
	case '-':
		tok = makeToken(token.MINUS, this.ch)
	case '!':
		if this.seeNextChar() == '=' {
			// save current character
			ch := this.ch
			this.readChar()
			tok = token.Token{Type: token.NOT_EQ, Value: string(ch) + string(this.ch)}
		} else {
			tok = makeToken(token.BANG, this.ch)
		}
	case '/':
		tok = makeToken(token.SLASH, this.ch)
	case '*':
		tok = makeToken(token.ASTERISK, this.ch)
	case '<':
		tok = makeToken(token.LT, this.ch)
	case '>':
		tok = makeToken(token.GT, this.ch)
	case 0:
		tok.Value = ""
		tok.Type = token.EOF
	default:
		if isLetter(this.ch) {
			tok.Value = this.readIdentifier()
			tok.Type = token.LookupreservedWords(tok.Value)
			return tok
		} else if isDigit(this.ch) {

			tok.Type = token.INT
			tok.Value = this.readNumber()
			return tok
		} else {
			tok = makeToken(token.ILLEGAL, this.ch)
		}
	}
	this.readChar()
	return tok
}
func makeToken(tokenType token.TokenType, valueLiteral byte) token.Token {
	return token.Token{Type: tokenType, Value: string(valueLiteral)}
}
func isLetter(ch byte) bool {
	// etend language allowed chars for identifiers
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func (this *Lexer) readIdentifier() string {
	start := this.position

	for isLetter(this.ch) {
		this.readChar()
	}
	return this.input[start:this.position]
}
func (this *Lexer) skipWitheSpace() {
	for this.ch == ' ' || this.ch == '\t' || this.ch == '\n' || this.ch == '\r' {
		this.readChar()
	}
}
func isDigit(ch byte) bool {

	return ch >= '0' && ch <= '9'
}
func (this *Lexer) readNumber() string {
	start := this.position

	for isDigit(this.ch) {
		this.readChar()
	}
	return this.input[start:this.position]
}
