package st600

import (
	"github.com/larixsource/suntech/lexer"
	"github.com/larixsource/suntech/st"
)

type StatusReport struct {
	CommonReport
	Mode             st.ModeType
	MsgNum           uint16
	DrivingHourMeter uint32
	BackupVolt       float32
	RealTime         bool
	ADC              float32
}

func parseSTTAscii(lex *lexer.Lexer, msg *Msg) {
	msg.Type = STTReport

	stt := &StatusReport{}
	msg.STT = stt
	stt.Hdr = STTReport

	parseCommonAscii(lex, msg, &stt.CommonReport)
	if msg.ParsingError != nil {
		return
	}

	mode, token, err := st.AsciiMode(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	stt.Mode = mode

	msgNum, token, err := st.AsciiMsgNum(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	stt.MsgNum = msgNum

	hmeter, token, err := st.AsciiDrivingHourMeter(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	stt.DrivingHourMeter = hmeter

	backupVolt, token, err := st.AsciiBackupVolt(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	stt.BackupVolt = backupVolt

	realTime, token, err := st.AsciiBit(lex, false)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	stt.RealTime = realTime

	adc, token, err := st.AsciiADC(lex, true)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	stt.ADC = adc

	return
}
