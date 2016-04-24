package lexer

import (
	"errors"
	"fmt"
	"io"
)

// TokenType indicates the type of token
type TokenType int

const (
	// EmptyToken is token with no content (just the delim byte)
	EmptyToken TokenType = iota

	// BitsToken is a token composed of 0 and 1
	BitsToken

	// DigitsToken is a token composed of digits (0-9)
	DigitsToken

	// HexToken is a token composed of hex digits (0-9a-fA-F)
	HexToken

	// FloatToken is a token composed of digits and, at most, one dot
	FloatToken

	// DataToken is token with arbitrary content
	DataToken
)

// Token is a single token recognized by the lexer
type Token struct {
	// Type indicates the type of the token (a broad category of the literal bytes)
	Type TokenType

	// Literal holds the token bytes, plus the separator at the end
	Literal []byte

	state dfaState
}

// OnlyDigits returns true if the token contains only digits (is a BitsToken or a DigitsToken)
func (t *Token) OnlyDigits() bool {
	return t.Type == BitsToken || t.Type == DigitsToken
}

// IsHex returns true if the token contains only hex digits (is a BitsToken, DigitsToken or HexToken)
func (t *Token) IsHex() bool {
	return t.Type == BitsToken || t.Type == DigitsToken || t.Type == HexToken
}

// EndsWith returns true if the last byte of the Literal is equals to the given delim byte.
func (t *Token) EndsWith(delim byte) bool {
	ll := len(t.Literal)
	if ll > 0 {
		return t.Literal[ll-1] == delim
	}
	return false
}

// WithoutSuffix returns Literal without its last byte
func (t *Token) WithoutSuffix() []byte {
	ll := len(t.Literal)
	if ll > 0 {
		return t.Literal[:ll-1]
	}
	return nil
}

// byte changes the Type field according to the byte c (make the internal dfa state to change).
func (t *Token) byte(c byte) {
	t.state = t.state.next(c)
	switch t.state {
	case emptyState:
		panic("impossibru!")
	case bitsState:
		t.Type = BitsToken
	case digitsState:
		t.Type = DigitsToken
	case hexState:
		t.Type = HexToken
	case signState:
		t.Type = DataToken
	case intState:
		t.Type = DataToken
	case dotState:
		t.Type = DataToken
	case floatState:
		t.Type = FloatToken
	case dataState:
		t.Type = DataToken
	default:
		panic(fmt.Errorf("unknown token type: %v", t.Type))
	}
}

// ErrTokenTooLong is returned by NextToken when the maximum length is reached without finding the delimiter byte
var ErrTokenTooLong = errors.New("token too long")

// Lexer is a very simple lexer, able to scan a reader using a delimiter byte and a maximum token length.
type Lexer struct {
	Reader io.Reader

	buf []byte
}

// Next scans the next token from the underlying reader, using a maximum length and a delimiter byte. If the maximum
// length is reached, an ErrTokenTooLong is returned.
// The delimiter byte is included in the Token literal and in the byte count.
func (l *Lexer) Next(max int, delim byte) (Token, error) {
	if max < 1 {
		return Token{}, fmt.Errorf("invalid max value, should be greater than 0")
	}
	if l.buf == nil {
		l.buf = make([]byte, 1)
	}
	t := Token{
		Type: EmptyToken,
	}
	for i := 0; i < max; i++ {
		_, err := io.ReadFull(l.Reader, l.buf)
		if err != nil {
			return t, err
		}
		c := l.buf[0]
		t.Literal = append(t.Literal, c)
		if c == delim {
			return t, nil
		}
		t.byte(c)
	}
	return t, ErrTokenTooLong
}

// NextFixed scans the next token from the underlying reader using a fixed length. If EOF is found before reading the
// token completely, an io.EOF is returned, along the resulting token (with a shorted literal obviously).
func (l *Lexer) NextFixed(length int) (Token, error) {
	if length < 1 {
		return Token{}, fmt.Errorf("invalid length value, should be greater than 0")
	}
	t := Token{
		Type:    EmptyToken,
		Literal: make([]byte, length),
	}
	_, err := io.ReadFull(l.Reader, t.Literal)
	switch err {
	case io.ErrUnexpectedEOF:
		for _, c := range t.Literal[:length-1] {
			t.byte(c)
		}
		t.Literal = t.Literal[:len(t.Literal)-1]
		return t, io.EOF
	case io.EOF:
		t.Literal = t.Literal[:len(t.Literal)-1]
		return t, err
	case nil:
	default:
		return t, err
	}

	for _, c := range t.Literal[:length-1] {
		t.byte(c)
	}
	return t, nil
}
