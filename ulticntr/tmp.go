package main

import (
	"io"
	"sync"
)

type tmpCounter struct {
	sync.Mutex
	internal io.ReadWriteCloser
}

func newTmpCounter() (*tmpCounter, error) {
	return &tmpCounter{}, nil
}

func (t tmpCounter) Count() (uint64, error) { return 5, nil }
