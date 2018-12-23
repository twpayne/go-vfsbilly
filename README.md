# `go-vfsbilly`

[![GoDoc](https://godoc.org/github.com/twpayne/go-vfsbilly?status.svg)](https://godoc.org/github.com/twpayne/go-vfsbilly)
[![Build Status](https://travis-ci.org/twpayne/go-vfsbilly.svg?branch=master)](https://travis-ci.org/twpayne/go-vfsbilly)
[![Report Card](https://goreportcard.com/badge/github.com/twpayne/go-vfsbilly)](https://goreportcard.com/report/github.com/twpayne/go-vfsbilly)

Package `vfsbilly` provides a compatibility later between
[`github.com/twpayne/go-vfs`](https://github.com/twpayne/go-vfs) and
[`github.com/src-d/go-billy`](https://github.com/src-d/go-billy).

This allows you to use `vfst` to test exisiting code that uses
[`billy.Filesystem`](https://godoc.org/github.com/src-d/go-billy#Filesystem).
See [the documentation](https://godoc.org/github.com/twpayne/go-vfsbilly) for
an example.


## License

The MIT License (MIT)

Copyright (c) 2018 Tom Payne

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
