package matcher

import "testing"

func Test_MatchesTokens(t *testing.T) {
	type args struct {
		includes []string
		msg      string
		fallback bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check matching with positive performs contain check",
			args: args{
				includes: []string{"test123"},
				msg:      "/var/log/something-test123",
				fallback: false,
			},
			want: true,
		},
		{
			name: "check matching with negative performs contain check",
			args: args{
				includes: []string{"test123"},
				msg:      "/var/log/something-test23",
				fallback: false,
			},
			want: false,
		},
		{
			name: "check empty includes results in pass through",
			args: args{
				includes: []string{},
				msg:      "/var/log/something-test23",
				fallback: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchesTokens(tt.args.includes, tt.args.msg, tt.args.fallback); got != tt.want {
				t.Errorf("matchesTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}
