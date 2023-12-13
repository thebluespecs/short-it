package services

import (
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	testCases := []struct {
		input uint64
	}{
		{0},
		{1},
		{42},
		{123},
		{987654321},
	}

	for _, tc := range testCases {
		t.Run("TestEncodeDecode_"+encode(tc.input), func(t *testing.T) {
			encoded := encode(tc.input)
			decoded := decode(encoded)

			if decoded != tc.input {
				t.Errorf("Expected decoding of %s to be %d, but got %d", encoded, tc.input, decoded)
			}
		})
	}
}

func TestEdgeCases(t *testing.T) {
	// Test edge cases
	zeroEncoded := encode(0)
	if zeroEncoded != string(alphabet[0]) {
		t.Errorf("Expected encoding of 0 to be %s, but got %s", string(alphabet[0]), zeroEncoded)
	}

	emptyStringDecoded := decode("")
	if emptyStringDecoded != 0 {
		t.Errorf("Expected decoding of an empty string to be 0, but got %d", emptyStringDecoded)
	}
}

func TestInvalidInput(t *testing.T) {
	// Test invalid input for decode
	invalidInput := "abc123)" // contains characters not in the alphabet
	result := decode(invalidInput)

	if result != 0 {
		t.Errorf("Expected decoding of invalid input to be 0, but got %d", result)
	}
}

