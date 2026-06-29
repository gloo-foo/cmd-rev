package rev_test

import (
	"fmt"
	"os"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-rev"
)

// This example demonstrates reading from a file instead of inline input.
func ExampleRev_fromFile_basic() {
	// rev testdata/text.txt
	data, err := os.ReadFile("testdata/text.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "read testdata: %v\n", err)
		return
	}
	output, _ := testable.Test(command.Rev(), string(data))
	fmt.Print(output)
	// Output:
	// dlroW olleH
}
