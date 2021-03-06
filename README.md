# gring

[![GoDoc](https://godoc.org/github.com/atedja/gring?status.svg)](https://godoc.org/github.com/atedja/gring) [![Build Status](https://travis-ci.org/atedja/gring.svg?branch=master)](https://travis-ci.org/atedja/gring)

Circular Linked List implemented with array backend.
Compared to `container/ring`, `gring` is using array as backend and integer indexes as pointers to nodes to make it easier to swap/remove/insert nodes without having to manage node pointers.
In addition, the added benefits to having index pointers are simpler cloning and marshalling, and O(1) operation to get the size of the ring.

Features:
* Various operations: node insertion, detachment, swaps, 2-opt swaps, direction reversal
* Get size in O(1)
* JSON marshalling
* Cloneable
