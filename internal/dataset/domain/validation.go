package domain

import (
	"fmt"
	"regexp"
	"strings"
)

var idPattern = regexp.MustCompile(`^[a-z][a-z0-9_-]{1,63}$`)

func ValidateID(id string) error {
	if !idPattern.MatchString(id) {
		return fmt.Errorf("invalid identifier %q", id)
	}
	return nil
}
func ValidateCRS(crs string) error {
	crs = strings.ToUpper(strings.TrimSpace(crs))
	if crs != "EPSG:4326" && crs != "EPSG:3857" {
		return fmt.Errorf("unsupported crs %s", crs)
	}
	return nil
}
func ValidateLayers(layers []string) error {
	seen := map[string]bool{}
	for _, l := range layers {
		l = strings.TrimSpace(l)
		if l == "" {
			return fmt.Errorf("empty layer")
		}
		if seen[l] {
			return fmt.Errorf("duplicate layer %s", l)
		}
		seen[l] = true
	}
	if len(layers) > 128 {
		return fmt.Errorf("too many layers")
	}
	return nil
}
func (d Dataset) StrictValidate() error {
	if err := d.Validate(); err != nil {
		return err
	}
	if err := ValidateID(d.ID); err != nil {
		return err
	}
	if err := ValidateCRS(d.CRS); err != nil {
		return err
	}
	return ValidateLayers(d.Layers)
}
func NormalizeLayers(layers []string) []string {
	out := make([]string, 0, len(layers))
	seen := map[string]bool{}
	for _, l := range layers {
		l = strings.TrimSpace(strings.ToLower(l))
		if l != "" && !seen[l] {
			seen[l] = true
			out = append(out, l)
		}
	}
	return out
}
