package st300

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestST300CGF(t *testing.T) {
	cgf := ST300CGF{
		DevID:     "100850000",
		SwVer:     2,
		GeoID:     1,
		Active:    true,
		Latitude:  37.0,
		Longitude: 127.0,
		Radius:    50,
		In:        true,
		Out:       true,
	}
	assert.Equal(t, []byte("ST300CGF;100850000;02;1;1;37.000000;127.000000;50;1;1"), cgf.Command())
}

func TestST300CGFRes(t *testing.T) {
	frame := "ST300CGF;Res;100850000;010;1;1;+37.000000;+127.000000;50;1;1\r"
	p := ParseString(frame, ParserOpts{})
	assert.True(t, p.Next())
	assert.Nil(t, p.Error())
	msg := p.Msg()
	require.NotNil(t, msg)

	assert.Equal(t, CGFCmd, msg.Type)
	cgf := msg.CGF
	assert.NotNil(t, cgf)
	assert.True(t, cgf.Resp)
	assert.Equal(t, "100850000", cgf.DevID)
	assert.Equal(t, "010", cgf.SwVer)
	assert.Equal(t, 1, cgf.GeoID)
	assert.True(t, cgf.Active)
	assert.Equal(t, 37.0, cgf.Latitude)
	assert.Equal(t, 127.0, cgf.Longitude)
	assert.Equal(t, 50, cgf.Radius)
	assert.True(t, cgf.In)
	assert.True(t, cgf.Out)
}
