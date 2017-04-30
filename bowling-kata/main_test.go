package main

import "testing"

// The first test
func TestGutterGame(t *testing.T) {
	g := Game{}
	g.roll(make([]int, 20)...)
	if g.traditionalScore() != 0 {
		t.Fail()
	}
}

// The second test
func TestAllOnes(t *testing.T) {
	g := Game{}
	for i := 0; i < 20; i++ {
		g.roll(1)
	}
	if g.traditionalScore() != 20 {
		t.Fail()
	}
}

// The third test
func TestOneSpare(t *testing.T) {
	g := Game{}
	g.roll([]int{5, 5, 3, 4}...)
	g.roll(make([]int, 16)...)
	if g.traditionalScore() != 17 {
		t.Errorf("expected 17 but got %v\n", g.traditionalScore())
	}
}

// The fourth test
func TestOneStrike(t *testing.T) {
	g := Game{}
	g.roll([]int{10, 3, 4}...)
	g.roll(make([]int, 16)...)
	if g.traditionalScore() != 24 {
		t.Errorf("Expected 24 but got %v\n", g.traditionalScore())
	}
}
