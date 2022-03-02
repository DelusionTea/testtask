package main

import (
	"testing"
)

func TestArabToRoman(t *testing.T) {
	type args struct {
		arabNum int
	}
	tests := []struct {
		name    string
		args    args
		arabNum int
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "positive test",
			args: args{
				arabNum: 55,
			},
			want:    "LV",
			wantErr: false,
		},
		{
			name: "negative test 1 ",
			args: args{
				arabNum: 5999,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "negative test 0 ",
			args: args{
				arabNum: -1,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ArabToRoman(tt.args.arabNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArabToRoman() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ArabToRoman() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//func Test_romanToArab(t *testing.T) {
//	type args struct {
//		romanNum string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    int
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{
//			name: "positive test 1",
//			args: args{
//				romanNum: "LV",
//			},
//			want:    55,
//			wantErr: false,
//		},
//		{
//			name: "negative test 1",
//			args: args{
//				romanNum: "AAB",
//			},
//			want:    0,
//			wantErr: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := romanToArab(tt.args.romanNum)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("romanToArab() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("romanToArab() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
