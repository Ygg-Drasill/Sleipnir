package example

import "testing"

func TestExample(t *testing.T) {
	got := 1
	want := 1

	if got != want {
		t.Errorf("Got %q, expected %q", got, want)
	}
}
