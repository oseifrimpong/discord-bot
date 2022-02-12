package internal

import (
	"reflect"
	"testing"

	"github.com/andersfylling/snowflake/v5"
)

func Test_convertStringtoSnowflake(t *testing.T) {
	type args struct {
		userIDStr string
	}
	tests := []struct {
		name string
		args args
		want snowflake.Snowflake
	}{
		{
			name: "Test convertStringtoSnowflake good",
			args: args{userIDStr: "<@!647002363202371594>"},
			want: 647002363202371594,
		},
		{
			name: "Test convertStringtoSnowflake bad",
			args: args{userIDStr: "usge-kajuej"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertStringToSnowflake(tt.args.userIDStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertStringtoSnowflake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddressChecker(t *testing.T) {
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
			if got := AddressChecker(tt.args.walletAddress); got != tt.want {
				t.Errorf("addressChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
