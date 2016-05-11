package st600

import (
	"errors"

	"time"

	"github.com/larixsource/suntech/lexer"
	"github.com/larixsource/suntech/st"
)

var ErrInvalidIO = errors.New("invalid IO")

type CommonReport struct {
	Hdr        MsgType
	DevID      string
	Model      st.Model
	SwVer      uint16
	Timestamp  time.Time
	Cell       Cell
	Latitude   float32
	Longitude  float32
	Speed      float32
	Course     float32
	Satellites uint8
	GPSFixed   bool
	Distance   uint32
	PowerVolt  float32
	IO         string
}

func knownModel(model st.Model) bool {
	switch model {
	case st.ST600V, st.ST600R:
		return true
	default:
		return false
	}
}

func asciiIO(lex *lexer.Lexer) (ioStatus string, token lexer.Token, err error) {
	token, err = lex.Next(9, st.Separator)
	if err != nil {
		return
	}
	if token.Type != lexer.BitsToken {
		err = ErrInvalidIO
		return
	}
	if !token.EndsWith(st.Separator) {
		err = st.ErrSeparator
		return
	}
	ioStatus = string(token.WithoutSuffix())
	return
}

func parseCommonAscii(lex *lexer.Lexer, msg *Msg, cmn *CommonReport) {
	devID, token, err := st.AsciiDevID(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	cmn.DevID = devID

	model, token, err := st.AsciiModel(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	msg.Model = model
	cmn.Model = model
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
	cmn.SwVer = swVer

	ts, tokens, err := st.AsciiTimestamp(lex)
	msg.Frame = append(msg.Frame, tokens[0].Literal...)
	msg.Frame = append(msg.Frame, tokens[1].Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	cmn.Timestamp = ts

	cell, tokens, cellErr := asciiCell3G(lex)
	for _, t := range tokens {
		msg.Frame = append(msg.Frame, t.Literal...)
	}
	if cellErr != nil {
		msg.ParsingError = cellErr
		return
	}
	cmn.Cell = cell

	lat, token, err := st.AsciiLat(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	cmn.Latitude = lat

	lon, token, err := st.AsciiLon(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	cmn.Longitude = lon

	speed, token, err := st.AsciiSpeed(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	cmn.Speed = speed

	course, token, err := st.AsciiCourse(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	cmn.Course = course

	satellites, token, err := st.AsciiSatellites(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	cmn.Satellites = satellites

	fix, token, err := st.AsciiFix(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	cmn.GPSFixed = fix

	distance, token, err := st.AsciiDistance(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	cmn.Distance = distance

	powerVolt, token, err := st.AsciiPowerVolt(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	cmn.PowerVolt = powerVolt

	ioStatus, token, err := asciiIO(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	cmn.IO = ioStatus

	return
}
