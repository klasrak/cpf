package cpf

import "testing"

func TestMask(t *testing.T) {
	type testCase struct {
		cpf      any
		expected string
	}

	testCases := []testCase{
		{11144477735, "111.444.777-35"},
		{2147483647, "021.474.836-47"},
		{98765432100, "987.654.321-00"},
		{int32(214748364), "002.147.483-64"},
		{int64(12345678901), "123.456.789-01"},
		{uint(1234567890), "012.345.678-90"},
		{uint32(123456789), "001.234.567-89"},
		{uint64(123456789012), "234.567.890-12"}, // Note: uint64 may exceed valid CPF range, but we still format it
		{"invalid", ""},
		{nil, ""},
		{testCase{}, ""},
		{-1, ""},
		{0, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			result := Mask(tc.cpf)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func BenchmarkMask(b *testing.B) {
	for b.Loop() {
		Mask(11144477735)
	}
}
