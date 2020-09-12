package cmd

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestDiscoverCmd(t *testing.T) {
	convey.Convey("test cmd", t, func() {
		cmd := discoverCmd
		args := []string{"discover"}
		cmd.SetArgs(args)

		err := cmd.Execute()

		assert.NoError(t, err)
	})
}
