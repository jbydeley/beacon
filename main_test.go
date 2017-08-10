package main

import (
	"testing"
)

func Test_hasValidConfig(t *testing.T) {
	tests := []struct {
		input    beaconConfig
		expected error
	}{
		{
			beaconConfig{
				author{
					Name:  "John Smith",
					Email: "john.smith@example.com",
				},
			},
			nil,
		},
		{
			beaconConfig{
				author{
					Name:  "John Smith",
					Email: "",
				},
			},
			errAuthorEmailRequired,
		},
		{
			beaconConfig{
				author{
					Name:  "",
					Email: "john.smith@example.com",
				},
			},
			errAuthorNameRequired,
		},
		{
			beaconConfig{
				author{
					Name:  "",
					Email: "",
				},
			},
			errAuthorNameRequired,
		},
	}

	for _, tt := range tests {
		actual := hasValidConfig(tt.input)
		if actual != tt.expected {
			t.Errorf("hasValidConfig(%+v) did not return expected value (%v)", tt.input, tt.expected)
		}
	}
}
