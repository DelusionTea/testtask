package FindNum

import "testing"

func Test_finder(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "positive test 1",
			args: args{
				text: "fRETre3445 43#$% 32424234 #$@$34 2323",
			},
			want:    32424234,
			wantErr: false,
		},
		{
			name: "bignum test ",
			args: args{
				text: "fRETre3445 43#$% 3242423999994 #$@$34 2323",
			},
			want:    3242423999994,
			wantErr: false,
		},
		{
			name: "really bignum test ",
			args: args{
				text: "fRE324242399999000003Tre3445 43#$% 324242399999000004 #$@$34 2323",
			},
			want:    324242399999000004,
			wantErr: false,
		},
		{
			name: "mistik test ",
			args: args{
				text: "fRETre00003445 000043#$% 000003242423#$@$000000034 00002323",
			},
			want:    3242423,
			wantErr: false,
		},
		{
			name: "without nums test ",
			args: args{
				text: "fRETre #$% #$@$ ",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty string test ",
			args: args{
				text: "",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "negative NUM",
			args: args{
				text: "fRETre3445 43#$% -32424234 #$@$34 2323",
			},
			want:    3445,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := finder(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("finder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("finder() = %v, want %v", got, tt.want)
			}
		})
	}
}
