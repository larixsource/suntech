package st300

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/larixsource/suntech/st"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func equalEVT(t *testing.T, expected *EventReport, actual *EventReport) {
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
	assert.Equal(t, expected.EvtID, actual.EvtID)
	assert.Equal(t, expected.DrivingHourMeter, actual.DrivingHourMeter)
	assert.Equal(t, expected.BackupVolt, actual.BackupVolt)
	assert.Equal(t, expected.RealTime, actual.RealTime)
}

var testEVT = EventReport{
	Hdr:              EVTReport,
	DevID:            "100850000",
	Model:            st.ST300,
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
	EvtID:            st.Input1GroundEvt,
	DrivingHourMeter: 0,
	BackupVolt:       4.5,
	RealTime:         true,
}

func TestEVT300(t *testing.T) {
	// add to buf all type of EVT frames
	frameTemplate := "ST300EVT;100850000;01;010;20081017;07:41:56;00100;+37.478519;+126.886819;000.012;000.00;9;1;0;15.30;001100;%d;0;4.5;1\r"
	idList := []int{1, 2, 3, 4, 5, 6}
	var frames []string
	var buf bytes.Buffer
	for _, id := range idList {
		frame := fmt.Sprintf(frameTemplate, id)
		frames = append(frames, frame)
		buf.WriteString(frame)
	}

	// parse content of buf, save extracted msgs
	var msgs []*Msg
	p := ParseBytes(buf.Bytes(), ParserOpts{})
	for p.Next() {
		msgs = append(msgs, p.Msg())
	}
	assert.Nil(t, p.Error())

	// check every extracted msg
	expectedEVT := testEVT
	expectedEVTID := []st.EventType{st.Input1GroundEvt, st.Input1OpenEvt, st.Input2GroundEvt,
		st.Input2OpenEvt, st.Input3GroundEvt, st.Input3OpenEvt}
	assert.Len(t, msgs, len(idList))
	for i, msg := range msgs {
		expectedEVT.EvtID = expectedEVTID[i]

		assert.EqualValues(t, st.ST300, msg.Model)
		assert.Equal(t, []byte(frames[i]), msg.Frame)
		assert.Nil(t, msg.ParsingError)

		assert.Equal(t, msg.Type, EVTReport)
		equalEVT(t, &expectedEVT, msg.EVT)
	}
	assert.False(t, p.Next())
}
