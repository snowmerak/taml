package strcase_test

import (
	"testing"

	"github.com/snowmerak/taml/strcase"
)

func TestPascalToSnake(t *testing.T) {
	pascals := []string{"FooBar", "FooBarBaz", "FooBarBazQux"}
	snakes := []string{"foo_bar", "foo_bar_baz", "foo_bar_baz_qux"}
	for i, p := range pascals {
		s := strcase.PascalToSnake(p)
		if s != snakes[i] {
			t.Errorf("PascalToSnake(%q) = %q, want %q", p, s, snakes[i])
		}
	}
}

func TestSnakeToPascal(t *testing.T) {
	pascals := []string{"FooBar", "FooBarBaz", "FooBarBazQux"}
	snakes := []string{"foo_bar", "foo_bar_baz", "foo_bar_baz_qux"}
	for i, s := range snakes {
		p := strcase.SnakeToPascal(s)
		if p != pascals[i] {
			t.Errorf("SnakeToPascal(%q) = %q, want %q", s, p, pascals[i])
		}
	}
}

func FuzzFuzzCross(f *testing.F) {
	f.Fuzz(func(t *testing.T, target string) {
		p2s := strcase.PascalToSnake(target)
		s2p := strcase.SnakeToPascal(target)

		if p2s != strcase.PascalToSnake(s2p) {
			t.Errorf("PascalToSnake(%q) != PascalToSnake(%q)", p2s, s2p)
		}
		if s2p != strcase.SnakeToPascal(p2s) {
			t.Errorf("SnakeToPascal(%q) != SnakeToPascal(%q)", s2p, p2s)
		}
	})
}
