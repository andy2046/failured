#!/usr/bin/env bash

set -euo pipefail

go test -count=1 -v -cover -race
go test -bench=. -run=none -benchtime=3s
go fmt
go vet
golint
