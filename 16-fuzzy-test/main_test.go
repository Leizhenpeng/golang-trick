package main

import (
	"strings"
	"testing"

	"github.com/go-playground/assert"
)

func TestEmailRead(t *testing.T) {
	tests := []struct {
		name  string
		email string
		wantU string
		wantD string
	}{
		{
			"test 1",
			"22@11.com",
			"22",
			"11.com",
		},
		{
			"test 2",
			"qwq@qq.com",
			"qwq",
			"qq.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, d := EmailRead(tt.email)
			if u != tt.wantU {
				t.Errorf("EmailRead() gotUsername = %v, want %v", u, tt.wantU)
			}
			if d != tt.wantD {
				t.Errorf("EmailRead() gotDomain = %v, want %v", d, tt.wantD)
			}
		})
	}
}

func FuzzTestEmailRead(f *testing.F) {
	testcases := []string{
		"222@11.com",
		"12313@qq.com",
		"22@222.com",
	}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, email string) {
		username, domain :=
			EmailRead(email)
		if strings.Contains(email, "@") {
			assert.Equal(
				t, username+"@"+domain, email)
		}
	})
}
