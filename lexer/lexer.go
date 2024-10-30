package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.NewFromString(token.EQ, string(ch)+string(l.ch))
		} else {
			tok = token.NewFromChar(token.ASSIGN, l.ch)
		}
	case '+':
		tok = token.NewFromChar(token.PLUS, l.ch)
	case '-':
		tok = token.NewFromChar(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.NewFromString(token.NOTEQ, string(ch)+string(l.ch))
		} else {
			tok = token.NewFromChar(token.BANG, l.ch)
		}
	case '/':
		tok = token.NewFromChar(token.SLASH, l.ch)
	case '*':
		tok = token.NewFromChar(token.ASTERISK, l.ch)
	case '<':
		tok = token.NewFromChar(token.LT, l.ch)
	case '>':
		tok = token.NewFromChar(token.GT, l.ch)
	case ';':
		tok = token.NewFromChar(token.SEMICOLON, l.ch)
	case ',':
		tok = token.NewFromChar(token.COMMA, l.ch)
	case '(':
		tok = token.NewFromChar(token.LPAREN, l.ch)
	case ')':
		tok = token.NewFromChar(token.RPAREN, l.ch)
	case '{':
		tok = token.NewFromChar(token.LBRACE, l.ch)
	case '}':
		tok = token.NewFromChar(token.RBRACE, l.ch)
	case 0:
		tok = token.NewFromString(token.EOF, "")
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = token.NewFromChar(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && 'z' >= ch ||
		'A' <= ch && 'Z' >= ch ||
		ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && '9' >= ch
}
