package main

import "testing"

func main() {

}

func TestArabToRoman(t *testing.T) {
	type args struct {
		arabNum int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
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

func Test_romanToArab(t *testing.T) {
	type args struct {
		romanNum string
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
			if got := romanToArab(tt.args.romanNum); got != tt.want {
				t.Errorf("romanToArab() = %v, want %v", got, tt.want)
			}
		})
	}
}
