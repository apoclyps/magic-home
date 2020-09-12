package cmd

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	convey.Convey("test cmd", t, func() {
		cmd := rootCmd
		args := []string{}
		cmd.SetArgs(args)

		err := cmd.Execute()

		assert.NoError(t, err)
	})
}
