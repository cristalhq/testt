# testt

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Testing tools for Go.

## Rationale

TODO

## Features

* Simple API.

See [GUIDE.md](https://github.com/cristalhq/testt/blob/main/GUIDE.md) for more details

## Install

Go version 1.17+

```
go get github.com/cristalhq/testt
```

## Example

```go
func TestExample(t *testing.T) {
	tt := testt.New(t)
	var err error
	tt.NoError(err)
}
```

Also see examples: [examples_test.go](https://github.com/cristalhq/testt/blob/main/examples_test.go).

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/testt/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/testt/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/testt
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/testt
[reportcard-img]: https://goreportcard.com/badge/cristalhq/testt
[reportcard-url]: https://goreportcard.com/report/cristalhq/testt
[coverage-img]: https://codecov.io/gh/cristalhq/testt/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/testt
[version-img]: https://img.shields.io/github/v/release/cristalhq/testt
[version-url]: https://github.com/cristalhq/testt/releases
