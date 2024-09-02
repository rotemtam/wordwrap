package wordwrap

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mitchellh/go-wordwrap"
	"os"
	"path/filepath"
	"strings"
)

var (
	path   = flag.String("path", "", "Path to search for .txt files (optional, defaults to cwd)")
	width  = flag.Uint("width", 40, "Maximum line width for formatting (default is 40 characters)")
	strict = flag.Bool("strict", false, "Enforce that no line exceeds the specified width")
)

// Run wordwrap and return the exit code.
func Run() int {
	if err := wrapFiles(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}

func wrapFiles() error {
	flag.Parse()

	// If no path is provided, default to current working directory
	if *path == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("unable to determine current working directory: %w", err)
		}
		*path = cwd
	}
	// Process files in the specified or default directory
	return filepath.Walk(*path, handleFile)
}

func handleFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	// Process only .txt files
	if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
		f, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}
		wrapped := wordwrap.WrapString(string(f), *width)
		if *strict {
			if err := validateLen(wrapped); err != nil {
				return fmt.Errorf("file %s: %w", path, err)
			}
		}
		return os.WriteFile(path, []byte(wrapped), os.ModePerm)
	}
	return nil
}

func validateLen(wrapped string) error {
	scanner := bufio.NewScanner(
		strings.NewReader(wrapped),
	)
	var l = 0
	for scanner.Scan() {
		l++
		if len(scanner.Text()) > int(*width) {
			return fmt.Errorf("line %d exceeds specified width %d", l, *width)
		}
	}
	return nil
}
