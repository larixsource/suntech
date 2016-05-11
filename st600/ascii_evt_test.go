package st600

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
	equalCommon(t, &expected.CommonReport, &actual.CommonReport)
	assert.Equal(t, expected.EvtID, actual.EvtID)
	assert.Equal(t, expected.DrivingHourMeter, actual.DrivingHourMeter)
	assert.Equal(t, expected.BackupVolt, actual.BackupVolt)
	assert.Equal(t, expected.RealTime, actual.RealTime)
}

var testEVT = EventReport{
	CommonReport: CommonReport{
		Hdr:       EVTReport,
		DevID:     "205951725",
		Model:     st.ST600R,
		SwVer:     325,
		Timestamp: time.Date(2015, 12, 23, 13, 32, 30, 0, time.UTC),
		Cell: Cell{
			Type: Cell3GType,
			Cell3G: Cell3G{
				CellID:      "001cbf75",
				MCC:         "730",
				MNC:         "2",
				LAC:         "4e39",
				SignalLevel: 33,
			},
		},
		Latitude:   -33.363867,
		Longitude:  -070.670218,
		Speed:      0.122,
		Course:     0,
		Satellites: 5,
		GPSFixed:   true,
		Distance:   190269102,
		PowerVolt:  12.89,
		IO:         "000000",
	},
	EvtID:            st.Input1GroundEvt,
	DrivingHourMeter: 183230,
	BackupVolt:       4.5,
	RealTime:         false,
	ADC:              0,
}

func TestEVT600(t *testing.T) {
	// add to buf all type of EVT frames
	frameTemplate := "ST600EVT;205951725;20;325;20151223;13:32:30;001cbf75;730;2;4e39;33;-33.363867;-070.670218;000.122;000.00;5;1;190269102;12.89;000000;%d;183230;4.5;0;0.00\r"
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

		assert.EqualValues(t, st.ST600R, msg.Model)
		assert.Equal(t, []byte(frames[i]), msg.Frame)
		assert.Nil(t, msg.ParsingError)

		assert.Equal(t, msg.Type, EVTReport)
		equalEVT(t, &expectedEVT, msg.EVT)
	}
	assert.False(t, p.Next())
}
