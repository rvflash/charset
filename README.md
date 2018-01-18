# Charset

[![GoDoc](https://godoc.org/github.com/rvflash/charset?status.svg)](https://godoc.org/github.com/rvflash/charset)
[![Build Status](https://img.shields.io/travis/rvflash/charset.svg)](https://travis-ci.org/rvflash/charset)
[![Code Coverage](https://img.shields.io/codecov/c/github/rvflash/charset.svg)](http://codecov.io/github/rvflash/charset?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/rvflash/charset)](https://goreportcard.com/report/github.com/rvflash/charset)

Golang package to limit the number of bytes per character in a string.


### Why?

Because MySQL has multiple unicode support and behind the utf8 character set, it uses a maximum of three bytes per character.
This charset contains only BMP characters.
An other, named utf8mb4 uses a maximum of four bytes per character and supports supplementary characters.
This package provides method to limit the number of bytes per character in a utf8 string.

> https://dev.mysql.com/doc/refman/5.7/en/charset-unicode-utf8mb4.html

### Installation

```bash
$ go get -u github.com/rvflash/charset
```

### Usage

```go
import "github.com/rvflash/charset"
// ...
s := charset.FormatUtf8("Warning: ⚠ 1 minute left ⌚ before boom 💥 ➡ click on the button !")
fmt.Println(s)
// Output: Warning: ⚠ 1 minute left ⌚ before boom  ➡ click on the button !
```
