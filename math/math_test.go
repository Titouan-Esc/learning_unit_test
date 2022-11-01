package math

import (
	"testing"
)

type AddData struct {
	x, y   int
	result int
}

func TestAdd(t *testing.T) {
	testData := []AddData{
		{1, 2, 3},
		{2, 3, 5},
		{1, 1, 2},
	}

	for _, value := range testData {
		result := Add(value.x, value.y)

		if result != value.result {
			t.Errorf("Add(%d, %d) FAILED. Expected %d, got %d\n", value.x, value.y, value.result, result)
		} else {
			t.Logf("Add(%d, %d) PASSED. Expected %d, got %d\n", value.x, value.y, value.result, result)
		}
	}
}
