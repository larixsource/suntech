package lexer

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLexerGTEPS(t *testing.T) {
	// not a suntech frame :D
	frame := "+RESP:GTEPS,060100,135790246811220,,13500,00,1,1,4.3,92,70.0,121.354335,31.222073,20090214013254,0460,0000,18d8,6141,00,2000.0,20090214093254,11F0$"
	lexer := Lexer{
		Reader: strings.NewReader(frame),
	}

	// header
	token, err := lexer.Next(15, ',')
	require.Nil(t, err)
	assert.Equal(t, DataToken, token.Type)
	assert.Equal(t, []byte("+RESP:GTEPS,"), token.Literal)
	assert.False(t, token.OnlyDigits())

	// protocol version
	token, err = lexer.Next(7, ',')
	require.Nil(t, err)
	assert.Equal(t, DigitsToken, token.Type)
	assert.Equal(t, []byte("060100,"), token.Literal)
	assert.True(t, token.OnlyDigits())

	// dev id
	token, err = lexer.NextFixed(16)
	require.Nil(t, err)
	assert.Equal(t, DigitsToken, token.Type)
	assert.Equal(t, []byte("135790246811220,"), token.Literal)
	assert.True(t, token.OnlyDigits())

	// dev name
	token, err = lexer.Next(21, ',')
	require.Nil(t, err)
	assert.Equal(t, EmptyToken, token.Type)
	assert.Equal(t, []byte(","), token.Literal)
	assert.False(t, token.OnlyDigits())

	// ext. power
	token, err = lexer.Next(6, ',')
	require.Nil(t, err)
	assert.Equal(t, DigitsToken, token.Type)
	assert.Equal(t, []byte("13500,"), token.Literal)
	assert.True(t, token.OnlyDigits())

	// report type/id
	token, err = lexer.Next(3, ',')
	require.Nil(t, err)
	assert.Equal(t, BitsToken, token.Type)
	assert.Equal(t, []byte("00,"), token.Literal)
	assert.True(t, token.OnlyDigits())

	// number of positions
	token, err = lexer.Next(3, ',')
	require.Nil(t, err)
	assert.Equal(t, BitsToken, token.Type)
	assert.Equal(t, []byte("1,"), token.Literal)
	assert.True(t, token.OnlyDigits())

	// gps accuracy
	token, err = lexer.Next(3, ',')
	require.Nil(t, err)
	assert.Equal(t, BitsToken, token.Type)
	assert.Equal(t, []byte("1,"), token.Literal)
	assert.True(t, token.OnlyDigits())

	// speed
	token, err = lexer.Next(6, ',')
	require.Nil(t, err)
	assert.Equal(t, FloatToken, token.Type)
	assert.Equal(t, []byte("4.3,"), token.Literal)
	assert.False(t, token.OnlyDigits())

	// azimuth
	token, err = lexer.Next(4, ',')
	require.Nil(t, err)
	assert.Equal(t, DigitsToken, token.Type)
	assert.Equal(t, []byte("92,"), token.Literal)
	assert.True(t, token.OnlyDigits())

	// altitude
	token, err = lexer.Next(9, ',')
	require.Nil(t, err)
	assert.Equal(t, FloatToken, token.Type)
	assert.Equal(t, []byte("70.0,"), token.Literal)
	assert.False(t, token.OnlyDigits())

	// longitude
	token, err = lexer.Next(12, ',')
	require.Nil(t, err)
	assert.Equal(t, FloatToken, token.Type)
	assert.Equal(t, []byte("121.354335,"), token.Literal)
	assert.False(t, token.OnlyDigits())

	// latitude
	token, err = lexer.Next(11, ',')
	require.Nil(t, err)
	assert.Equal(t, FloatToken, token.Type)
	assert.Equal(t, []byte("31.222073,"), token.Literal)
	assert.False(t, token.OnlyDigits())

	// GPS timestamp
	token, err = lexer.NextFixed(15)
	require.Nil(t, err)
	assert.Equal(t, DigitsToken, token.Type)
	assert.Equal(t, []byte("20090214013254,"), token.Literal)
	assert.True(t, token.OnlyDigits())

	// MCC
	token, err = lexer.NextFixed(5)
	require.Nil(t, err)
	assert.Equal(t, DigitsToken, token.Type)
	assert.Equal(t, []byte("0460,"), token.Literal)
	assert.True(t, token.OnlyDigits())

	// MNC
	token, err = lexer.NextFixed(5)
	require.Nil(t, err)
	assert.Equal(t, BitsToken, token.Type)
	assert.Equal(t, []byte("0000,"), token.Literal)
	assert.True(t, token.OnlyDigits())

	// LAC
	token, err = lexer.NextFixed(5)
	require.Nil(t, err)
	assert.Equal(t, HexToken, token.Type)
	assert.Equal(t, []byte("18d8,"), token.Literal)
	assert.False(t, token.OnlyDigits())

	// Cell ID
	token, err = lexer.NextFixed(5)
	require.Nil(t, err)
	assert.Equal(t, DigitsToken, token.Type)
	assert.Equal(t, []byte("6141,"), token.Literal)
	assert.True(t, token.OnlyDigits())

	// reserved
	token, err = lexer.NextFixed(3)
	require.Nil(t, err)
	assert.Equal(t, BitsToken, token.Type)
	assert.Equal(t, []byte("00,"), token.Literal)
	assert.True(t, token.OnlyDigits())

	// mileage
	token, err = lexer.Next(10, ',')
	require.Nil(t, err)
	assert.Equal(t, FloatToken, token.Type)
	assert.Equal(t, []byte("2000.0,"), token.Literal)
	assert.False(t, token.OnlyDigits())

	// send time
	token, err = lexer.NextFixed(15)
	require.Nil(t, err)
	assert.Equal(t, DigitsToken, token.Type)
	assert.Equal(t, []byte("20090214093254,"), token.Literal)
	assert.True(t, token.OnlyDigits())

	// count number
	token, err = lexer.NextFixed(5)
	require.Nil(t, err)
	assert.Equal(t, HexToken, token.Type)
	assert.Equal(t, []byte("11F0$"), token.Literal)
	assert.False(t, token.OnlyDigits())

	// no more tokens
	token, err = lexer.NextFixed(1)
	assert.Equal(t, io.EOF, err)
}

func TestLexerFixedTooShort(t *testing.T) {
	frame := "too-short"
	lexer := Lexer{
		Reader: strings.NewReader(frame),
	}
	token, err := lexer.NextFixed(len(frame) + 1)
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, DataToken, token.Type)
	assert.Equal(t, []byte("too-short"), token.Literal)
}

func TestLexerFixedEmpty(t *testing.T) {
	frame := ""
	lexer := Lexer{
		Reader: strings.NewReader(frame),
	}
	token, err := lexer.NextFixed(1)
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, EmptyToken, token.Type)
	assert.Empty(t, token.Literal)
}
