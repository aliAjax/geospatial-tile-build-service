package domain

import (
	"fmt"
	"strings"
)

var channels = map[string]bool{"stable": true, "canary": true, "preview": true}

func ValidateChannel(v string) error {
	v = strings.TrimSpace(strings.ToLower(v))
	if !channels[v] {
		return fmt.Errorf("unsupported channel %s", v)
	}
	return nil
}
func IsActive(r Release) bool { return r.Active && !r.UpdatedAt.IsZero() }
