package golexa

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimestamp_MarshalJSON(t *testing.T) {
	tt, _ := time.Parse(time.RFC3339Nano, "2015-05-13T12:34:56Z")
	ts := timestamp(tt)
	buf, _ := ts.MarshalJSON()

	assert.Equal(t, "2015-05-13T12:34:56Z", string(buf), "should marshal into correct string representation")
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	ts := timestamp(time.Time{})
	ts.UnmarshalJSON([]byte(`"2015-05-13T12:34:56Z"`))

	tt := time.Time(ts)

	year, month, day := tt.Date()
	hour, min, sec := tt.Clock()

	assert.EqualValues(t, 2015, year, "should have correct year")
	assert.EqualValues(t, 5, month, "should have correct month")
	assert.EqualValues(t, 13, day, "should have correct day")
	assert.Equal(t, "UTC", tt.Location().String(), "should have correct location")
	assert.EqualValues(t, 12, hour, "should have correct hour")
	assert.EqualValues(t, 34, min, "should have correct min")
	assert.EqualValues(t, 56, sec, "should have correct sec")
}
