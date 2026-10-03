# go-ttml
A library for parsing and accessing TTML files in a somewhat more human-welcoming manner, written for Go.

## Installing

```
go get github.com/Endg4meZer0/go-ttml
```

## Usage

As of right now, the only thing this library can offer is `ParseString(string)`. This will try to parse the given string and, if successful, give a TTML struct back. More about TTML object structure in `type_*.go` files.

## Updates?

Someday. I have plans to make more options for parsing and also write an option to marshal lyrics into TTML as well, mostly to use in [lrcsnc](https://github.com/Endg4meZer0/lrcsnc). But yeah, the point stands - someday. For now it's mostly a proof of concept.