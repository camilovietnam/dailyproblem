package ranges

import "testing"

func TestMinRange(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		listeners []int
		antennas  []int
		want      int
	}{
		{
			name:      "one listener, two antennas",
			listeners: []int{2},
			antennas:  []int{4, 10},
			want:      2,
		},
		{
			name:      "two listener, two antennas",
			listeners: []int{2, 15},
			antennas:  []int{4, 10},
			want:      5,
		},
		{
			name:      "two listener, two antennas from example",
			listeners: []int{1, 5, 11, 20},
			antennas:  []int{4, 8, 15},
			want:      5,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := MinRange(tt.listeners, tt.antennas)
			if tt.want != got {
				t.Fatalf("failed: %s - expected %d, got %d", tt.name, tt.want, got)
			} else {
				t.Logf("passed: %s", tt.name)
			}
		})
	}
}
