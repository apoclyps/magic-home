package cmd

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestPowerCmd(t *testing.T) {
	convey.Convey("test cmd", t, func() {
		cmd := powerCmd
		args := []string{}
		cmd.SetArgs(args)

		err := cmd.Execute()

		assert.NoError(t, err)
	})
}
