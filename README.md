```
go test -benchmem -cpu 1,2,4,8,16 -run=^$ -bench ^BenchmarkGenerate$ github.com/klasrak/cpf

goos: linux
goarch: amd64
pkg: github.com/klasrak/cpf
cpu: 13th Gen Intel(R) Core(TM) i5-13600K
BenchmarkGenerate               47688402                24.83 ns/op            0 B/op          0 allocs/op
BenchmarkGenerate-2             48220190                24.83 ns/op            0 B/op          0 allocs/op
BenchmarkGenerate-4             48244668                24.85 ns/op            0 B/op          0 allocs/op
BenchmarkGenerate-8             46726274                24.82 ns/op            0 B/op          0 allocs/op
BenchmarkGenerate-16            48149306                24.77 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/klasrak/cpf  5.939s
```