# Ultimate Tic Tac Toe next best move prediction

## Why?

To find joy of coding. With AI taking over a big chunk of my development flow, I was missing the lightbulb moments which I used to get before AI got involved.

This project is an attempt to bring back that joy for me.

## What I got?

- The joy of coding!
- Did bit manipulation like never before
- MiniMax algorithm with Alpha-Beta pruning
- Using 2x2 uint64 to get the board
- Extreme memory and compute optimization

## Run benchmark and test

```bash
go test ./... -bench=. -count=5
```

Results on my machine -

```bash
goos: darwin
goarch: arm64
pkg: github.com/SirusCodes/9x9-analysis/engine
cpu: Apple M4
BenchmarkRunEngine_Depth10-10    	       1	1667939000 ns/op	5871642320 B/op	183482083 allocs/op
BenchmarkRunEngine_Depth10-10    	       1	1670770958 ns/op	5871584528 B/op	183481741 allocs/op
BenchmarkRunEngine_Depth10-10    	       1	1681113458 ns/op	5871569616 B/op	183481613 allocs/op
BenchmarkRunEngine_Depth10-10    	       1	1692574625 ns/op	5871549848 B/op	183481512 allocs/op
BenchmarkRunEngine_Depth10-10    	       1	1691782792 ns/op	5871584096 B/op	183481690 allocs/op
```
