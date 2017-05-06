# c-go-benchmarks

## Description

This repository contains some benchmarks I ran to study the possibility of using Go's concurrency to parallely run procedural C code. Most of the benchmarks out there ([example](https://benchmarksgame.alioth.debian.org/u64q/compare.php?lang=go&lang2=gcc)) show that C code can _sometimes_ be several times faster than Go code. That said, writing safe concurrent code in Go is much easier, so if we could mix the two it would be a great success for concurrent programs that need to go as fast as possible.

The question is: can we find any real improvements in example cases? It's perfectly possible that that isn't the case, for several different reasons (for example, there may be some optimizations that the go compiler manages to do on its concurrent code but that can't be done when some of the code is written in C). This is what these benchmarks are made to test.