package maxes

import (
	"reflect"
	"testing"
)

func Test_calculateMaxSubarray(t *testing.T) {
	type args struct {
		numbers []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic test",
			args: args{
				numbers: []int{2, 1, 1, 3},
				k:       3,
			},
			want: []int{2, 3},
		},
		{
			name: "test k = n",
			args: args{
				numbers: []int{7, 13, 1, 4, 9},
				k:       5,
			},
			want: []int{13},
		}, {
			name: "test with lots of small numbers",
			args: args{
				numbers: []int{7, 1, 1, 1, 1, 1, 5, 1, 1, 1, 1, 1, 9},
				k:       10,
			},
			want: []int{7, 5, 5, 9},
		}, {
			name: "array smaller than k",
			args: args{
				numbers: []int{1, 2, 3},
				k:       4,
			},
			want: nil,
		},
		{
			name: "example from email",
			args: args{
				numbers: []int{10, 5, 2, 7, 8, 7},
				k:       3,
			},
			want: []int{10, 7, 8, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMaxSubarray(tt.args.numbers, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateMaxSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
