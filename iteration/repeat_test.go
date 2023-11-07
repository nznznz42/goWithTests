package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("z")
	expected := "zzzzzzzzzz"

	if repeated != expected {
		t.Errorf("expected: %q\n recieved: %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("z")
	}
}