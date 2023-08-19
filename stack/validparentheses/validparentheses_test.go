package validparentheses

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", want: true, args: args{"{}[]()"}},
		{name: "test2", want: false, args: args{"{}[]("}},
		{name: "test3", want: false, args: args{"(]"}},
		{name: "test4", want: false, args: args{"({[)"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
