package st300

import (
	"time"

	"github.com/larixsource/suntech/lexer"
	"github.com/larixsource/suntech/st"
)

type AlertReport struct {
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
	AltID            st.AlertType
	DrivingHourMeter uint32
	BackupVolt       float32
	RealTime         bool
}

func parseALTAscii(lex *lexer.Lexer, msg *Msg) {
	msg.Type = ALTReport

	alt := &AlertReport{}
	msg.ALT = alt

	devID, token, err := st.AsciiDevID(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.DevID = devID

	model, token, err := st.AsciiModel(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	msg.Model = model
	alt.Model = model
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
	alt.SwVer = swVer

	ts, tokens, err := st.AsciiTimestamp(lex)
	msg.Frame = append(msg.Frame, tokens[0].Literal...)
	msg.Frame = append(msg.Frame, tokens[1].Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.Timestamp = ts

	cell, token, err := st.AsciiCell(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.Cell = cell

	lat, token, err := st.AsciiLat(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.Latitude = lat

	lon, token, err := st.AsciiLon(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.Longitude = lon

	speed, token, err := st.AsciiSpeed(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.Speed = speed

	course, token, err := st.AsciiCourse(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.Course = course

	satellites, token, err := st.AsciiSatellites(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.Satellites = satellites

	fix, token, err := st.AsciiFix(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.GPSFixed = fix

	distance, token, err := st.AsciiDistance(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.Distance = distance

	powerVolt, token, err := st.AsciiPowerVolt(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.PowerVolt = powerVolt

	ioStatus, token, err := asciiIO(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.IO = ioStatus

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

	var unknownTail bool
	if model != st.ST300 && model != st.ST340 && model != st.ST340LC {
		unknownTail = true
	}

	realTime, token, err := st.AsciiBit(lex, !unknownTail)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	alt.RealTime = realTime

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
