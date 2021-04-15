package calculator

import (
	"testing"
)

func TestNeedCalc(t *testing.T) {
	const str = "/calculator Добавить НДС 100"
	const want = "120"
	got := NeedCalc(str, 0.2, 0.3)
	if got != want {
		t.Fatalf("Command: %q => got = %q, want=%q", str, got, want)
	}
}
