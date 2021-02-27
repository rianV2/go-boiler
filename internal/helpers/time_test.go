// +build unit

package helpers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetBodTime_ShouldReturnBeginningOfDayTime_WhenGivenUTCCurrentTime(t *testing.T) {

	// INIT
	currentTime := time.Now().UTC()
	y, m, d := currentTime.Date()
	loc := currentTime.Location()

	// CODE UNDER TEST
	bodTime := GetBodTime(currentTime)

	// EXPECTATION
	require.Equal(t, y, bodTime.Year())
	require.Equal(t, m, bodTime.Month())
	require.Equal(t, d, bodTime.Day())
	require.Equal(t, 0, bodTime.Hour())
	require.Equal(t, 0, bodTime.Minute())
	require.Equal(t, 0, bodTime.Second())
	require.Equal(t, loc, bodTime.Location())
}

func TestGetValidTimezone(t *testing.T) {

	expectations := []struct {
		input  string
		output *string
	}{
		{
			"Asia/Jakarta", StringPointer("Asia/Jakarta"),
		},
		{
			"Asia/Bandung", nil,
		},
		{
			"America/New_York", StringPointer("America/New_York"),
		},
		{
			"tz", nil,
		},
	}
	for _, exp := range expectations {
		ret := GetValidTimezone(exp.input)
		require.Equal(t, ret, exp.output)
	}
}
