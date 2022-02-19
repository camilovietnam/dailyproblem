package water_contained

import (
	"testing"
)

var tests = []struct {
	name     string
	input    []uint
	expected uint
}{
	{
		name:     "no walls",
		input:    []uint{},
		expected: 0,
	},
	{
		name:     "3 walls, 1 unit of water",
		input:    []uint{1, 0, 1},
		expected: 1,
	},
	{
		name:     "3 walls, 1 unit of water (from example)",
		input:    []uint{2, 1, 2},
		expected: 1,
	},
	{
		name:     "6 walls, 8 units of water",
		input:    []uint{3, 0, 1, 3, 0, 5},
		expected: 8,
	},
	{
		name:     "6 walls, 0 units of water",
		input:    []uint{0, 1, 2, 2, 1, 0},
		expected: 0,
	},
	{
		name:     "5 walls, 0 units of water",
		input:    []uint{5, 4, 3, 2, 1},
		expected: 0,
	},
}

func Test_WaterContained(t *testing.T) {
	for _, test := range tests {
		t.Log("Running test....", test.name)
		got := calculateWater(test.input)

		if got != test.expected {
			t.Fatalf("Failed test: %s", test.name)
		}
	}
}
