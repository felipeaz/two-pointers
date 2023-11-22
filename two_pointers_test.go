package twopointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := map[string]struct {
		text           string
		expectResponse bool
	}{
		"Valid palindrome": {
			text:           "A man, a plan, a canal: Panama",
			expectResponse: true,
		},
		"Invalid palindrome": {
			text:           "This is definitely not a palindrome",
			expectResponse: false,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			resp := IsPalindrome(tt.text)

			assert.Equal(t, tt.expectResponse, resp)
		})
	}
}
