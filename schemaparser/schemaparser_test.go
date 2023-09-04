package schemaparser

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseTypes(t *testing.T) {
	schema, err := ioutil.ReadFile("../graph/schema.graphqls")
	require.NoError(t, err)

	t.Run("should return only the type names without interfaces", func(t *testing.T) {
		actualTypeNames, err := ParseTypes(schema)
		require.NoError(t, err)

		expectedTypeNames := []string{
			"Query",
			"Human",
			"Droid",
		}

		assert.Equal(t, expectedTypeNames, actualTypeNames)
	})
}

func TestCountEnumValues(t *testing.T) {
	schema, err := ioutil.ReadFile("../graph/schema.graphqls")
	require.NoError(t, err)

	t.Run("should return 11 for all enum values", func(t *testing.T) {
		actualEnumValues, err := CountEnumValues(schema)
		require.NoError(t, err)

		assert.Equal(t, 11, actualEnumValues)
	})
}
