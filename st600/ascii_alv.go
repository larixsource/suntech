package st600

import (
	"github.com/larixsource/suntech/lexer"
	"github.com/larixsource/suntech/st"
)

type AliveReport struct {
	Hdr   MsgType
	DevID string
}

func parseALVAscii(lex *lexer.Lexer, msg *Msg) {
	msg.Type = ALVReport

	alv := &AliveReport{
		Hdr: ALVReport,
	}
	msg.ALV = alv

	devID, token, err := st.AsciiDevIDAtEnd(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alv.DevID = devID

	return
}
