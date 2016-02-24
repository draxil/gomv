# gomv

Move files in go without caring about whether it's possible to do this as a reanme (with os.Rename) or if we have to copy-and-then-delete.

Ex:

> import "github.com/draxil/gomv"

> ....

> err := gomv.MoveFile(filea, fileb)

# details

Currently all this does is to attempt an os.Rename, and if that fails with a cross device error it'll attempt a "copy and them remove" style move. At time of writing this is very fresh from scratching the itch so any issues please let me know.

So far I have tested this moving accross devices in Linux and OSX. To really get the best out of the test suite currently (actually test the key copy-and-move case) you need to set an envrionment variable to somewhere that will proc the error.

# TODO

* some better way to simulate the errors for a better default test
* Windows..
* More errors which would block a rename.

---
[![GoDoc](https://godoc.org/github.com/draxil/gomv?status.svg)](https://godoc.org/github.com/draxil/gomv)
[![Build Status](https://travis-ci.org/draxil/gomv.svg?branch=master)](https://travis-ci.org/draxil/gomv)
