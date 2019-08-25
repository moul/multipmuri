# multipmuri

> A library to parse project-management URIs (inspired by @multiformats)

[![CircleCI](https://circleci.com/gh/moul/multipmuri.svg?style=shield)](https://circleci.com/gh/moul/multipmuri)
[![GoDoc](https://godoc.org/moul.io/multipmuri?status.svg)](https://godoc.org/moul.io/multipmuri)
[![License](https://img.shields.io/github/license/moul/multipmuri.svg)](https://github.com/moul/multipmuri/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/moul/multipmuri.svg)](https://github.com/moul/multipmuri/releases)
[![Go Report Card](https://goreportcard.com/badge/moul.io/multipmuri)](https://goreportcard.com/report/moul.io/multipmuri)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/multipmuri/badge)](https://www.codefactor.io/repository/github/moul/multipmuri)
[![codecov](https://codecov.io/gh/moul/multipmuri/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/multipmuri)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)
<!--[![Docker Metrics](https://images.microbadger.com/badges/image/moul/multipmuri.svg)](https://microbadger.com/images/moul/multipmuri)-->


## Usage

This libraries parses an URI that you could find in a README, markdown, comment field of the populat project-management tools.

You can use the library to parse an URI:
* without context (i.e. `github.com/moul/depviz#42`)
* with a context (i.e. `#42` in the context of `github.com/moul/depviz`)

```golang
import "moul.io/multipmuri"

depviz42, _ := multipmuri.DecodeString("github.com/moul/depviz#42")
fmt.Println(depviz42) // https://github.com/moul/depviz/issues/42

depviz43, _ := depviz42.RelDecodeString("#43")
fmt.Println(depviz43) // https://github.com/moul/depviz/issues/43
```

## Install

```console
$ go get -u moul.io/multipmuri
```

## License

© 2019 [Manfred Touron](https://manfred.life) -
[Apache-2.0 License](https://github.com/moul/multipmuri/blob/master/LICENSE)
