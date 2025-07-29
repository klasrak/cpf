package cpf

import (
	"fmt"
	"testing"
)

const (
	EXAMPLE_VALID_CPF          = "620.223.434-25" // This is not a real CPF, just a valid example
	EXAMPLE_VALID_CPF_INT      = 62022343425
	EXAMPLE_VALID_CPF_UINT     = uint(62022343425)
	EXAMPLE_VALID_CPF_UINT64   = uint64(62022343425)
	EXAMPLE_VALID_CPF_INT64    = int64(62022343425)
	EXAMPLE_VALID_CPF_STRING   = "62022343425"
	EXAMPLE_INVALID_CPF        = "111.111.111-11"
	EXAMPLE_INVALID_CPF_INT    = 11111111111
	EXAMPLE_INVALID_CPF_UINT   = uint(11111111111)
	EXAMPLE_INVALID_CPF_UINT64 = uint64(11111111111)
	EXAMPLE_INVALID_CPF_INT64  = int64(11111111111)
)

func TestIsValid(t *testing.T) {
	type testCase struct {
		name     string
		cpf      any
		expected bool
	}

	var testCases []testCase

	for i := range 1000 {
		testCases = append(testCases, testCase{
			name:     fmt.Sprintf("Valid CPF #%d", i+1),
			cpf:      New("SP"),
			expected: true,
		})
	}

	testCases = append(testCases, []testCase{
		{
			name:     "Invalid CPF - All digits equal",
			cpf:      "111.111.111-11",
			expected: false,
		},
		{
			name:     "Invalid CPF - Incorrect verifier",
			cpf:      "123.456.789-00",
			expected: false,
		},
		{
			name:     "Invalid CPF - Too short",
			cpf:      "123.456",
			expected: false,
		},
		{
			name:     "Invalid CPF - Too long",
			cpf:      "123.456.789-012",
			expected: false,
		},
		{
			name:     "Invalid CPF - Non-numeric characters",
			cpf:      "123.456.789-AB",
			expected: false,
		},
		{
			name:     "Invalid CPF - Empty string",
			cpf:      "",
			expected: false,
		},
		{
			name:     "Invalid CPF - Nil",
			cpf:      nil,
			expected: false,
		},
		{
			name:     "Invalid CPF - Negative integer",
			cpf:      -12345678901,
			expected: false,
		},
		{
			name:     "Invalid CPF - Zero",
			cpf:      0,
			expected: false,
		},
		{
			name:     "Invalid CPF - Non-integer string",
			cpf:      "not_a_cpf",
			expected: false,
		},
		{
			name:     "Invalid CPF - Float",
			cpf:      1234567890.12,
			expected: false,
		},
		{
			name:     "Invalid CPF - Boolean",
			cpf:      true,
			expected: false,
		},
		{
			name:     "Invalid CPF - int32",
			cpf:      int32(2147483647),
			expected: false,
		},
		{
			name:     "Invalid CPF - int64",
			cpf:      int64(1234567890123456789),
			expected: false,
		},
		{
			name:     "Invalid CPF - uint",
			cpf:      uint(1234567890123456789),
			expected: false,
		},
		{
			name:     "Invalid CPF - uint32",
			cpf:      uint32(4294967295),
			expected: false,
		},
		{
			name:     "Invalid CPF - uint64",
			cpf:      uint64(18446744073709551615),
			expected: false,
		},
		{
			name:     "Invalid CPF - Special characters",
			cpf:      "!@#$%^&*()",
			expected: false,
		},
		{
			name:     "Valid CPF int64",
			cpf:      int64(New("SP")),
			expected: true,
		},
		{
			name:     "Valid CPF uint64",
			cpf:      uint64(New("SP")),
			expected: true,
		},
		{
			name:     "Valid CPF string",
			cpf:      fmt.Sprintf("%d", New("SP")),
			expected: true,
		},
		{
			name:     "Valid CPF with mask #1",
			cpf:      WithMask("SP"),
			expected: true,
		},
		{
			name:     "Valid CPF with mask #2",
			cpf:      WithMask("RJ"),
			expected: true,
		},
		{
			name:     "Valid CPF with mask #3",
			cpf:      WithMask("SC"),
			expected: true,
		},
		{
			name:     "Valid CPF with mask #4",
			cpf:      WithMask("MG"),
			expected: true,
		},
		{
			name:     "Valid CPF with mask #5",
			cpf:      WithMask("RS"),
			expected: true,
		},
		{
			name:     "Valid CPF with unmask",
			cpf:      Unmask(Mask(New("SP"))),
			expected: true,
		},
	}...)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsValid(tc.cpf)
			if result != tc.expected {
				t.Errorf("Expected %v for CPF %v, got %v", tc.expected, tc.cpf, result)
			}
		})
	}
}
func BenchmarkIsValidInt(b *testing.B) {
	for b.Loop() {
		IsValid(EXAMPLE_VALID_CPF_INT)
	}
}

func BenchmarkIsValidUint(b *testing.B) {
	for b.Loop() {
		IsValid(EXAMPLE_VALID_CPF_UINT)
	}
}

func BenchmarkIsValidUint64(b *testing.B) {
	for b.Loop() {
		IsValid(EXAMPLE_VALID_CPF_UINT64)
	}
}

func BenchmarkIsValidInt64(b *testing.B) {
	for b.Loop() {
		IsValid(EXAMPLE_VALID_CPF_INT64)
	}
}

func BenchmarkIsValidString(b *testing.B) {
	for b.Loop() {
		IsValid(EXAMPLE_VALID_CPF_STRING)
	}
}

func BenchmarkIsValidWithMask(b *testing.B) {
	for b.Loop() {
		IsValid(EXAMPLE_VALID_CPF)
	}
}
