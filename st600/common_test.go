package st600

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func equalCommon(t *testing.T, expected *CommonReport, actual *CommonReport) {
	require.NotNil(t, expected)
	require.NotNil(t, actual)
	assert.Equal(t, expected.Hdr, actual.Hdr)
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
}
