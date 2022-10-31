package g2engineconfigurationjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, err error) {
	if err != nil {
		assert.FailNow(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestBuildSimpleSystemConfigurationJson(test *testing.T) {
	actual, err := BuildSimpleSystemConfigurationJson("postgresql://postgres:postgres@$10.0.0.1:5432/G2")
	testError(test, err)
	test.Log("Actual:", actual)
}

func TestBuildSimpleSystemConfigurationJsonFailure(test *testing.T) {
	actual, err := BuildSimpleSystemConfigurationJson("")
	if err == nil {
		assert.FailNow(test, "There should be an error.")
	}
	test.Log("Actual:", actual)
}
