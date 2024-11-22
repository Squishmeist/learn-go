# Learn Go

This project follows the modules from [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration).

## Setup

To get started, install `godoc`:

```sh
go install golang.org/x/tools/cmd/godoc@latest
```

## Commands

### Run Tests

To run all tests, use:

```sh
go test ./...
```

### Run Benchmarks

To run benchmarks, use:

```sh
go test -bench=.
```

### View Documentation

To view the documentation locally, run:

```sh
godoc -http=localhost:8000
```

Then open your browser and navigate to `http://localhost:8000`.
