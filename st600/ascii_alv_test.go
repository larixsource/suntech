package st600

import (
	"testing"

	"github.com/larixsource/suntech/st"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestALV600R(t *testing.T) {
	frame := "ST600ALV;600850777\r"
	p := ParseString(frame, ParserOpts{})
	assert.True(t, p.Next())
	assert.Nil(t, p.Error())
	msg := p.Msg()
	require.NotNil(t, msg)

	assert.EqualValues(t, st.UnknownModel, msg.Model)
	assert.Equal(t, []byte(frame), msg.Frame)
	assert.Nil(t, msg.ParsingError)

	assert.Equal(t, ALVReport, msg.Type)
	assert.Equal(t, ALVReport, msg.ALV.Hdr)
	assert.Equal(t, "600850777", msg.ALV.DevID)

	assert.False(t, p.Next())
}
