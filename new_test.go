package cpf

import "testing"

func TestNew(t *testing.T) {
	type testCase struct {
		state    string
		expected bool
	}

	testCases := []testCase{
		{"RS", true},
		{"DF", true},
		{"GO", true},
		{"MS", true},
		{"MT", true},
		{"TO", true},
		{"AC", true},
		{"AM", true},
		{"AP", true},
		{"PA", true},
		{"RO", true},
		{"RR", true},
		{"CE", true},
		{"MA", true},
		{"PI", true},
		{"AL", true},
		{"PB", true},
		{"PE", true},
		{"RN", true},
		{"BA", true},
		{"SE", true},
		{"MG", true},
		{"ES", true},
		{"RJ", true},
		{"SP", true},
		{"PR", true},
		{"SC", true},
		{"XX", false},    // Invalid state
		{"", false},      // Empty state
		{"SP123", false}, // Invalid state format
		{"!@#$", false},  // Invalid characters
	}

	for _, tc := range testCases {
		t.Run(tc.state, func(t *testing.T) {
			cpf := New(tc.state)

			got := IsValid(cpf)
			if got != tc.expected {
				t.Errorf("Expected CPF validity for state %s to be %v, got %v", tc.state, tc.expected, got)
			}

		})
	}
}

func BenchmarkGenerate(b *testing.B) {
	for b.Loop() {
		New("SP")
	}
}
