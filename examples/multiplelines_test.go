package rev_test

import (
	"fmt"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-rev"
)

func ExampleRev_multipleLines() {
	// echo with multiple lines piped to rev
	output, _ := testable.Test(command.Rev(), "First line\nSecond line\nThird line\n")
	fmt.Print(output)
	// Output:
	// enil tsriF
	// enil dnoceS
	// enil drihT
}
