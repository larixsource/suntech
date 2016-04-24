package st300

import (
	"time"

	"bytes"

	"github.com/larixsource/suntech/lexer"
	"github.com/larixsource/suntech/st"
)

type EventReport struct {
	Hdr              MsgType
	DevID            string
	Model            st.Model
	SwVer            uint16
	Timestamp        time.Time
	Cell             string
	Latitude         float32
	Longitude        float32
	Speed            float32
	Course           float32
	Satellites       uint8
	GPSFixed         bool
	Distance         uint32
	PowerVolt        float32
	IO               string
	EvtID            st.EventType
	DrivingHourMeter uint32
	BackupVolt       float32
	RealTime         bool
}

func parseEVTAscii(lex *lexer.Lexer, msg *Msg) {
	msg.Type = EVTReport

	evt := &EventReport{}
	msg.EVT = evt

	token, err := lex.Next(10, st.Separator)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	switch {
	case token.OnlyDigits() && len(token.WithoutSuffix()) == 9:
		evt.DevID = string(token.WithoutSuffix())
	case bytes.Equal(st.ResLiteral, token.WithoutSuffix()):
		msg.Type = UnknownMsg
		msg.ParsingError = ErrUnknownHdr
		return
	default:
		msg.ParsingError = st.ErrInvalidDevID
		return
	}

	model, token, err := st.AsciiModel(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	msg.Model = model
	evt.Model = model
	if !knownModel(model) {
		msg.ParsingError = st.ErrUnsupportedModel
		return
	}

	swVer, token, err := st.AsciiSwVer(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.SwVer = swVer

	ts, tokens, err := st.AsciiTimestamp(lex)
	msg.Frame = append(msg.Frame, tokens[0].Literal...)
	msg.Frame = append(msg.Frame, tokens[1].Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.Timestamp = ts

	cell, token, err := st.AsciiCell(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.Cell = cell

	lat, token, err := st.AsciiLat(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.Latitude = lat

	lon, token, err := st.AsciiLon(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.Longitude = lon

	speed, token, err := st.AsciiSpeed(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.Speed = speed

	course, token, err := st.AsciiCourse(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.Course = course

	satellites, token, err := st.AsciiSatellites(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.Satellites = satellites

	fix, token, err := st.AsciiFix(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.GPSFixed = fix

	distance, token, err := st.AsciiDistance(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.Distance = distance

	powerVolt, token, err := st.AsciiPowerVolt(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.PowerVolt = powerVolt

	ioStatus, token, err := asciiIO(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.IO = ioStatus

	evtID, token, err := st.AsciiEvtID(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.EvtID = evtID

	hmeter, token, err := st.AsciiDrivingHourMeter(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.DrivingHourMeter = hmeter

	backupVolt, token, err := st.AsciiBackupVolt(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.BackupVolt = backupVolt

	var unknownTail bool
	if model != st.ST300 && model != st.ST340 && model != st.ST340LC {
		unknownTail = true
	}

	realTime, token, err := st.AsciiMsgType(lex, !unknownTail)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	evt.RealTime = realTime

	if unknownTail {
		// TODO: estimate the maximum length of the unknown tail
		token, err := st.AsciiUnknownTail(lex, 64)
		msg.Frame = append(msg.Frame, token.Literal...)
		if err != nil {
			msg.ParsingError = err
			return
		}
	}

	return
}
