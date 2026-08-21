package domain

import "testing"

func TestChunkOwnsPlanStorage(t *testing.T) {
	items := []TilePlan{{Zoom: 1, Tiles: 1}, {Zoom: 2, Tiles: 2}}
	chunks := Chunk(items, 1)
	items[0].Zoom = 9
	if chunks[0][0].Zoom != 1 {
		t.Fatal("chunk aliases caller plan")
	}
}
