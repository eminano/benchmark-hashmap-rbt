# Benchmark hash map vs red-black tree

## Results

```
goos: darwin
goarch: amd64
pkg: benchmark-rbt-hashmap
BenchmarkInsert_hashmap-4   	 2113599	       572 ns/op	      97 B/op	       2 allocs/op
BenchmarkInsert_rbt-4       	 1741177	       734 ns/op	      96 B/op	       4 allocs/op
BenchmarkGet_hashmap-4      	 7684614	       157 ns/op	      15 B/op	       1 allocs/op
BenchmarkGet_rbt-4          	 5560808	       217 ns/op	      31 B/op	       2 allocs/op
BenchmarkGetAll_hashmap-4   	   74847	     16382 ns/op	    8843 B/op	       0 allocs/op
BenchmarkGetAll_rbt-4       	   66264	     17622 ns/op	   11780 B/op	       1 allocs/op
```
