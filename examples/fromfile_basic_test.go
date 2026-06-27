package rev_test

import (
	"fmt"
	"os"

	command "github.com/gloo-foo/cmd-rev"
	"github.com/gloo-foo/testable"
)

// This example demonstrates reading from a file instead of inline input.
func ExampleRev_fromFile_basic() {
	// rev testdata/text.txt
	data, err := os.ReadFile("testdata/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "read testdata: %v\n", err)
		return
	}
	output, _ := testable.Test(command.Rev(), string(data))
	fmt.Print(output)
	// Output:
	// dlroW olleH
}
