package st300

import (
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
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
	assert.InEpsilon(t, expected.Latitude, actual.Latitude, epsilon)
	assert.InEpsilon(t, expected.Longitude, actual.Longitude, epsilon)
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
}

func TestSTT340LC(t *testing.T) {
	// not exactly the same of the spec, because there wasn't any with model 03
	frame := "ST300STT;100850000;03;010;20081017;07:41:56;00100;+37.478519;+126.886819;000.012;000.00;9;1;0;15.30;001100;1;0072;0;4.5;1\r"
	p := ParseString(frame, ParserOpts{})
	assert.True(t, p.Next())
	assert.Nil(t, p.Error())
	msg := p.Msg()
	require.NotNil(t, msg)
	spew.Dump(msg)

	expectedSTT := &StatusReport{
		Hdr:              STTReport,
		DevID:            "100850000",
		Model:            st.ST340LC,
		SwVer:            10,
		Timestamp:        time.Date(2008, 10, 17, 7, 41, 56, 0, time.UTC),
		Cell:             "00100",
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
	}

	assert.EqualValues(t, st.ST340LC, msg.Model)
	assert.Equal(t, []byte(frame), msg.Frame)
	assert.Nil(t, msg.ParsingError)

	assert.Equal(t, msg.Type, STTReport)
	equalSTT(t, expectedSTT, msg.StatusReport)

	assert.False(t, p.Next())
}

func TestSTT340Buff(t *testing.T) {
	frame := "ST300STT;205150043;02;529;20150716;19:33:30;6d6113;-32.644923;-071.424437;000.039;000.00;10;1;724692;12.89;000000;1;5069;001257;4.2;0\r"
	p := ParseString(frame, ParserOpts{})
	assert.True(t, p.Next())
	assert.Nil(t, p.Error())
	msg := p.Msg()
	require.NotNil(t, msg)

	spew.Dump(msg)

	expectedSTT := &StatusReport{
		Hdr:              STTReport,
		DevID:            "205150043",
		Model:            st.ST340,
		SwVer:            529,
		Timestamp:        time.Date(2015, 7, 16, 19, 33, 30, 0, time.UTC),
		Cell:             "6d6113",
		Latitude:         -32.644923,
		Longitude:        -71.424437,
		Speed:            0.039,
		Course:           0,
		Satellites:       10,
		GPSFixed:         true,
		Distance:         724692,
		PowerVolt:        12.89,
		IO:               "000000",
		Mode:             st.IdleMode,
		MsgNum:           5069,
		DrivingHourMeter: 1257,
		BackupVolt:       4.2,
		RealTime:         false,
	}

	assert.EqualValues(t, st.ST340, msg.Model)
	assert.Equal(t, []byte(frame), msg.Frame)
	assert.Nil(t, msg.ParsingError)

	assert.Equal(t, msg.Type, STTReport)
	equalSTT(t, expectedSTT, msg.StatusReport)

	assert.False(t, p.Next())
}

func TestSTT300K(t *testing.T) {
	frame := "ST300STT;600850802;12;999;20141212;09:47:21;04600;+37.479370;+126.888552;000.120;000.00;3;1;10660;12.25;000000;2;0036;002068;0.0;1;3.10;302799;0.00;215.86;01488BF1160000;1\r"
	p := ParseString(frame, ParserOpts{})
	assert.True(t, p.Next())
	assert.Nil(t, p.Error())
	msg := p.Msg()
	require.NotNil(t, msg)

	spew.Dump(msg)

	expectedSTT := &StatusReport{
		Hdr:              STTReport,
		DevID:            "600850802",
		Model:            st.ST300K,
		SwVer:            999,
		Timestamp:        time.Date(2014, 12, 12, 9, 47, 21, 0, time.UTC),
		Cell:             "04600",
		Latitude:         37.479370,
		Longitude:        126.888552,
		Speed:            0.12,
		Course:           0,
		Satellites:       3,
		GPSFixed:         true,
		Distance:         10660,
		PowerVolt:        12.25,
		IO:               "000000",
		Mode:             st.ActiveMode,
		MsgNum:           36,
		DrivingHourMeter: 2068,
		BackupVolt:       0.0,
		RealTime:         true,
	}

	assert.EqualValues(t, st.ST300K, msg.Model)
	assert.Equal(t, []byte(frame), msg.Frame)
	assert.Nil(t, msg.ParsingError)

	assert.Equal(t, msg.Type, STTReport)
	equalSTT(t, expectedSTT, msg.StatusReport)

	assert.False(t, p.Next())
}
