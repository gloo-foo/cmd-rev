package rev_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-rev"
	"github.com/gloo-foo/testable"
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
