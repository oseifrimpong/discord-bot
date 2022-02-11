package internal

import "testing"

func Test_addressChecker(t *testing.T) {
	type args struct {
		walletAddress string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				walletAddress: "0x0",
			},
			want: false,
		},

		{
			name: "Test 2",
			args: args{
				walletAddress: "AAeS8tdU9sXwJDGrwg7TkrNhwbLKhkwbkE9bkWfvwiCp", //correct for now.
				// walletAddress: "EPi9bC1qhd1hLh9QqxMkHHxK2Lzw3BHTWrt4cjPLwhPU", //correct going forward.
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addressChecker(tt.args.walletAddress); got != tt.want {
				t.Errorf("addressChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
