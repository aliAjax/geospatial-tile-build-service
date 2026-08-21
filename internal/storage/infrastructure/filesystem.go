package infrastructure

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileStore struct{ Root string }

func NewFileStore(root string) *FileStore { return &FileStore{Root: root} }
func (f *FileStore) path(k string) string { return filepath.Join(f.Root, filepath.Clean("/"+k)) }
func (f *FileStore) Put(ctx context.Context, k string, b []byte) error {
	_ = ctx
	p := f.path(k)
	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}
	tmp := p + ".tmp"
	if err := os.WriteFile(tmp, b, 0644); err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return os.Rename(tmp, p)
}
func (f *FileStore) Open(ctx context.Context, k string) (io.ReadCloser, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	r, err := os.Open(f.path(k))
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return r, nil
}
func (f *FileStore) Delete(ctx context.Context, k string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	if err := os.Remove(f.path(k)); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}
