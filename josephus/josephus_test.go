package main

import "testing"

// Recreational Mathematics 2003, chapter 37
const want = 73
const n = 100

func TestPow(t *testing.T) {
	want := 8
	got := pow(2, 3)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestJosephus1(t *testing.T) {
	got := josephus1(n)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestJosephus2(t *testing.T) {
	got := josephus2(n)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}
