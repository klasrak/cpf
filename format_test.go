package cpf

import (
	"sync"
	"testing"
)

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
		{"001.234.567-89", "001.234.567-89"},
		{"12345678901", "123.456.789-01"},
		{"123", ""},
		{"!@#$%^&*()", ""},
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

func TestMaskAsync(t *testing.T) {
	type testCase struct {
		cpf      any
		expected string
	}

	type testResult struct {
		index  int
		result string
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
		{"001.234.567-89", "001.234.567-89"},
		{"12345678901", "123.456.789-01"},
		{"123", ""},
		{"!@#$%^&*()", ""},
		{nil, ""},
		{testCase{}, ""},
		{-1, ""},
		{0, ""},
	}

	var wg sync.WaitGroup
	results := make([]testResult, len(testCases))

	for i, tc := range testCases {
		wg.Add(1)
		go func(i int, tc testCase) {
			defer wg.Done()
			result := Mask(tc.cpf)
			results[i] = testResult{index: i, result: result}
		}(i, tc)
	}
	wg.Wait()

	for _, res := range results {
		expected := testCases[res.index].expected
		if res.result != expected {
			t.Errorf("Test case %d: Expected %s, got %s", res.index, expected, res.result)
		}
	}
}

func TestWithMask(t *testing.T) {
	type testCase struct {
		state    string
		expected int
	}

	testCases := []testCase{
		{"RS", 14},
		{"DF", 14},
		{"GO", 14},
		{"MS", 14},
		{"MT", 14},
		{"TO", 14},
		{"AC", 14},
		{"AM", 14},
		{"AP", 14},
		{"PA", 14},
		{"RO", 14},
		{"RR", 14},
		{"CE", 14},
		{"MA", 14},
		{"PI", 14},
		{"AL", 14},
		{"PB", 14},
		{"PE", 14},
		{"RN", 14},
		{"BA", 14},
		{"SE", 14},
		{"MG", 14},
		{"ES", 14},
		{"RJ", 14},
		{"SP", 14},
		{"PR", 14},
		{"SC", 14},
		{"XX", 0},    // Invalid state
		{"", 0},      // Empty state
		{"SP123", 0}, // Invalid state format
		{"!@#$", 0},  // Invalid characters
	}

	for _, tc := range testCases {
		t.Run(tc.state, func(t *testing.T) {
			cpf := WithMask(tc.state)

			if len(cpf) != tc.expected {
				t.Errorf("Expected CPF with length %d for state %s, got %s", tc.expected, tc.state, cpf)
			}
		})
	}
}

func TestUnmask(t *testing.T) {
	type testCase struct {
		cpf      string
		expected string
	}

	testCases := []testCase{
		{"111.444.777-35", "11144477735"},
		{"001.234.567-89", "00123456789"},
		{"12345678901", "12345678901"},
		{"123", ""},
		{"!@#$%^&*()", ""},
		{"", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.cpf, func(t *testing.T) {
			result := Unmask(tc.cpf)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestUnmaskToInt(t *testing.T) {
	type testCase struct {
		cpf      string
		expected int
	}

	testCases := []testCase{
		{"111.444.777-35", 11144477735},
		{"001.234.567-89", 123456789},
		{"12345678901", 12345678901},
		{"123", -1},
		{"!@#$%^&*()", -1},
		{"", -1},
	}

	for _, tc := range testCases {
		t.Run(tc.cpf, func(t *testing.T) {
			result := UnmaskToInt(tc.cpf)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func BenchmarkMask(b *testing.B) {
	for b.Loop() {
		Mask(11144477735)
	}
}

func BenchmarkWithMask(b *testing.B) {
	for b.Loop() {
		WithMask("SP")
	}
}

func BenchmarkUnmask(b *testing.B) {
	for b.Loop() {
		Unmask("111.444.777-35")
	}
}

func BenchmarkUnmaskToInt(b *testing.B) {
	for b.Loop() {
		UnmaskToInt("111.444.777-35")
	}
}
