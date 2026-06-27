package rev_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-rev"
	"github.com/gloo-foo/testable"
)

func ExampleRev_basic() {
	// echo "Hello World" | rev
	output, _ := testable.Test(command.Rev(), "Hello World\n")
	fmt.Print(output)
	// Output:
	// dlroW olleH
}
