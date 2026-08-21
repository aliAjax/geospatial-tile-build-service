package application

import (
	"context"
	"sync"
	"github.com/example/geospatial-tile-build-service/internal/storage/infrastructure"
	"testing"
)

func TestBuildStartKeepsCallerInput(t *testing.T) {
	input := make([]byte, 3, 64)
	copy(input, []byte("abc"))
	s := New(infrastructure.New()); start:=make(chan struct{}); var wg sync.WaitGroup
	for i:=0;i<2;i++{wg.Add(1);go func(id string){defer wg.Done();<-start;if _,err:=s.Start(context.Background(),id,"v1",input);err!=nil{t.Error(err)}}(string(rune('a'+i)))}
	close(start);wg.Wait()
	for _, b := range input[3:cap(input)] {
		if b != 0 {
			t.Fatalf("caller buffer was mutated: %v", input[:cap(input)])
		}
	}
}

func TestCacheKeepsDataOwnership(t *testing.T) {
	c := NewCache(2)
	src := []byte("cache-value")
	c.Put("k",src); src[0]='X'; start:=make(chan struct{});var wg sync.WaitGroup
	for i:=0;i<2;i++{wg.Add(1);go func(){defer wg.Done();<-start;for j:=0;j<50;j++{_,_=c.Get("k")}}()};close(start);wg.Wait()
	got, ok := c.Get("k")
	if !ok || string(got) != "cache-value" {
		t.Fatalf("cache retained caller reference: %q", got)
	}
	got[0] = 'Y'
	again, _ := c.Get("k")
	if string(again) != "cache-value" {
		t.Fatalf("cache returned internal reference: %q", again)
	}
}
