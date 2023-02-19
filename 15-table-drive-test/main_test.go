package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

//bad case
func TestSplitEmail(t *testing.T) {
	username, domain := SplitEmail("1233@qq.com")
	require.NoError(t, nil)
	assert.Equal(t, username, "1233")
	assert.Equal(t, domain, "qq.com")

	username, domain = SplitEmail("22qqew@gami.com")
	require.NoError(t, nil)
	assert.Equal(t, username, "22qqew")
	assert.Equal(t, domain, "gami.com")

	username, domain = SplitEmail("123213@88.com")
	require.NoError(t, nil)
	assert.Equal(t, username, "123213")
	assert.Equal(t, domain, "88.com")

	username, domain = SplitEmail("123213@228.com")
	require.NoError(t, nil)
	assert.Equal(t, username, "123213")
	assert.Equal(t, domain, "228.com")

}

// table drive test
func TestSplitEmail2(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		username string
		domain   string
	}{
		{
			name:     "test 1",
			email:    "22qqew@gami.com",
			username: "22qqew",
			domain:   "gami.com",
		},
		{
			name:     "test 2",
			email:    "1233@qq.com",
			username: "1233",
			domain:   "qq.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			username, domain := SplitEmail(tt.email)
			require.NoError(t, nil)
			assert.Equal(t, username, tt.username)
			assert.Equal(t, domain, tt.domain)
		})
	}
}
