package domain

import "testing"

func TestFeatureCloneOwnsMutableData(t *testing.T) {
	original := Feature{Properties: map[string]any{"name": "road"}, Geometry: LineString{{X: 1, Y: 1}, {X: 2, Y: 2}}}
	clone := original.Clone()
	clone.Properties["name"] = "changed"
	clone.Geometry.(LineString)[0].X = 9
	if original.Properties["name"] != "road" || original.Geometry.(LineString)[0].X != 1 {
		t.Fatal("clone aliases original")
	}
}
