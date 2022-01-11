package main

import "testing"

func Test_findFirstStringInBracket(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "return ing",
			args: args{
				str: "test(ing)",
			},
			want: "ing",
		},
		{
			name: "return empty string",
			args: args{
				str: "bastian()",
			},
			want: "",
		},
		{
			name: "return whitespace",
			args: args{
				str: "bastian ( )",
			},
			want: " ",
		},
		{
			name: "return empty string with str is empty string",
			args: args{
				str: "",
			},
			want: "",
		},
		{
			name: "return empty string since there is no brackets in str",
			args: args{
				str: "bastian)",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstStringInBracket(tt.args.str); got != tt.want {
				t.Errorf("findFirstStringInBracket() = %v, want %v", got, tt.want)
			}
		})
	}
}
