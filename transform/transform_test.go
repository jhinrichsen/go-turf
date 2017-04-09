package transform

import "testing"

func TestTransform(t *testing.T) {
	if 500 != transform(5, 1, 10, 100, 1000) {
		t.Fail()
	}
}
