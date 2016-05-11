package st600

import (
	"testing"
	"time"

	"github.com/larixsource/suntech/st"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func equalUEX(t *testing.T, expected *ExtDataReport, actual *ExtDataReport) {
	require.NotNil(t, expected)
	require.NotNil(t, actual)
	equalCommon(t, &expected.CommonReport, &actual.CommonReport)
	assert.Equal(t, expected.Len, actual.Len)
	assert.Equal(t, expected.Data, actual.Data)
	assert.Equal(t, expected.Checksum, actual.Checksum)
	assert.Equal(t, expected.DrivingHourMeter, actual.DrivingHourMeter)
	assert.Equal(t, expected.BackupVolt, actual.BackupVolt)
	assert.Equal(t, expected.RealTime, actual.RealTime)
}

func TestUEX600R(t *testing.T) {
	frame := "ST600UEX;205951719;20;325;20160202;18:32:54;001cbf75;730;2;4e39;42;-33.364026;-070.670234;000.056;184.17;7;1;4;9.14;100000;144;$FMS1,0,3,1.15,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0\r\n$FMS4,0,0,4265,3.3,4311,3.3,0,0,0,23,64,0,0,0\r\n$FMS8,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,@\r\n;81;000450;0.0;1\r"
	p := ParseString(frame, ParserOpts{})
	assert.True(t, p.Next())
	assert.Nil(t, p.Error())
	msg := p.Msg()
	require.NotNil(t, msg)

	expectedUEX := &ExtDataReport{
		CommonReport: CommonReport{
			Hdr:       UEXReport,
			DevID:     "205951719",
			Model:     st.ST600R,
			SwVer:     325,
			Timestamp: time.Date(2016, 2, 2, 18, 32, 54, 0, time.UTC),
			Cell: Cell{
				Type: Cell3GType,
				Cell3G: Cell3G{
					CellID:      "001cbf75",
					MCC:         "730",
					MNC:         "2",
					LAC:         "4e39",
					SignalLevel: 42,
				},
			},
			Latitude:   -33.364026,
			Longitude:  -070.670234,
			Speed:      0.056,
			Course:     184.17,
			Satellites: 7,
			GPSFixed:   true,
			Distance:   4,
			PowerVolt:  9.14,
			IO:         "100000",
		},
		Len:              144,
		Data:             []byte("$FMS1,0,3,1.15,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0\r\n$FMS4,0,0,4265,3.3,4311,3.3,0,0,0,23,64,0,0,0\r\n$FMS8,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,@\r\n"),
		Checksum:         0x81,
		DrivingHourMeter: 450,
		BackupVolt:       0,
		RealTime:         true,
	}

	assert.EqualValues(t, st.ST600R, msg.Model)
	assert.Equal(t, []byte(frame), msg.Frame)
	assert.Nil(t, msg.ParsingError)

	assert.Equal(t, msg.Type, UEXReport)
	equalUEX(t, expectedUEX, msg.UEX)
	assert.True(t, msg.UEX.Valid())

	assert.False(t, p.Next())
}

func TestUEX600R2(t *testing.T) {
	frame := "ST600UEX;205951719;20;325;20160202;19:02:45;001cbf72;730;2;4e39;42;-33.364049;-070.670220;000.063;000.00;7;1;21;9.14;100000;47;$FMS8,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,@\r\n;99;000479;0.0;0\r"
	p := ParseString(frame, ParserOpts{})
	assert.True(t, p.Next())
	assert.Nil(t, p.Error())
	msg := p.Msg()
	require.NotNil(t, msg)

	expectedUEX := &ExtDataReport{
		CommonReport: CommonReport{
			Hdr:       UEXReport,
			DevID:     "205951719",
			Model:     st.ST600R,
			SwVer:     325,
			Timestamp: time.Date(2016, 2, 2, 19, 2, 45, 0, time.UTC),
			Cell: Cell{
				Type: Cell3GType,
				Cell3G: Cell3G{
					CellID:      "001cbf72",
					MCC:         "730",
					MNC:         "2",
					LAC:         "4e39",
					SignalLevel: 42,
				},
			},
			Latitude:   -33.364049,
			Longitude:  -070.670220,
			Speed:      0.063,
			Course:     0,
			Satellites: 7,
			GPSFixed:   true,
			Distance:   21,
			PowerVolt:  9.14,
			IO:         "100000",
		},
		Len:              47,
		Data:             []byte("$FMS8,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,@\r\n"),
		Checksum:         0x99,
		DrivingHourMeter: 479,
		BackupVolt:       0,
		RealTime:         false,
	}

	assert.EqualValues(t, st.ST600R, msg.Model)
	assert.Equal(t, []byte(frame), msg.Frame)
	assert.Nil(t, msg.ParsingError)

	assert.Equal(t, msg.Type, UEXReport)
	equalUEX(t, expectedUEX, msg.UEX)
	assert.True(t, msg.UEX.Valid())

	assert.False(t, p.Next())
}
