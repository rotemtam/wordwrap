package wordwrap_test

import (
	"github.com/rogpeppe/go-internal/testscript"
	"os"
	"rotemtam.com/wordwrap"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"wordwrap": wordwrap.Run,
	}))
}

func TestScript(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata",
	})
}
