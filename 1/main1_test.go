package main

import (
	"testing"
)

// тест функции convert_to_string
func TestConvertToString(t *testing.T) {
	got := convert_to_string()
	want := "4242423.14Golangtrue(1+2i)"

	if got != want {
		t.Errorf("convert_to_string() = %q; want %q", got, want)
	}
}

// тест функции convert_to_rune_slice
func TestConvertToRuneSlice(t *testing.T) {
	input := "Go"
	got := convert_to_rune_slice(input)
	want := []rune{'G', 'o'}

	if len(got) != len(want) {
		t.Fatalf("Длина неверна: got %d, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("convert_to_rune_slice()[%d] = %q; want %q", i, got[i], want[i])
		}
	}
}

// тест функции rune_with_salt
func TestRuneWithSalt(t *testing.T) {
	input := []rune("abcdef")
	got := rune_with_salt(input)
	want := "abcgo-2024def"

	if got != want {
		t.Errorf("rune_with_salt() = %q; want %q", got, want)
	}
}
