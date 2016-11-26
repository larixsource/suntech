package st600

import (
	"github.com/larixsource/suntech/lexer"
	"github.com/larixsource/suntech/st"
)

type EmergencyReport struct {
	CommonReport
	EmgID            st.EmergencyType
	DrivingHourMeter uint32
	BackupVolt       float32
	RealTime         bool
	ADC              float32
}

func parseEMGAscii(lex *lexer.Lexer, msg *Msg) {
	msg.Type = EMGReport

	emg := &EmergencyReport{}
	msg.EMG = emg
	emg.Hdr = EMGReport

	parseCommonAscii(lex, msg, &emg.CommonReport)
	if msg.ParsingError != nil {
		return
	}

	emgID, token, err := st.AsciiEmgID(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	emg.EmgID = emgID

	hmeter, token, err := st.AsciiDrivingHourMeter(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	emg.DrivingHourMeter = hmeter

	backupVolt, token, err := st.AsciiBackupVolt(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	emg.BackupVolt = backupVolt

	realTime, token, err := st.AsciiBit(lex, false)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	emg.RealTime = realTime

	adc, token, err := st.AsciiADC(lex, true)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	emg.ADC = adc

	return
}
