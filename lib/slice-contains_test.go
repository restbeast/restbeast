package lib

import "testing"

func TestSliceContains(t *testing.T) {
	set := []struct {
		Slice  []string
		Item   string
		Result bool
	}{
		{[]string{"x", "y"}, "x", true},
		{[]string{"x", "y"}, "z", false},
	}

	for _, run := range set {
		result := sliceContains(run.Slice, run.Item)

		if result != run.Result {
			t.Errorf("sliceContains incorrect, got: %t, want: %t.", result, run.Result)
		}
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("sliceContains did not panic with wrong typed element")
		}
	}()

	sliceContains("not-a-slice", "a-string")
}
