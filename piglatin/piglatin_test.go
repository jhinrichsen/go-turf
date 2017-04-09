
package main

import "testing"

func TestVowel(t *testing.T) {
	if !IsVowel("another") ||
	  !IsVowel("explain") ||
	  IsVowel("belong") ||
	  IsVowel("zero") {
	  		t.Fail()
	}
}

func TestBeast(t *testing.T)  {
	if piglatin("beast") != "east-bay" { t.Fail() }
}

func TestDough(t *testing.T)  {
	if piglatin("dough") != "ough-day" { t.Fail() }
}

func TestHappy(t *testing.T)  {
	if piglatin("happy") != "appy-hay" { t.Fail() }
}

func TestQuestion(t *testing.T)  {
	if piglatin("question") != "uestion-qay" { t.Fail() }
}

func TestAnother(t *testing.T) {
	if piglatin("another") != "another-way" { t.Fail() }
}

func TestIf(t *testing.T) {
	if piglatin("if") != "if-way" { t.Fail() }
}
