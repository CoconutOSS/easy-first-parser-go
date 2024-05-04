# easy-first-parser-go
[![CircleCI](https://circleci.com/gh/CoconutOSS/easy-first-parser-go.svg?style=shield)](https://circleci.com/gh/CoconutOSS/easy-first-parser-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/CoconutOSS/easy-first-parser-go)](https://goreportcard.com/report/github.com/CoconutOSS/easy-first-parser-go)
[![Coverage Status](https://coveralls.io/repos/github/CoconutOSS/easy-first-parser-go/badge.svg?branch=coveralls)](https://coveralls.io/github/CoconutOSS/easy-first-parser-go?branch=coveralls)

easy-first-parser-go - An efficient open-source Dependency Parser utilizing the Easy-First Algorithm written in Go.

# Build from source

```sh
% git clone https://github.com/CoconutOSS/easy-first-parser-go.git
% cd easy-first-parser-go
% make deps && make bindata && make build
```

# Usage
easy-first-parser-go has `train` (training a parser phase) and `eval` (evaluating a trained parser phase) modes. To see the detail options, type `./easy-first-parser-go --help`.

## Training a parser
To see the detail options, type `./easy-first-parser-go train --help`.

```sh
% ./easy-first-parser-go train --train-filename path/to/train.txt --dev-filename path/to/dev.txt --max-iter 10 --model-filename mo