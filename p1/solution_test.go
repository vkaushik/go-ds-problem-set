package p1

import "testing"

func TestLenOfLongestSubStWithUniqueChars(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"edge case 1, empty string: ", args{""}, 0},
		{"edge case 2, single char string: a", args{"a"}, 1},
		{"edge case 3, all same char string: aa", args{"aa"}, 1},
		{"edge case 4, all different char string: ab", args{"ab"}, 2},
		{"edge case 5, all different char string: abc", args{"abc"}, 3},
		{"positive 1, string: abca", args{"abca"}, 3},
		{"positive 2, string: abcabcdac", args{"abcabcdac"}, 4},
		{"positive 3, string: abac", args{"abac"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LenOfLongestSubStWithUniqueChars(tt.args.str); got != tt.want {
				t.Errorf("LenOfLongestSubStWithUniqueChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
