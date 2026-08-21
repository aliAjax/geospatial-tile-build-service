package health

import "sync/atomic"

type State struct{ ready atomic.Bool }

func (s *State) SetReady(v bool) { s.ready.Store(v) }
func (s *State) Ready() bool     { return s.ready.Load() }
