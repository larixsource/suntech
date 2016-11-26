package st600

import (
	"github.com/larixsource/suntech/lexer"
	"github.com/larixsource/suntech/st"
)

type ExtDataReport struct {
	CommonReport
	Len              uint16
	Data             []byte
	Checksum         uint8
	DrivingHourMeter uint32
	BackupVolt       float32
	RealTime         bool
}

func (edr *ExtDataReport) Valid() bool {
	var sum byte
	for _, b := range edr.Data {
		sum += b
	}
	return sum == edr.Checksum
}

func parseUEXAscii(lex *lexer.Lexer, msg *Msg) {
	msg.Type = UEXReport

	uex := &ExtDataReport{}
	msg.UEX = uex
	uex.Hdr = UEXReport

	parseCommonAscii(lex, msg, &uex.CommonReport)
	if msg.ParsingError != nil {
		return
	}

	length, token, err := st.AsciiLen(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	uex.Len = length

	token, err = lex.NextFixed(int(length) + 1)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
	}
	uex.Data = token.Literal[:int(length)]

	chk, token, err := st.AsciiChecksum(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	uex.Checksum = chk

	hmeter, token, err := st.AsciiDrivingHourMeter(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	uex.DrivingHourMeter = hmeter

	backupVolt, token, err := st.AsciiBackupVolt(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	uex.BackupVolt = backupVolt

	realTime, token, err := st.AsciiBit(lex, true)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	uex.RealTime = realTime

	return
}
