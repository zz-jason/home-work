## Introduction

This is a homework. It's an aggregate evaluator, you need to:
1. add an implementation of the `Aggregate` interface to calculate the result of
   `select a, avg(b) from t group by a;`.

2. add a bench mark test in the file `bench_test.go` like the
   `BenchmarkNaiveAggregate()` function to benchmark the newly added
   implementation by you.

There are two kinds of readers can be used to complete this task:

1. `TypeColumn` returns a `[][]int64`, data belongs to the same column are
   continuously stored in the `[]int64` slice, there are only two columns to
   return, the first column represent column `a`, and the second represent
   column `b`.

2. `TypeRow` returns a `[][]int64`, data belongs to the same row are stored
   continuously in the `[]int64` slice, there are two columns in one row.  Like
   `TypeColumn`, the first column represent column `a`, and the second represent
   column `b`.

Both `TypeColumn` and `TypeRow` returns a batch of data, the only difference is
the memory layout of the returned data.

You only need to complete one of the following benchmarks, please chose the one
that takes the less execution time:

The first one uses the row-oriented memory reader:
```go
func BenchmarkYourAggregate(b *testing.B) {
	opReader := memreader.Build(memreader.TypeRow)
	opReader.Populate()

	opAgg := new(YourAggregate)
	opAgg.SetReader(opReader)

	b.ResetTimer()
	for counter := 0; counter < b.N; counter++ {
		opAgg.GetResult()
	}
}
```

The second one uses the column-oriented memory reader:
```go
func BenchmarkYourAggregate(b *testing.B) {
	opReader := memreader.Build(memreader.TypeColumn)
	opReader.Populate()

	opAgg := new(YourAggregate)
	opAgg.SetReader(opReader)

	b.ResetTimer()
	for counter := 0; counter < b.N; counter++ {
		opAgg.GetResult()
	}
}
```

**NOTE**:
1. The only difference of the first and second one is the mem-reader.
2. You only need to consider **one** of the readers, please chose it carefully.
3. The goal is to obtain the **lowest** execution time, the faster the better.
