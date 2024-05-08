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
% ./easy-first-parser-go train --train-filename path/to/train.txt --dev-filename path/to/dev.txt --max-iter 10 --model-filename model.bin
0, 0.907, 0.893
1, 0.920, 0.901
2, 0.929, 0.904
3, 0.935, 0.906
4, 0.940, 0.907
5, 0.944, 0.907
6, 0.947, 0.908
7, 0.950, 0.908
8, 0.953, 0.908
9, 0.955, 0.908
```

## Evaluating a trained parser
To see the detail 