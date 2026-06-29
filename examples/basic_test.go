package rev_test

import (
	"fmt"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-rev"
)

func ExampleRev_basic() {
	// echo "Hello World" | rev
	output, _ := testable.Test(command.Rev(), "Hello World\n")
	fmt.Print(output)
	// Output:
	// dlroW olleH
}
