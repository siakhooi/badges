#!/usr/bin/env bash

set -euo pipefail

diff="$(gofmt -d -s .)"
if [ -n "$diff" ]; then
  printf '%s\n' "$diff"
  exit 1
fi

