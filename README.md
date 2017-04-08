# DMY

Command line interface to create dummy data.

## Install

```bash
$ go get github.com/tzmfreedom/dmy
```

## Usage

```bash
$ dmy -N 10 hoge "foo_{{.Index}}" {{date .Index}}
```