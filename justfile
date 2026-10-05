default:
  @just --list

run:
  gofmt -w -s .
  go run .

