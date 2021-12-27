package bot

import (
	"reflect"
	"testing"

	"github.com/andersfylling/disgord"
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
			if got := convertStringtoSnowflake(tt.args.userIDStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertStringtoSnowflake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_msgHandler(t *testing.T) {
	type args struct {
		session disgord.Session
		evt     *disgord.MessageCreate
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msgHandler(tt.args.session, tt.args.evt)
		})
	}
}
