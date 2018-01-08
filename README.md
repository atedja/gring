# gring

[![GoDoc](https://godoc.org/github.com/atedja/gring?status.svg)](https://godoc.org/github.com/atedja/gring) [![Build Status](https://travis-ci.org/atedja/gring.svg?branch=master)](https://travis-ci.org/atedja/gring)

Circular Linked List implemented with array backend.
Compared to `container/ring`, `gring` is using array as backend and integer indexes as pointers to nodes in the ring to make it easier to swap/remove/insert nodes without having to manage node pointers.
In addition, the added benefits to having integer indexes are simpler cloning and marshalling.
