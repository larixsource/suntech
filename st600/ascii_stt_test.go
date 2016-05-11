package st600

import (
	"testing"
	"time"

	"github.com/larixsource/suntech/st"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const epsilon = 0.00001

func equalSTT(t *testing.T, expected *StatusReport, actual *StatusReport) {
	require.NotNil(t, expected)
	require.NotNil(t, actual)
	assert.Equal(t, expected.DevID, actual.DevID)
	assert.Equal(t, expected.Model, actual.Model)
	assert.Equal(t, expected.SwVer, actual.SwVer)
	assert.Equal(t, expected.Timestamp, actual.Timestamp)
	assert.Equal(t, expected.Cell, actual.Cell)
	if expected.Latitude != 0 && actual.Latitude != 0 {
		assert.InEpsilon(t, expected.Latitude, actual.Latitude, epsilon)
	}
	if expected.Longitude != 0 && actual.Longitude != 0 {
		assert.InEpsilon(t, expected.Longitude, actual.Longitude, epsilon)
	}
	assert.Equal(t, expected.Speed, actual.Speed)
	assert.Equal(t, expected.Course, actual.Course)
	assert.Equal(t, expected.Satellites, actual.Satellites)
	assert.Equal(t, expected.GPSFixed, actual.GPSFixed)
	assert.Equal(t, expected.Distance, actual.Distance)
	assert.Equal(t, expected.PowerVolt, actual.PowerVolt)
	assert.Equal(t, expected.IO, actual.IO)
	assert.Equal(t, expected.Mode, actual.Mode)
	assert.Equal(t, expected.MsgNum, actual.MsgNum)
	assert.Equal(t, expected.DrivingHourMeter, actual.DrivingHourMeter)
	assert.Equal(t, expected.BackupVolt, actual.BackupVolt)
	assert.Equal(t, expected.RealTime, actual.RealTime)
	if expected.ADC != 0 && actual.ADC != 0 {
		assert.InEpsilon(t, expected.ADC, actual.ADC, epsilon)
	}
}

func TestSTT600R(t *testing.T) {
	// not exactly the same of the spec, because there wasn't any with model 03
	frame := "ST600STT;100850000;20;010;20081017;07:41:56;00100;+37.478519;+126.886819;000.012;000.00;9;1;0;15.30;001100;1;0072;0;4.5;1;12.35\r"
	p := ParseString(frame, ParserOpts{})
	assert.True(t, p.Next())
	assert.Nil(t, p.Error())
	msg := p.Msg()
	require.NotNil(t, msg)

	expectedSTT := &StatusReport{
		Hdr:       STTReport,
		DevID:     "100850000",
		Model:     st.ST600R,
		SwVer:     10,
		Timestamp: time.Date(2008, 10, 17, 7, 41, 56, 0, time.UTC),
		Cell: Cell{
			Type:   Cell2GType,
			Cell2G: "00100",
		},
		Latitude:         37.478519,
		Longitude:        126.886819,
		Speed:            0.012,
		Course:           0,
		Satellites:       9,
		GPSFixed:         true,
		Distance:         0,
		PowerVolt:        15.3,
		IO:               "001100",
		Mode:             st.IdleMode,
		MsgNum:           72,
		DrivingHourMeter: 0,
		BackupVolt:       4.5,
		RealTime:         true,
		ADC:              12.35,
	}

	assert.EqualValues(t, st.ST600R, msg.Model)
	assert.Equal(t, []byte(frame), msg.Frame)
	assert.Nil(t, msg.ParsingError)

	assert.Equal(t, msg.Type, STTReport)
	equalSTT(t, expectedSTT, msg.STT)

	assert.False(t, p.Next())
}

func TestSTT600R3G(t *testing.T) {
	// real frame
	frame := "ST600STT;205951725;20;325;20151224;10:10:44;001cbf72;730;2;4e39;47;-33.363627;-070.670525;000.056;000.00;6;1;190269159;12.79;000000;1;0053;183231;0.0;0;0.00\r"
	p := ParseString(frame, ParserOpts{})
	assert.True(t, p.Next())
	assert.Nil(t, p.Error())
	msg := p.Msg()
	require.NotNil(t, msg)

	expectedSTT := &StatusReport{
		Hdr:       STTReport,
		DevID:     "205951725",
		Model:     st.ST600R,
		SwVer:     325,
		Timestamp: time.Date(2015, 12, 24, 10, 10, 44, 0, time.UTC),
		Cell: Cell{
			Type: Cell3GType,
			Cell3G: Cell3G{
				CellID:      "001cbf72",
				MCC:         "730",
				MNC:         "2",
				LAC:         "4e39",
				SignalLevel: 47,
			},
		},
		Latitude:         -33.363627,
		Longitude:        -070.670525,
		Speed:            0.056,
		Course:           0,
		Satellites:       6,
		GPSFixed:         true,
		Distance:         190269159,
		PowerVolt:        12.79,
		IO:               "000000",
		Mode:             st.IdleMode,
		MsgNum:           53,
		DrivingHourMeter: 183231,
		BackupVolt:       0,
		RealTime:         false,
		ADC:              0,
	}

	assert.EqualValues(t, st.ST600R, msg.Model)
	assert.Equal(t, []byte(frame), msg.Frame)
	assert.Nil(t, msg.ParsingError)

	assert.Equal(t, msg.Type, STTReport)
	equalSTT(t, expectedSTT, msg.STT)

	assert.False(t, p.Next())
}

func TestSTT600VBuff(t *testing.T) {
	frame := "ST600STT;205150043;21;529;20150716;19:33:30;6d6113;-32.644923;-071.424437;000.039;000.00;10;1;724692;12.89;00110000;1;5069;001257;4.2;0;12.35\r"
	p := ParseString(frame, ParserOpts{})
	assert.True(t, p.Next())
	assert.Nil(t, p.Error())
	msg := p.Msg()
	require.NotNil(t, msg)

	expectedSTT := &StatusReport{
		Hdr:       STTReport,
		DevID:     "205150043",
		Model:     st.ST600V,
		SwVer:     529,
		Timestamp: time.Date(2015, 7, 16, 19, 33, 30, 0, time.UTC),
		Cell: Cell{
			Type:   Cell2GType,
			Cell2G: "6d6113",
		},
		Latitude:         -32.644923,
		Longitude:        -71.424437,
		Speed:            0.039,
		Course:           0,
		Satellites:       10,
		GPSFixed:         true,
		Distance:         724692,
		PowerVolt:        12.89,
		IO:               "00110000",
		Mode:             st.IdleMode,
		MsgNum:           5069,
		DrivingHourMeter: 1257,
		BackupVolt:       4.2,
		RealTime:         false,
		ADC:              12.35,
	}

	assert.EqualValues(t, st.ST600V, msg.Model)
	assert.Equal(t, []byte(frame), msg.Frame)
	assert.Nil(t, msg.ParsingError)

	assert.Equal(t, msg.Type, STTReport)
	equalSTT(t, expectedSTT, msg.STT)

	assert.False(t, p.Next())
}
