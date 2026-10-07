default:
  @just --list

source := "./cmd/badges"

format:
  gofmt -w -s {{ source }}

run: clean format
  go run {{ source }}

clean:
  rm -rf docs
