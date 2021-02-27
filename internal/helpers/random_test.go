// +build unit

package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {

	t.Run("ShouldReturnRandomStringWithGivenLength", func(t *testing.T) {
		t.Parallel()
		expected := 6
		actual := RandomString(expected)
		require.Equal(t, expected, len(actual))
	})

	t.Run("ShouldReturnUniqueString", func(t *testing.T) {
		t.Parallel()

		expectedUniqueStrNum := 1000
		generated := make(map[string]int, 0)
		for i := 0; i < expectedUniqueStrNum; i++ {
			generated[RandomString(6)] = 1
		}
		require.Equal(t, expectedUniqueStrNum, len(generated), "Should have generated 100 unique str")
	})

}
