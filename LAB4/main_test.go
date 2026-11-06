package main

import "testing"

func TestVerifyFlag(t *testing.T) {
	cases := []struct {
		candidate string
		want      bool
	}{
		{"cryptoCTF{meowmeow}", true},
		{"", false},
		{"cryptoCTF{meow}", false},
		{"cryptoCTF{meowmeowmeow}", false},
		{"cryptoCTF{Meowmeow}", false},
		{"prefixcryptoCTF{meowmeow}", false},
		{"cryptoCTF{meowmeow}suffix", false},
		{"cryptoCTF{meowmeow}\n", false},
	}
	for _, tc := range cases {
		t.Run(tc.candidate, func(t *testing.T) {
			got, err := verifyFlag(tc.candidate)
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.want {
				t.Errorf("verifyFlag(%q) = %v; want %v", tc.candidate, got, tc.want)
			}
		})
	}
}
