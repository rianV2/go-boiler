package helpers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestEncodeByteToBase64(t *testing.T) {
	byteTest := []byte{72, 101, 108, 108, 111, 32, 84, 101, 115, 116}
	expectation := "SGVsbG8gVGVzdA=="

	result := EncodeByteToBase64(byteTest)
	require.Equal(t, expectation, result, "Does encoded correctly")
}

func TestDecodeBase64ToByte(t *testing.T) {
	base64Test := "SGVsbG8gVGVzdA=="
	expectation := []byte{72, 101, 108, 108, 111, 32, 84, 101, 115, 116}

	result := DecodeBase64ToByte(base64Test)
	require.Equal(t, expectation, result, "Does decoded correctly")
}

func TestSecondOrNotNilString(t *testing.T) {
	expectations := []struct {
		inputStr1 *string
		inputStr2 *string
		output    *string
	}{
		{
			nil,
			nil,
			nil,
		},
		{
			StringPointer("shortlyst"),
			nil,
			StringPointer("shortlyst"),
		},
		{
			nil,
			StringPointer("shortlyst lab"),
			StringPointer("shortlyst lab"),
		},
		{
			StringPointer("shortlyst"),
			StringPointer("shortlyst"),
			StringPointer("shortlyst"),
		},
		{
			StringPointer("shortlyst"),
			StringPointer("shortlyst lab"),
			StringPointer("shortlyst lab"),
		},
	}

	for _, expectation := range expectations {
		actual := SecondOrNotNilString(expectation.inputStr1, expectation.inputStr2)
		require.Equal(t, expectation.output, actual)
	}
}

func TestLatestTime(t *testing.T) {
	expectations := []struct {
		inputTime1 *time.Time
		inputTime2 *time.Time
		output     *time.Time
	}{
		{
			nil,
			nil,
			nil,
		},
		{
			TimePointer(time.Date(time.Now().Year(), time.November, 1, 0, 0, 0, 0, time.UTC)),
			nil,
			TimePointer(time.Date(time.Now().Year(), time.November, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			nil,
			TimePointer(time.Date(time.Now().Year(), time.October, 1, 0, 0, 0, 0, time.UTC)),
			TimePointer(time.Date(time.Now().Year(), time.October, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			TimePointer(time.Date(time.Now().Year(), time.November, 1, 0, 0, 0, 0, time.UTC)),
			TimePointer(time.Date(time.Now().Year(), time.November, 1, 0, 0, 0, 0, time.UTC)),
			TimePointer(time.Date(time.Now().Year(), time.November, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			TimePointer(time.Date(time.Now().Year(), time.November, 1, 0, 0, 0, 0, time.UTC)),
			TimePointer(time.Date(time.Now().Year(), time.October, 1, 0, 0, 0, 0, time.UTC)),
			TimePointer(time.Date(time.Now().Year(), time.November, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			TimePointer(time.Date(time.Now().Year(), time.October, 1, 0, 0, 0, 0, time.UTC)),
			TimePointer(time.Date(time.Now().Year(), time.November, 1, 0, 0, 0, 0, time.UTC)),
			TimePointer(time.Date(time.Now().Year(), time.November, 1, 0, 0, 0, 0, time.UTC)),
		},
	}

	for _, expectation := range expectations {
		actual := LatestTime(expectation.inputTime1, expectation.inputTime2)
		require.Equal(t, expectation.output, actual)
	}
}
