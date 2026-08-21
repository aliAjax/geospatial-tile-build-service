package infrastructure

import (
	"context"
	"sync"
	"testing"
)

func TestBuildInputSnapshotUnderConcurrentAccess(t *testing.T) {
	m := New()
	ctx := context.Background()
	want := []byte("tile-input-v1")
	if err := m.Put(ctx, "builds/a/input", want); err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 200; j++ {
				if err := m.Put(ctx, "builds/a/input", []byte("tile-input-v1")); err != nil {
					t.Errorf("put %d: %v", i, err)
				}
				r, err := m.Open(ctx, "builds/a/input")
				if err != nil {
					t.Errorf("open %d: %v", i, err)
					continue
				}
				buf := make([]byte, len(want))
				if _, err := r.Read(buf); err != nil {
					t.Errorf("read: %v", err)
				}
				_ = r.Close()
				if string(buf) != string(want) {
					t.Errorf("snapshot changed: %q", buf)
				}
			}
		}(i)
	}
	wg.Wait()
}
