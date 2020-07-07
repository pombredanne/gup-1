# gup

Golang binary updater!

`gup` takes a go binary and uses the symbol table in the executable to
find out where its from and then updates it in case it is out of date.

This is useful for small CLI tools that people make that may become
out of date.


## Usage

```
gup [binary ...]

gup updates go binaries. It finds the go import path of the binary and does 'go get -u' and 'go install' for that path if found.
```
