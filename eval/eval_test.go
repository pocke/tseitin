package eval

import (
	"testing"
)

func TestEvaluatorSucc(t *testing.T) {
	ev := Evaluator{x: "A"}
	if got := ev.Succ(); got != "A" {
		t.Errorf("Expected: A, but got %s", got)
	}
	if got := ev.x; got != "B" {
		t.Errorf("Expected: B, but got %s", got)
	}

	ev = Evaluator{x: "Z"}
	if got := ev.Succ(); got != "Z" {
		t.Errorf("Expected: Z, but got %s", got)
	}
	if got := ev.x; got != "AA" {
		t.Errorf("Expected: AA, but got %s", got)
	}

	ev = Evaluator{x: "AZ"}
	if got := ev.Succ(); got != "AZ" {
		t.Errorf("Expected: AZ, but got %s", got)
	}
	if got := ev.x; got != "BA" {
		t.Errorf("Expected: BA, but got %s", got)
	}

	ev = Evaluator{x: "ZZ"}
	if got := ev.Succ(); got != "ZZ" {
		t.Errorf("Expected: ZZ, but got %s", got)
	}
	if got := ev.x; got != "AAA" {
		t.Errorf("Expected: AAA, but got %s", got)
	}
}
