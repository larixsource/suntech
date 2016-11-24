package st300

import (
	"bytes"
	"strconv"

	"github.com/larixsource/suntech/lexer"
	"github.com/larixsource/suntech/st"
)

type ST300CGF struct {
	DevID     string
	SwVer     uint16
	GeoID     int
	Active    bool
	Latitude  float32
	Longitude float32
	Radius    int
	In        bool
	Out       bool

	// Resp is true when this is a response
	Resp bool
}

func (gf *ST300CGF) Command() []byte {
	var buf bytes.Buffer
	buf.WriteString("ST300CGF;")
	buf.WriteString(gf.DevID)
	buf.WriteByte(';')
	buf.WriteString(strconv.FormatUint(uint64(gf.SwVer), 10))
	buf.WriteByte(';')
	buf.WriteString(strconv.Itoa(gf.GeoID))
	buf.WriteByte(';')
	if gf.Active {
		buf.WriteString("1")
	} else {
		buf.WriteString("0")
	}
	buf.WriteByte(';')
	buf.WriteString(strconv.FormatFloat(float64(gf.Latitude), 'f', 6, 32))
	buf.WriteByte(';')
	buf.WriteString(strconv.FormatFloat(float64(gf.Longitude), 'f', 6, 32))
	buf.WriteByte(';')
	buf.WriteString(strconv.Itoa(gf.Radius))
	buf.WriteByte(';')
	if gf.In {
		buf.WriteString("1")
	} else {
		buf.WriteString("0")
	}
	buf.WriteByte(';')
	if gf.Out {
		buf.WriteString("1")
	} else {
		buf.WriteString("0")
	}
	return buf.Bytes()
}

func parseCGF(lex *lexer.Lexer, msg *Msg) {
	msg.Type = CGFCmd

	cgf := &ST300CGF{}
	msg.CGF = cgf

	// Res and DevID, or just DevID
	isDevID, devID, token, err := st.AsciiDevIDOrRes(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	if isDevID {
		cgf.DevID = devID
	} else {
		cgf.Resp = true

		// get devID
		devID, token, err := st.AsciiDevID(lex)
		msg.Frame = append(msg.Frame, token.Literal...)
		if err != nil {
			msg.ParsingError = err
			return
		}
		cgf.DevID = devID
	}

	swVer, token, err := st.AsciiSwVer(lex)
	msg.Frame = append(msg.Frame, token.Literal...)
	if err != nil {
		msg.ParsingError = err
		return
	}
	cgf.SwVer = swVer

	// TODO
}
