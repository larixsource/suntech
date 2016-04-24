package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexerEmptyState(t *testing.T) {
	dfa := emptyState
	assert.Equal(t, dfa.next('0'), bitsState)
	assert.Equal(t, dfa.next('1'), bitsState)
	assert.Equal(t, dfa.next('2'), digitsState)
	assert.Equal(t, dfa.next('+'), signState)
	assert.Equal(t, dfa.next('-'), signState)
	for _, c := range []byte("abcdefABCDEF") {
		assert.Equal(t, dfa.next(c), hexState)
	}
	assert.Equal(t, dfa.next('.'), dataState)
	assert.Equal(t, dfa.next('z'), dataState)
	assert.Equal(t, dfa.next('\r'), dataState)
}

func TestLexerBitsState(t *testing.T) {
	dfa := bitsState
	assert.Equal(t, dfa.next('0'), bitsState)
	assert.Equal(t, dfa.next('1'), bitsState)
	assert.Equal(t, dfa.next('2'), digitsState)
	assert.Equal(t, dfa.next('+'), dataState)
	assert.Equal(t, dfa.next('-'), dataState)
	for _, c := range []byte("abcdefABCDEF") {

		assert.Equal(t, dfa.next(c), hexState)
	}
	assert.Equal(t, dfa.next('.'), dotState)
	assert.Equal(t, dfa.next('z'), dataState)
	assert.Equal(t, dfa.next('\r'), dataState)
}

func TestLexerDigitsState(t *testing.T) {
	dfa := digitsState
	assert.Equal(t, dfa.next('0'), digitsState)
	assert.Equal(t, dfa.next('1'), digitsState)
	assert.Equal(t, dfa.next('2'), digitsState)
	assert.Equal(t, dfa.next('+'), dataState)
	assert.Equal(t, dfa.next('-'), dataState)
	for _, c := range []byte("abcdefABCDEF") {

		assert.Equal(t, dfa.next(c), hexState)
	}
	assert.Equal(t, dfa.next('.'), dotState)
	assert.Equal(t, dfa.next('z'), dataState)
	assert.Equal(t, dfa.next('\r'), dataState)
}

func TestLexerHexState(t *testing.T) {
	dfa := hexState
	assert.Equal(t, dfa.next('0'), hexState)
	assert.Equal(t, dfa.next('1'), hexState)
	assert.Equal(t, dfa.next('2'), hexState)
	assert.Equal(t, dfa.next('+'), dataState)
	assert.Equal(t, dfa.next('-'), dataState)
	for _, c := range []byte("abcdefABCDEF") {

		assert.Equal(t, dfa.next(c), hexState)
	}
	assert.Equal(t, dfa.next('.'), dataState)
	assert.Equal(t, dfa.next('z'), dataState)
	assert.Equal(t, dfa.next('\r'), dataState)
}

func TestLexerSignState(t *testing.T) {
	dfa := signState
	assert.Equal(t, dfa.next('0'), intState)
	assert.Equal(t, dfa.next('1'), intState)
	assert.Equal(t, dfa.next('2'), intState)
	assert.Equal(t, dfa.next('+'), dataState)
	assert.Equal(t, dfa.next('-'), dataState)
	for _, c := range []byte("abcdefABCDEF") {

		assert.Equal(t, dfa.next(c), dataState)
	}
	assert.Equal(t, dfa.next('.'), dataState)
	assert.Equal(t, dfa.next('z'), dataState)
	assert.Equal(t, dfa.next('\r'), dataState)
}

func TestLexerIntState(t *testing.T) {
	dfa := intState
	assert.Equal(t, dfa.next('0'), intState)
	assert.Equal(t, dfa.next('1'), intState)
	assert.Equal(t, dfa.next('2'), intState)
	assert.Equal(t, dfa.next('+'), dataState)
	assert.Equal(t, dfa.next('-'), dataState)
	for _, c := range []byte("abcdefABCDEF") {

		assert.Equal(t, dfa.next(c), dataState)
	}
	assert.Equal(t, dfa.next('.'), dotState)
	assert.Equal(t, dfa.next('z'), dataState)
	assert.Equal(t, dfa.next('\r'), dataState)
}

func TestLexerDotState(t *testing.T) {
	dfa := dotState
	assert.Equal(t, dfa.next('0'), floatState)
	assert.Equal(t, dfa.next('1'), floatState)
	assert.Equal(t, dfa.next('2'), floatState)
	assert.Equal(t, dfa.next('+'), dataState)
	assert.Equal(t, dfa.next('-'), dataState)
	for _, c := range []byte("abcdefABCDEF") {

		assert.Equal(t, dfa.next(c), dataState)
	}
	assert.Equal(t, dfa.next('.'), dataState)
	assert.Equal(t, dfa.next('z'), dataState)
	assert.Equal(t, dfa.next('\r'), dataState)
}

func TestLexerFloatState(t *testing.T) {
	dfa := floatState
	assert.Equal(t, dfa.next('0'), floatState)
	assert.Equal(t, dfa.next('1'), floatState)
	assert.Equal(t, dfa.next('2'), floatState)
	assert.Equal(t, dfa.next('+'), dataState)
	assert.Equal(t, dfa.next('-'), dataState)
	for _, c := range []byte("abcdefABCDEF") {

		assert.Equal(t, dfa.next(c), dataState)
	}
	assert.Equal(t, dfa.next('.'), dataState)
	assert.Equal(t, dfa.next('z'), dataState)
	assert.Equal(t, dfa.next('\r'), dataState)
}

func TestLexerDataState(t *testing.T) {
	dfa := dataState
	assert.Equal(t, dfa.next('0'), dataState)
	assert.Equal(t, dfa.next('1'), dataState)
	assert.Equal(t, dfa.next('2'), dataState)
	assert.Equal(t, dfa.next('+'), dataState)
	assert.Equal(t, dfa.next('-'), dataState)
	for _, c := range []byte("abcdefABCDEF") {

		assert.Equal(t, dfa.next(c), dataState)
	}
	assert.Equal(t, dfa.next('.'), dataState)
	assert.Equal(t, dfa.next('z'), dataState)
	assert.Equal(t, dfa.next('\r'), dataState)
}

func TestLexerUnknownState(t *testing.T) {
	dfa := dfaState(69)
	assert.Panics(t, func() {
		dfa.next('0')
	})
}
