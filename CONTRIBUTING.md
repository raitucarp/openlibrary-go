# Contributing to openlibrary-go

Thank you for considering contributing!

## Development Setup

```bash
git clone https://github.com/raitucarp/openlibrary-go
cd openlibrary-go
go mod tidy
```

## Running Tests

Tests are located in `tests/`:

```bash
go test ./...
```

## Code Style

- Follow standard Go formatting (`go fmt ./...`)
- Public types and methods should have GoDoc comments
- Add runnable usage examples in `examples_test.go` when adding new API features

## Pull Requests

1. Create a feature branch: `git checkout -b feature/my-change`
2. Commit using clear messages
3. Ensure tests pass: `go test ./...`
4. Open Pull Request against `main`

## Reporting Issues

Open an issue at:
https://github.com/raitucarp/openlibrary-go/issues
