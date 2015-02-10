# goigor

`goigor` generates `.gitignore` (or `.hgignore`) files from the command line. a bit like [joe](https://github.com/karan/joe) but written in [Go](http://golang.org/).

## features
* easy to install
* even easier to use
* works wherever [Go works](http://golang.org/doc/install#requirements)
* generates files for Git and Mercurial

## installation

```shell
go get github.com/mlen108/goigor
cd goigor
go build goigor.go
```

## usage

```shell
./goigor go   # output the Golang file to stdout
```

#### multiple languages

```shell
./goigor python java fortran
```

#### save to file

```shell
./goigor go > .gitignore
```
or
```shell
./goigor go > .hgignore
```

#### append to file

```shell
./goigor go >> .gitignore
```

## license

MIT
