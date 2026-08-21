package domain

import (
	"sync"
	"testing"
)

func TestOptionsCanonicalIsStableAndNonMutating(t *testing.T) {
	o := Options{Layers: []string{"roads", "buildings"}, MinZoom: 2, MaxZoom: 4}
	start:=make(chan struct{});var wg sync.WaitGroup
	for i:=0;i<2;i++{wg.Add(1);go func(){defer wg.Done();<-start;if got:=o.Canonical();got!="buildings,roads:2:4"{t.Errorf("unexpected canonical options: %s",got)}}()};close(start);wg.Wait()
	if o.Layers[0] != "roads" {
		t.Fatalf("canonicalization mutated caller layers: %#v", o.Layers)
	}
}
