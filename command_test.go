package command_test

import (
	"fmt"
	"strings"
	"testing"

	command "github.com/gloo-foo/cmd-rev"
	"github.com/gloo-foo/testable"
)

func TestRev_SingleLine(t *testing.T) {
	lines, err := testable.TestLines(command.Rev(), "hello\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "olleh" {
		t.Fatalf("got %q, want [olleh]", lines)
	}
}

func TestRev_MultipleLines(t *testing.T) {
	lines, err := testable.TestLines(command.Rev(), "hello\nworld\ntest\n")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"olleh", "dlrow", "tset"}
	if len(lines) != len(want) {
		t.Fatalf("got %d lines, want %d", len(lines), len(want))
	}
	for i, w := range want {
		if lines[i] != w {
			t.Errorf("line %d: got %q, want %q", i, lines[i], w)
		}
	}
}

func TestRev_EmptyInput(t *testing.T) {
	lines, err := testable.TestLines(command.Rev(), "")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 0 {
		t.Fatalf("got %q, want empty", lines)
	}
}

func TestRev_Unicode_Japanese(t *testing.T) {
	lines, err := testable.TestLines(command.Rev(), "日本語\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "語本日" {
		t.Fatalf("got %q, want [語本日]", lines)
	}
}

func TestRev_Unicode_Emoji(t *testing.T) {
	lines, err := testable.TestLines(command.Rev(), "abc🌍def👋ghi\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "ihg👋fed🌍cba" {
		t.Fatalf("got %q, want [ihg👋fed🌍cba]", lines)
	}
}

func TestRev_NoError(t *testing.T) {
	_, err := testable.TestLines(command.Rev(), "anything\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRev_Palindrome(t *testing.T) {
	lines, err := testable.TestLines(command.Rev(), "racecar\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "racecar" {
		t.Fatalf("got %q, want [racecar]", lines)
	}
}

func TestRev_MixedContent(t *testing.T) {
	lines, err := testable.TestLines(command.Rev(), "abc\n\n日本語\n123\n")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"cba", "", "語本日", "321"}
	if len(lines) != len(want) {
		t.Fatalf("got %d lines, want %d", len(lines), len(want))
	}
	for i, w := range want {
		if lines[i] != w {
			t.Errorf("line %d: got %q, want %q", i, lines[i], w)
		}
	}
}

func TestRev_Spaces(t *testing.T) {
	lines, err := testable.TestLines(command.Rev(), "abc 123\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "321 cba" {
		t.Fatalf("got %q, want [321 cba]", lines)
	}
}

func TestRev_SingleChar(t *testing.T) {
	lines, err := testable.TestLines(command.Rev(), "x\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "x" {
		t.Fatalf("got %q, want [x]", lines)
	}
}

func TestRev_LongLine(t *testing.T) {
	input := "start" + strings.Repeat("x", 5000) + "end\n"
	lines, err := testable.TestLines(command.Rev(), input)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 {
		t.Fatalf("got %d lines, want 1", len(lines))
	}
	if !strings.HasPrefix(lines[0], "dne") {
		t.Errorf("expected prefix 'dne', got %q", lines[0][:10])
	}
	if !strings.HasSuffix(lines[0], "trats") {
		t.Errorf("expected suffix 'trats', got %q", lines[0][len(lines[0])-10:])
	}
}

func ExampleRev() {
	lines, _ := testable.TestLines(command.Rev(), "hello\nworld\n")
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// olleh
	// dlrow
}

func ExampleRev_unicode() {
	lines, _ := testable.TestLines(command.Rev(), "日本語\n")
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// 語本日
}
