package domain

import (
	"context"
	"io"
)

func ReadAllAndClose(r io.ReadCloser) ([]byte, error) {
	if r == nil {
		return nil, io.ErrUnexpectedEOF
	}
	return io.ReadAll(r)
}

type Store interface {
	Put(context.Context, string, []byte) error
	Open(context.Context, string) (io.ReadCloser, error)
	Delete(context.Context, string) error
}
