default:
  @just --list

source := "./cmd/badges"
run:
  gofmt -w -s {{ source }}
  go run {{ source }}

clean:
  rm -rf docs
