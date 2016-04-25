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

func equalALT(t *testing.T, expected *AlertReport, actual *AlertReport) {
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
	assert.Equal(t, expected.AltID, actual.AltID)
	assert.Equal(t, expected.DrivingHourMeter, actual.DrivingHourMeter)
	assert.Equal(t, expected.BackupVolt, actual.BackupVolt)
	assert.Equal(t, expected.RealTime, actual.RealTime)
}

var testALT = AlertReport{
	Hdr:              EMGReport,
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
	AltID:            st.StartOverSpeedAlt,
	DrivingHourMeter: 0,
	BackupVolt:       4.5,
	RealTime:         true,
}

func TestALT300(t *testing.T) {
	// add to buf all type of ALT frames
	frameTemplate := "ST300ALT;100850000;01;010;20081017;07:41:56;00100;+37.478519;+126.886819;000.012;000.00;9;1;0;15.30;001100;%d;0;4.5;1\r"
	idList := []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 13, 14, 15, 16, 18, 19, 22, 23, 24, 25, 26, 27, 28, 33, 34, 40, 41,
		44, 45, 46, 47, 48, 49, 50, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 73}
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
	expectedALT := testALT
	expectedALTID := []st.AlertType{st.StartOverSpeedAlt, st.StopOverSpeedAlt, st.DisconnectedGPSAntennaAlt,
		st.ReconnectedGPSAntennaAlt, st.ExitedGeoFenceAlt, st.EnteredGeoFenceAlt, st.ShortedGPSAntennaAlt,
		st.EnterDeepSleepModeAlt, st.ExitDeepSleepModeAlt, st.BackupBatteryErrorAlt, st.BatteryLowLevelAlt,
		st.ShockedAlt, st.CollisionAlt, st.DeviatedFromRouteAlt, st.EnteredIntoRouteAlt, st.EngineExceedSpeedAlt,
		st.EngineVehicleSpeedAlt, st.EngineCoolantTempAlt, st.EngineOilPressureAlt, st.EngineRPMAlt,
		st.EngineHardBrakeAlt, st.EngineErrCodeAlt, st.IgnitionOnAlt, st.IgnitionOffAlt, st.ConnectedToMainPowerAlt,
		st.DisconnectedFromMainPowerAlt, st.ConnectedToBackupBatteryAlt, st.DisconnectedToBackupBatteryAlt,
		st.FastAccelerationFromDPAAlt, st.FastBrakingFromDPAAlt, st.SharpTurnFromDPAAlt, st.OverSpeedFromDPAAlt,
		st.JammingDetectedAlt, st.InsertedIButtonAlt, st.RemovedIButtonAlt, st.DriveLessThanPredefinedTimeAlt,
		st.StoppedMoreThanPredefinedTimeAlt, st.DeadCenterAlt, st.OverRPMAlt, st.CompletedAutoRPMCalibrationAlt,
		st.CompletedAutoOdometerCalibrationAlt, st.CompletedAutoOdometerCalibrationDualGearSystemAlt,
		st.StopLimitAtIgnitionONAlt, st.MovingAfterStopLimitAtIgnitionONAlt, st.RapidFuelReductionAlt}

	assert.Len(t, expectedALTID, len(idList))
	assert.Len(t, msgs, len(idList))
	for i, msg := range msgs {
		expectedALT.AltID = expectedALTID[i]

		assert.EqualValues(t, st.ST300, msg.Model)
		assert.Equal(t, []byte(frames[i]), msg.Frame)
		assert.Nil(t, msg.ParsingError)

		assert.Equal(t, msg.Type, ALTReport)
		equalALT(t, &expectedALT, msg.ALT)
	}
	assert.False(t, p.Next())
}
