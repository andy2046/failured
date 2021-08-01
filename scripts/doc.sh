#!/usr/bin/env bash

set -euo pipefail

godoc2md github.com/andy2046/failured \
  > $GOPATH/src/github.com/andy2046/failured/doc.md
