package alias

import (
	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-rev"
)

// Rev is a convenience constructor delegating to command.Rev.
func Rev(flags ...any) gloo.Command[[]byte, []byte] {
	return command.Rev(flags...)
}
