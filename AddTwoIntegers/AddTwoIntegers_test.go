package AddTwoIntegers

import "testing"

func TestSum(t *testing.T) {
	if sum(12, 5) != 17 {
		t.Error(`sum(12, 15) != 17`)
	}

	if sum(-10, 4) != -6 {
		t.Error(`sum(-10, 4) != -6`)
	}
}
