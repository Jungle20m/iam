package business

import "testing"

func TestAdd(t *testing.T) {
	got := 10
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSub(t *testing.T) {
	got := 5
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
