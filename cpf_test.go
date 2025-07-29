package cpf

import "testing"

func TestGenerate(t *testing.T) {
	type testCase struct {
		state     string
		expectErr bool
	}

	testCases := []testCase{
		{"RS", false},
		{"DF", false},
		{"GO", false},
		{"MS", false},
		{"MT", false},
		{"TO", false},
		{"AC", false},
		{"AM", false},
		{"AP", false},
		{"PA", false},
		{"RO", false},
		{"RR", false},
		{"CE", false},
		{"MA", false},
		{"PI", false},
		{"AL", false},
		{"PB", false},
		{"PE", false},
		{"RN", false},
		{"BA", false},
		{"SE", false},
		{"MG", false},
		{"ES", false},
		{"RJ", false},
		{"SP", false},
		{"PR", false},
		{"SC", false},
		{"XX", true},    // Invalid state
		{"", true},      // Empty state
		{"SP123", true}, // Invalid state format
		{"!@#$", true},  // Invalid characters
	}

	for _, tc := range testCases {
		t.Run(tc.state, func(t *testing.T) {
			cpf := Generate(tc.state)

			if !tc.expectErr {
				if cpf < 0 {
					t.Errorf("Expected valid CPF for state %s, got error code %d", tc.state, cpf)
				}
			} else {
				if cpf >= 0 {
					t.Errorf("Expected error for state %s, got valid CPF %d", tc.state, cpf)
				}
			}
		})
	}
}

func BenchmarkGenerate(b *testing.B) {
	for b.Loop() {
		Generate("SP")
	}
}
