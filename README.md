# gup

Golang binary updater!

`gup` takes a go binary and uses the symbol table in the executable to
find out where its from and then updates it in case it is out of date.

## Usage

```
gup [binary ...]

gup updates go binaries. It finds the go import path of the binary and does 'go get -u' and 'go install' for that path if found.
```

# Notes

This is useful for small CLI tools that people make that may become
out of date.

Uses code from [goman](https://github.com/christophberger/goman).

I intend to update this almost never since it should just work.

Maybe a cool feature to add would be an import in your go executables
that checks if it s out of date and lets you know to stderr.
