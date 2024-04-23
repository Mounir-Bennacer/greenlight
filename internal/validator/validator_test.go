package validator_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"greenlight.mounirbennacer.com/internal/validator"
)

func TestNotBlank(t *testing.T) {
	tests := []struct {
		value    string
		expected bool
	}{
		{"", false},
		{" ", true},
		{"hello", true},
	}

	for _, test := range tests {
		result := validator.NotBlank(test.value)
		assert.Equal(t, test.expected, result, "NotBlank result mismatch")
	}
}

func TestMaxChars(t *testing.T) {
	tests := []struct {
		value    string
		n        int
		expected bool
	}{
		{"", 5, true},
		{"hello", 10, true},
		{"hello", 4, false},
	}

	for _, test := range tests {
		result := validator.MaxChars(test.value, test.n)
		assert.Equal(t, test.expected, result, "MaxChars result mismatch")
	}
}

func TestPermittedValues(t *testing.T) {
	tests := []struct {
		value           interface{}
		permittedValues []interface{}
		expected        bool
	}{
		{1, []interface{}{1, 2, 3}, true},
		{4, []interface{}{1, 2, 3}, false},
		{"apple", []interface{}{"apple", "banana"}, true},
	}

	for _, test := range tests {
		result := validator.PermittedValue(test.value, test.permittedValues...)

		assert.Equal(t, test.expected, result, "PermittedValue result mismatch")
	}
}

func TestMatches(t *testing.T) {
	tests := []struct {
		value    string
		rx       *regexp.Regexp
		expected bool
	}{
		{"user@example.com", validator.EmailRX, true},
		{"invalid-email", validator.EmailRX, false},
	}

	for _, test := range tests {
		result := validator.Matches(test.value, test.rx)
		assert.Equal(t, test.expected, result, "Matches result mismatch")
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		values   []interface{}
		expected bool
	}{
		{values: []interface{}{1, 2, 3}, expected: true},
		{values: []interface{}{1, 2, 2}, expected: false},
		{values: []interface{}{"apple", "banana", "cherry"}, expected: true},
		{values: []interface{}{"apple", "banana", "apple"}, expected: false},
	}
	for _, test := range tests {
		result := validator.Unique(test.values)
		assert.Equal(t, test.expected, result, "Unique result mismatch")
	}
}
