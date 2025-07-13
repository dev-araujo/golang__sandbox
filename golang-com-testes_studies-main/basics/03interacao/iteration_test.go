package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repetead("a", 5)
	expectation := "aaaaa"

	if repeated != expectation {
		t.Errorf("expectation '%s', received '%s'", expectation, repeated)
	}
}
func BenchmarkRepeated(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Repetead("a", 5)
	}
}
