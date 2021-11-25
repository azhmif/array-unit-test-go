package main

import "testing"

func TestMaxA(t *testing.T) {
	t.Log("Max Value  A : ", findMaxRow(a))

	if findMaxRow(a) != maxA {
		t.Error("Correct Answer", maxA)
	}
}

func TestMaxB(t *testing.T) {
	t.Log("Max Value  b : ", findMaxRow(b))

	if findMaxRow(b) != maxB {
		t.Error("Correct Answer", maxA)
	}
}

func TestMaxC(t *testing.T) {
	t.Log("Max Value C : ", findMaxRow(c))

	if findMaxRow(c) != maxC {
		t.Error("Correct Answer", maxC)
	}
}
