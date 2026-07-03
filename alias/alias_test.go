package alias_test

import (
	"slices"
	"testing"

	"github.com/gloo-foo/testable"

	rev "github.com/gloo-foo/cmd-rev/alias"
)

// The alias package re-exports the constructor under an unprefixed name. A
// mis-wired re-export (Rev bound to the wrong function) compiles cleanly, so
// only behavior can prove the wiring. The test exercises the re-export and
// asserts the GNU rev output it must produce: each line reversed character by
// character (Unicode-aware), line order preserved.

func assertLines(t *testing.T, got, want []string) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestAlias_RevReversesEachLine(t *testing.T) {
	lines, err := testable.TestLines(rev.Rev(), "hello\nworld\n")
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"olleh", "dlrow"})
}

func TestAlias_RevIsUnicodeAware(t *testing.T) {
	// rev reverses by rune, not by byte: multi-byte characters stay intact.
	lines, err := testable.TestLines(rev.Rev(), "日本語\n")
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"語本日"})
}
