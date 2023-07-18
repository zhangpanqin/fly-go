package test

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add_aa(t *testing.T) {
	t.Run("测试 2", func(t *testing.T) {
		if got := add(1, 2); got != 4 {
			t.Errorf("add() = %v, want %v", got, 4)
		}
	})

	t.Run("测试 3", func(t *testing.T) {
		if got := add(2, 2); got != 4 {
			t.Errorf("add() = %v, want %v", got, 4)
		}
	})
}
