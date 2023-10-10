package tests

import (
	"fmt"
	"testing"
)

func TestTestFunc(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testFunc",
			args: args{
				x: 1,
				y: 2,
			},
			want: 3,
		},
		{
			name: "testFunc",
			args: args{
				x: 2,
				y: 2,
			},
			want: 4,
		},
		{
			name: "testFunc",
			args: args{
				x: 3,
				y: 2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("testFunc with args: %v\n", tt.args)
			if got := CalcFunc(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("testFunc() = %v, want %v", got, tt.want)
			}
		})
	}

}