package bot

import (
	"testing"

	"github.com/andersfylling/disgord"
)

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
