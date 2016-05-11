// Package st600 provides a parser for ST600 devices
package st600

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/larixsource/suntech/lexer"
	"github.com/larixsource/suntech/st"
)

var ErrUnknownHdr = errors.New("unknown HDR")

// ParserOpts holds configuration options that affect the behavior of the parser
type ParserOpts struct {
	// SkipUnknownFrames indicates to the parser if a frame with a unknown HDR should be consumed from the
	// underlying reader without stopping the parsing process.
	SkipUnknownFrames bool
}

// Parse returns a Parser to parse the content of a reader.
func Parse(r io.Reader, opts ParserOpts) *Parser {
	return &Parser{
		lex: &lexer.Lexer{
			Reader: r,
		},
		opts: opts,
	}
}

func ParseString(s string, opts ParserOpts) *Parser {
	return Parse(strings.NewReader(s), opts)
}

func ParseBytes(b []byte, opts ParserOpts) *Parser {
	return Parse(bytes.NewReader(b), opts)
}

// Parser is a ST300/ST340 parser
type Parser struct {
	lex  *lexer.Lexer
	opts ParserOpts

	last *Msg
	err  error
}

func (p *Parser) Next() bool {
	token, err := p.lex.NextFixed(1)
	if err != nil {
		if err != io.EOF {
			p.err = err
		}
		return false
	}

	switch token.Literal[0] {
	case st.STX:
		p.err = st.ErrZipUnsupported
		return false
	case 'S':
		p.last = p.parseAscii()
		return true
	default:
		p.err = fmt.Errorf("unexpected byte: %v", token.Literal[0])
		return false
	}
}

func (p *Parser) Msg() *Msg {
	return p.last
}

func (p *Parser) Error() error {
	return p.err
}

func (p *Parser) parseAscii() *Msg {
	msg := &Msg{}

	// get hdr tal (has to be T300CMD;)
	token, err := p.lex.NextFixed(8)
	if err != nil {
		msg.ParsingError = fmt.Errorf("error reading ascii hdr: %s", err)
	}
	hdr := asciiHdr(token)
	msg.Frame = append(msg.Frame, 'S')
	msg.Frame = append(msg.Frame, token.Literal...)

	switch hdr {
	case STTReport:
		parseSTTAscii(p.lex, msg)
	case EMGReport:
		parseEMGAscii(p.lex, msg)
	default:
		msg.ParsingError = ErrUnknownHdr
	}
	if msg.ParsingError == ErrUnknownHdr && p.opts.SkipUnknownFrames {
		token, err := p.lex.Next(1024, st.EndOfFrame)
		msg.Frame = append(msg.Frame, token.Literal...)
		if err != nil {
			msg.ParsingError = fmt.Errorf("error reading unknown frame: %+v", err)
		}
	}

	return msg
}

var (
	sttHdr = []byte("T600STT;")
	emgHdr = []byte("T600EMG;")
)

func asciiHdr(token lexer.Token) MsgType {
	if token.Type != lexer.DataToken {
		return UnknownMsg
	}
	switch {
	case bytes.Equal(token.Literal, sttHdr):
		return STTReport
	case bytes.Equal(token.Literal, emgHdr):
		return EMGReport
	default:
		return UnknownMsg
	}
}
