package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	for idx := 0; idx < b.N; idx++ {
		topKFrequent(nums, k)
	}
}
func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [1,1,1,2,2,3], k = 2",
			args: args{nums: []int{1, 1, 1, 2, 2, 3}, k: 2},
			want: []int{1, 2},
		},
		{
			name: "nums = [1], k = 1",
			args: args{nums: []int{1}, k: 1},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
