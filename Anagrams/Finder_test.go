package Anagram

import (
	"reflect"
	"testing"
)

func Test_finder(t *testing.T) {
	type args struct {
		text []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "positive test 1",
			args: args{
				text: []string{"cat", "act", "dinosaur", "testing", "setting"},
			},
			want:    []string{"cat", "act", "testing", "setting"},
			wantErr: false,
		},
		{
			name: "positive test 2",
			args: args{
				text: []string{"cat", "dododo", "act", "ododod", "dinosaur", "testing", "setting"},
			},
			want:    []string{"cat", "act", "dododo", "ododod", "testing", "setting"},
			wantErr: false,
		},
		{
			name: "negative test 1",
			args: args{
				text: []string{"cat"},
			},
			want:    []string{""},
			wantErr: true,
		},
		{
			name: "negative test 2",
			args: args{
				text: []string{"act", "dinosaur", "testing"},
			},
			want:    []string{""},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := finder(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("finder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("finder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
