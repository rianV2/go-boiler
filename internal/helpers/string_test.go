// +build unit

package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStrLimit(t *testing.T) {
	expectations := []struct {
		input  *string
		limit  int
		output *string
	}{
		{
			nil,
			3,
			nil,
		},
		{
			StringPointer("Too long char"),
			3,
			StringPointer("Too"),
		},
		{
			StringPointer("Too long char"),
			13,
			StringPointer("Too long char"),
		},
		{
			StringPointer("Too long char"),
			30,
			StringPointer("Too long char"),
		},
		{
			StringPointer("Too long char"),
			0,
			StringPointer(""),
		},
		{
			StringPointer("Too long char"),
			-1,
			StringPointer(""),
		},
	}

	for _, expectation := range expectations {
		actual := StrLimit(expectation.input, expectation.limit)
		require.Equal(t, expectation.output, actual)
	}
}
