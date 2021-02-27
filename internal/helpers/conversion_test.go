// +build unit

package helpers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestVal(t *testing.T) {
	var nilString *string

	expectations := []struct {
		input  interface{}
		output interface{}
	}{
		{
			"str", "str",
		},
		{
			StringPointer("str"), "str",
		},
		{
			nil, nil,
		},
		{
			123, 123,
		},
		{
			IntPointer(123), 123,
		},
		{
			nilString, nil,
		},
	}
	for _, exp := range expectations {
		ret := Val(exp.input)
		require.Equal(t, ret, exp.output)
	}

}

func TestValStr(t *testing.T) {
	testCases := map[string]struct {
		Input  *string
		Output string
	}{
		"ShouldReturnEmptyString_WhenInputIsNil": {
			Input:  nil,
			Output: "",
		},
		"ShouldReturnEmptyString_WhenInputIsEmptyString": {
			Input:  StringPointer(""),
			Output: "",
		},
		"ShouldReturnString_WhenInputIsNotNil": {
			Input:  StringPointer("SomeValue"),
			Output: "SomeValue",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := ValStr(testCase.Input)
			require.Equal(t, testCase.Output, actual)
		})
	}
}

func TestValTimeUnix(t *testing.T) {
	now := time.Now()
	testCases := map[string]struct {
		Input  *time.Time
		Output int64
	}{
		"ShouldReturnZero_WhenInputIsNil": {
			Input:  nil,
			Output: int64(0),
		},
		"ShouldReturnUnixtimestamp_WhenInputIsNotNil": {
			Input:  &now,
			Output: now.Unix(),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := ValTimeUnix(testCase.Input)
			require.Equal(t, testCase.Output, actual)
		})
	}
}

func TestInterfaceToString(t *testing.T) {
	expectations := []struct {
		input  interface{}
		output *string
	}{
		{
			"affiliate marketing staff",
			StringPointer("affiliate marketing staff"),
		},
		{
			0.111,
			StringPointer("0.111"),
		},
		{
			11111,
			StringPointer("11111"),
		},
	}

	for _, expectation := range expectations {
		actual := InterfaceToString(expectation.input)
		require.Equal(t, expectation.output, actual)
	}
}

func TestEqualValueStr(t *testing.T) {
	expectations := []struct {
		Param1         *string
		Param2         *string
		ExpectedOutput bool
	}{
		{
			Param1:         nil,
			Param2:         nil,
			ExpectedOutput: true,
		},
		{
			Param1:         StringPointer(""),
			Param2:         StringPointer(""),
			ExpectedOutput: true,
		},
		{
			Param1:         StringPointer("text1"),
			Param2:         StringPointer("text2"),
			ExpectedOutput: false,
		},
		{
			Param1:         nil,
			Param2:         StringPointer("text2"),
			ExpectedOutput: false,
		},
	}

	for _, expectation := range expectations {
		actual := EqualValueStr(expectation.Param1, expectation.Param2)
		require.Equal(t, expectation.ExpectedOutput, actual)
	}
}
