package st600

import (
	"github.com/larixsource/suntech/lexer"
	"github.com/larixsource/suntech/st"
)

type AlertReport struct {
	CommonReport
	AltID            st.AlertType
	DrivingHourMeter uint32
	BackupVolt       float32
	RealTime         bool
	ADC              float32
}

func parseALTAscii(lex *lexer.Lexer, msg *Msg) {
	msg.Type = ALTReport

	alt := &AlertReport{}
	msg.ALT = alt
	alt.Hdr = ALTReport

	parseCommonAscii(lex, msg, &alt.CommonReport)
	if msg.ParsingError != nil {
		return
	}

	altID, token, err := st.AsciiAltID(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.AltID = altID

	hmeter, token, err := st.AsciiDrivingHourMeter(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.DrivingHourMeter = hmeter

	backupVolt, token, err := st.AsciiBackupVolt(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.BackupVolt = backupVolt

	realTime, token, err := st.AsciiBit(lex, false)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.RealTime = realTime

	adc, token, err := st.AsciiADC(lex, true)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.ADC = adc

	return
}
