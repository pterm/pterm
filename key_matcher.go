package pterm

import (
	"atomicgo.dev/keyboard/keys"
)

type KeyMatcher struct {
	value keys.Key
}

// checks if a KeyMatcher is equal to a key (KeyCode or string)
// by checking key type and extracting KayMatcher for the given type
func (s KeyMatcher) Equal(key any) bool {
	switch k := key.(type) {
	case string:
		return s.value.String() == k
	case keys.KeyCode:
		return s.value.Code == k
	default:
		return false
	}
}
