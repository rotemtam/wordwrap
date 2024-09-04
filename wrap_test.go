package wordwrap_test

import (
	"bufio"
	"github.com/rogpeppe/go-internal/testscript"
	"os"
	"rotemtam.com/wordwrap"
	"strconv"
	"strings"
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
		Cmds: map[string]func(ts *testscript.TestScript, neg bool, args []string){
			"maxline": maxline,
		},
	})
}

// maxline verifies that the longest line in args[0] is shorter than args[1] chars.
// Usage: maxline <path> <maxline>
func maxline(ts *testscript.TestScript, neg bool, args []string) {
	if len(args) != 2 {
		ts.Fatalf("usage: maxline <path> <maxline>")
	}
	l, ok := strconv.Atoi(args[1])
	if ok != nil {
		ts.Fatalf("usage: maxline <path> <maxline>")
	}
	scanner := bufio.NewScanner(
		strings.NewReader(
			ts.ReadFile(args[0]),
		),
	)
	tooLong := false
	for scanner.Scan() {
		if len(scanner.Text()) > l {
			tooLong = true
			break
		}
	}
	if tooLong && !neg {
		ts.Fatalf("line too long in %s", args[0])
	}
	if !tooLong && neg {
		ts.Fatalf("no line too long in %s", args[0])
	}
}
