# Adaptive Accrual Failure Detector

[![Documentation](https://godoc.org/github.com/andy2046/failured?status.svg)](http://godoc.org/github.com/andy2046/failured)
[![GitHub issues](https://img.shields.io/github/issues/andy2046/failured.svg)](https://github.com/andy2046/failured/issues)
[![license](https://img.shields.io/github/license/andy2046/failured.svg)](https://github.com/andy2046/failured/LICENSE)
[![Release](https://img.shields.io/github/release/andy2046/failured.svg?label=Release)](https://github.com/andy2046/failured/releases)

----

## There is NO perfect failure detector.

## It's a trade-off between completeness and accuracy.

## Failure detection is not a binary value.

This is an implementation of a failure detector that uses
an adaptive accrual algorithm. The theory of this failure detector is taken from the paper
[A New Adaptive Accrual Failure Detector for Dependable Distributed Systems](https://dl.acm.org/citation.cfm?id=1244129).

This failure detector is useful for detecting connections failures between nodes in distributed systems.

for documentation, view the [API reference](./doc.md)


## Install

```
go get github.com/andy2046/failured
```

## Usage

```go
package main

import (
	"time"

	"github.com/andy2046/failured"
)

func main() {
	fd := failured.New()
	closer := make(chan struct{})

	// call RegisterHeartbeat every second
	go func() {
		for {
			select {
			case <-closer:
				return
			default:
			}

			time.Sleep(time.Second)
			fd.RegisterHeartbeat()
		}
	}()

	time.Sleep(3 * time.Second)
	close(closer)

	// check FailureProbability
	p := fd.FailureProbability()
	println("failure probability is", p)
}
```
