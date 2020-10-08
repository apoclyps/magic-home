package lights_test

import (
	"testing"

	"github.com/apoclyps/magic-home/pkg/lights"
)

// TODO: update test to mock ReadFile and supply colour map
func TestValue(t *testing.T) {
	value, _ := lights.NewValue("blue")

	actual, err := value.GetColorByName()
	expected := lights.Blue()

	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("actual %d expected %d", actual, expected)
	}
}
