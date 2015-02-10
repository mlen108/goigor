# goigor

`goigor` generates `.gitignore` (or `.hgignore` if you fancy it) files from the command line.

a bit like [joe](https://github.com/karan/joe) but written in [Go](http://golang.org/) (because `pip` has gone crazy once again..).

## features
* easy to install
* even easier to use
* works wherever [Go works](http://golang.org/doc/install#requirements)
* no stupid third-party dependencies
* generates files for Git and Mercurial

## installation

```bash
go get github.com/mlen108/goigor
cd goigor
go build goigor.go
```

## usage

```bash
./goigor go   # outputs the Golang file to stdout
```

#### multiple languages

```bash
./goigor python java fortran
```

#### save to file

```bash
./goigor go > .gitignore
```
or
```bash
./goigor go > .hgignore
```

#### append to file

```bash
./goigor go >> .gitignore
```

## license

MIT
