# [cpf](https://github.com/klasrak/cpf): High-Performance Brazilian CPF Utilities 🚀

[![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen.svg)](https://github.com/klasrak/cpf/actions)
[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://pkg.go.dev/github.com/klasrak/cpf)
[![GitHub Stars](https://img.shields.io/github/stars/klasrak/cpf.svg?style=social)](https://github.com/klasrak/cpf)

Welcome to **[cpf](https://github.com/klasrak/cpf)**, a *blazing-fast*, zero-allocation (where possible) Go package designed for seamless handling of Brazilian CPF (Cadastro de Pessoas Físicas) numbers. Whether you're generating valid CPFs tied to specific states, validating existing ones, or toggling between masked and unmasked formats, this library prioritizes **performance**, **simplicity**, and **reliability**. Built with efficiency in mind, it avoids unnecessary allocations and executes operations in nanoseconds—perfect for high-throughput applications like APIs, data processing pipelines, or compliance tools.

cpf leverages Go's strengths for optimal speed. It's open-source under the MIT license, so feel free to fork, contribute, or integrate it into your projects!

## 🌟 Why cpf?

- **Ultra-Fast Operations**: Benchmarks show operations completing in as little as 1ns/op with zero allocations for validation.
- **State-Specific Generation**: Generate CPFs linked to Brazilian states using the official modulo 11 algorithm.
- **Flexible Input/Output**: Supports ints, uints, strings, and more for seamless integration.
- **Masking & Unmasking**: Easily apply or remove the standard "000.000.000-00" format.
- **Validation Excellence**: Thorough checks for structure, check digits, and common invalid patterns (e.g., all zeros).
- **No Dependencies**: Pure Go, lightweight, and ready to go.
- **Creative Edge**: We even handle edge cases like oversized uint64 inputs by formatting from right to left!

CPF stands for "Cadastro de Pessoas Físicas," Brazil's unique taxpayer ID. It's an 11-digit number with two check digits calculated via modulo 11. This package demystifies it all while keeping things performant.

## 📦 Installation

Get started in seconds:

```bash
go get -u github.com/klasrak/cpf
```

Add it to your `go.mod` and import as `import "github.com/klasrak/cpf"`.

## 🛠️ Quick Start & Examples

Dive right in with these snippets. Assume you've imported the package.

### Generating a Valid CPF

```go
package main

import (
    "fmt"
    "github.com/klasrak/cpf"
)

func main() {
    // Generate a numeric CPF for São Paulo
    cpfNum := cpf.New("SP")
    if cpfNum == -1 {
        fmt.Println("Invalid state!")
    } else {
        fmt.Printf("Generated CPF: %d\n", cpfNum) // e.g., 12345678901
    }

    // Generate a masked CPF for Rio de Janeiro
    maskedCPF := cpf.WithMask("RJ")
    fmt.Println("Masked CPF:", maskedCPF) // e.g., "123.456.789-01"
}
```

### Validating a CPF

```go
isValid := cpf.IsValid("123.456.789-09")
fmt.Println("Is Valid?", isValid) // false (invalid check digits)

isValidNum := cpf.IsValid(98765432100)
fmt.Println("Is Valid?", isValidNum) // true (if valid)
```

### Masking & Unmasking

```go
// Mask a raw number
masked := cpf.Mask(11144477735)
fmt.Println(masked) // "111.444.777-35"

// Unmask a formatted string
unmasked := cpf.Unmask("111.444.777-35")
fmt.Println(unmasked) // "11144477735"

// Unmask to int
unmaskedInt := cpf.UnmaskToInt("111.444.777-35")
fmt.Println(unmaskedInt) // 11144477735
```

For more examples, check the [examples directory](examples/) in the repo.

## 🔍 Deep Dive: How CPF Generation Works

Brazilian CPFs are 11 digits: ABC.DEF.GHI-JK, where J and K are check digits.

1. **Base Digits**: Randomly generate 9 digits (or use state-specific logic).
2. **First Check Digit (J)**:
   - Weights: 10 to 2 (left to right).
   - Sum products, modulo 11.
   - If remainder < 2, J=0; else J=11-remainder.
3. **Second Check Digit (K)**:
   - Include J, weights 11 to 2.
   - Same modulo logic.
4. **State Binding**: The 9th digit ties to a state (e.g., 8 for SP).

Invalid states return -1 or empty strings.

## 📚 API Reference

All functions are in the `cpf` package. Here's a detailed breakdown:

### `New(state string) int`

- **Description**: Generates a valid CPF for the state (e.g., "SP" → São Paulo).
- **Returns**: Valid CPF as int, or -1 on invalid state.
- **Example**: `cpf.New("MG") → 12345678901`

### `WithMask(state string) string`

- **Description**: Like `New`, but returns masked string (e.g., "123.456.789-01").
- **Returns**: Masked CPF or "" on invalid state.

### `Mask(cpf any) string`

- **Description**: Formats input to "000.000.000-00".
- **Inputs**: int, int32/64, uint/32/64, string.
- **Returns**: Masked string or "" on invalid.
- **Note**: Handles large uint64 by right-aligning digits.

### `Unmask(cpf string) string`

- **Description**: Strips dots, dashes; returns plain digits.
- **Returns**: "00000000000" or "" on invalid.

### `UnmaskToInt(cpf string) int`

- **Description**: Unmasks and converts to int.
- **Returns**: Positive int or -1 on invalid.

### `IsValid(cpf any) bool`

- **Description**: Validates structure and check digits.
- **Inputs**: int/64, uint/64, string (masked/unmasked).
- **Returns**: true if valid.


## ⚡ Performance Benchmarks

Optimized for speed and low overhead. This is my benchmarks on a 13th Gen Intel® Core™ i5-13600K (Linux, amd64). We minimized allocations—many ops are zero-alloc!

| Benchmark          | Time (ns/op) | Alloc (B/op) | Allocs/op |
|--------------------|--------------|--------------|-----------|
| Mask              | 24.6        | 16          | 1        |
| WithMask          | 62.6        | 16          | 1        |
| Unmask            | 24.8        | 16          | 1        |
| UnmaskToInt       | 12.7        | 0           | 0        |
| New               | 30.1        | 0           | 0        |
| IsValidInt        | 16.8        | 0           | 0        |
| IsValidUint       | 1.0         | 0           | 0        |
| IsValidUint64     | 16.8        | 0           | 0        |
| IsValidInt64      | 16.8        | 0           | 0        |
| IsValidString     | 11.5        | 0           | 0        |
| IsValidWithMask   | 11.9        | 0           | 0        |

Run `go test -bench=.` to verify on your machine. These results highlight our focus on performance: **validation is sub-20ns and allocation-free!**

### Full benchmark

```sh
go test -benchmem -cpu 1,2,4,8 -bench=.

goos: linux
goarch: amd64
pkg: github.com/klasrak/cpf
cpu: 13th Gen Intel(R) Core(TM) i5-13600K
```

| Benchmark                | Iterations   | Time (ns/op) | Alloc (B/op) | Allocs/op |
|--------------------------|--------------|--------------|--------------|-----------|
| BenchmarkMask            | 48442154     | 24.72        | 16           | 1         |
| BenchmarkMask-2          | 46789185     | 24.42        | 16           | 1         |
| BenchmarkMask-4          | 48796629     | 24.71        | 16           | 1         |
| BenchmarkMask-8          | 48979684     | 24.61        | 16           | 1         |
| BenchmarkWithMask        | 19398380     | 59.78        | 16           | 1         |
| BenchmarkWithMask-2      | 19855608     | 59.63        | 16           | 1         |
| BenchmarkWithMask-4      | 19924069     | 59.63        | 16           | 1         |
| BenchmarkWithMask-8      | 20157704     | 59.79        | 16           | 1         |
| BenchmarkUnmask          | 48586246     | 24.85        | 16           | 1         |
| BenchmarkUnmask-2        | 49273590     | 24.41        | 16           | 1         |
| BenchmarkUnmask-4        | 47833567     | 24.72        | 16           | 1         |
| BenchmarkUnmask-8        | 48135663     | 24.84        | 16           | 1         |
| BenchmarkUnmaskToInt     | 93930397     | 12.56        | 0            | 0         |
| BenchmarkUnmaskToInt-2   | 94553982     | 12.45        | 0            | 0         |
| BenchmarkUnmaskToInt-4   | 96315583     | 12.44        | 0            | 0         |
| BenchmarkUnmaskToInt-8   | 93799954     | 12.54        | 0            | 0         |
| BenchmarkNew             | 44703962     | 26.97        | 0            | 0         |
| BenchmarkNew-2           | 43534274     | 26.90        | 0            | 0         |
| BenchmarkNew-4           | 44517864     | 26.94        | 0            | 0         |
| BenchmarkNew-8           | 44245804     | 26.91        | 0            | 0         |
| BenchmarkIsValidInt      | 71274181     | 16.54        | 0            | 0         |
| BenchmarkIsValidInt-2    | 72173035     | 16.61        | 0            | 0         |
| BenchmarkIsValidInt-4    | 72483140     | 16.54        | 0            | 0         |
| BenchmarkIsValidInt-8    | 72489699     | 16.52        | 0            | 0         |
| BenchmarkIsValidUint     | 1000000000   | 1.000        | 0            | 0         |
| BenchmarkIsValidUint-2   | 1000000000   | 1.000        | 0            | 0         |
| BenchmarkIsValidUint-4   | 1000000000   | 1.000        | 0            | 0         |
| BenchmarkIsValidUint-8   | 1000000000   | 1.000        | 0            | 0         |
| BenchmarkIsValidUint64   | 72778605     | 16.53        | 0            | 0         |
| BenchmarkIsValidUint64-2 | 70614046     | 16.56        | 0            | 0         |
| BenchmarkIsValidUint64-4 | 70101028     | 16.58        | 0            | 0         |
| BenchmarkIsValidUint64-8 | 70111623     | 16.55        | 0            | 0         |
| BenchmarkIsValidInt64    | 71954649     | 16.63        | 0            | 0         |
| BenchmarkIsValidInt64-2  | 72113073     | 16.62        | 0            | 0         |
| BenchmarkIsValidInt64-4  | 71014800     | 16.64        | 0            | 0         |
| BenchmarkIsValidInt64-8  | 68481709     | 16.66        | 0            | 0         |
| BenchmarkIsValidString   | 100000000    | 11.35        | 0            | 0         |
| BenchmarkIsValidString-2 | 100000000    | 11.33        | 0            | 0         |
| BenchmarkIsValidString-4 | 100000000    | 11.34        | 0            | 0         |
| BenchmarkIsValidString-8 | 100000000    | 11.36        | 0            | 0         |
| BenchmarkIsValidWithMask | 100000000    | 11.98        | 0            | 0         |
| BenchmarkIsValidWithMask-2 | 100000000  | 11.99        | 0            | 0         |
| BenchmarkIsValidWithMask-4 | 100000000  | 12.03        | 0            | 0         |
| BenchmarkIsValidWithMask-8 | 94558876   | 11.99        | 0            | 0         |

## 🗺️ Supported States

All 26 Brazilian states plus Brasilia (capital) via two-letter codes:
- AC, AL, AP, AM, BA, CE, DF, ES, GO, MA, MT, MS, MG, PA, PB, PR, PE, PI, RJ, RN, RS, RO, RR, SC, SP, SE, TO.


## 🤝 Contributing

Fork the repo, do your things and be happy <3.

## 📜 License

MIT License—free to use, modify, and distribute. See [LICENSE](LICENSE) for full text.

## 🎉 Acknowledgments

- Thanks to the Go community for inspiration.
- Built with ❤️ by [klasrak](https://github.com/klasrak). Star the repo if it helps!

For questions, open an issue. Happy coding! 🇧🇷