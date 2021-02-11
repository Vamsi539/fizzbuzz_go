package main

import (
	"testing"
)

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{name: "add", args: args{
			2,
			3,
		}, want:5}, // TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sub(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sub", args: args{5,
		6}, want:-1},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sub(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mul(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "mul", args: args{5 ,
		6}, want:30},
		{name: "mul1", args:args{-5,-3}, want:15},
		{name: "mul2", args: args{-5,5}, want:-25},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mul(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("mul() = %v, want %v", got, tt.want)
			}
		})
	}
}
/*
func Test_div(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "div", args: args{6, 3}, want:2}, // TODO: Add test cases.
		{name: "div1", args: args{6, 2}, want:3},
		{name: "div2", args:args{10,2	}, want:5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := div(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("div() = %v, want %v", got, tt.want)
			}
		})
	}
}*/

func Test_modl(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name:"modl", args:args{10,3}, want:1}, // TODO: Add test cases.
		{name: "modl1", args: args{12, 5}, want:2},
		{name: "modl2", args: args{7,3	}, want:1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modl(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("modl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_div1(t *testing.T) {
	type args struct {
		a float32
		b float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{name: "div5", args: args{5,2}, want:2.5},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := div(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fizzbuzz(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "for 1", args: args{1}, want:""}, // TODO: Add test cases.
		{name: "for 3", args: args{3}, want: "fizz"},
		{name: "for 5", args: args{5}, want: "buzz"},
		{name: "for 15", args: args{15}, want: "fizzbuzz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzbuzz(tt.args.i); got != tt.want {
				t.Errorf("fizzbuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_re_fizzbuzz(t *testing.T) {
	type args struct {
		i float32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "re fizz 1", args: args{1	}, want:""}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := re_fizzbuzz(tt.args.i); got != tt.want {
				t.Errorf("re_fizzbuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}