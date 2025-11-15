package cmd

import (
	"testing"
)

func TestValidateSizeStr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "valid number", input: "100", expected: true},
		{name: "valid with unit k", input: "100k", expected: true},
		{name: "valid with unit m", input: "100m", expected: true},
		{name: "valid with unit g", input: "100g", expected: true},
		{name: "invalid empty", input: "", expected: true},
		{name: "invalid letters", input: "abc", expected: false},
		{name: "invalid negative", input: "-100", expected: false},
		{name: "edge case zero", input: "0", expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateSizeStr(tt.input)
			if got != tt.expected {
				t.Errorf("ValidateSizeStr(%q) = %v, want %v",
					tt.input, got, tt.expected)
			}
		})
	}
}
func TestValidateReturn(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "valid return", input: "301 https://example.com", expected: true},
		{name: "invalid http status code", input: "600 https://example.com", expected: false},
		{name: "invalid url", input: "301 htt://example.com/", expected: false},
		{name: "invalid url and status code", input: "600 htt://example.com/", expected: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateReturn(tt.input)
			if got != tt.expected {
				t.Errorf("ReturnValidate(%q) = %v, want %v",
					tt.input, got, tt.expected)
			}
		})
	}
}

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "valid url", input: "https://example.com", expected: true},
		{name: "invalid url", input: "htt://example.com/", expected: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateURL(tt.input)
			if got != tt.expected {
				t.Errorf("URLValidate(%q) = %v, want %v",
					tt.input, got, tt.expected)
			}
		})
	}
}
