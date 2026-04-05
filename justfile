set dotenv-load

[group('meta')]
default:
  @just --list --unsorted

[group('app')]
run:
  go run ./cmd/server

[group('quality')]
lint:
  golangci-lint run

[group('quality')]
lint-fix:
  golangci-lint run --fix

[group('quality')]
format:
  golangci-lint fmt

[group('migrations')]
goose-install:
  go install github.com/pressly/goose/v3/cmd/goose@latest

[group('migrations')]
goose-status:
  goose status

[group('migrations')]
goose-up:
  goose up

[group('migrations')]
goose-down:
  goose down

[group('migrations')]
goose-create name:
  goose create {{name}} sql
