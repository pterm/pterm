package putils

import (
	"cmp"
	"fmt"
	"maps"
	"reflect"
	"slices"

	"github.com/pterm/pterm"
)

// KeyValueFromStruct turns the exported fields of a struct, or a pointer to
// one, into key value pairs in declaration order. Field tags work as in
// TableFromStructs. Any other value returns nil.
//
//	pterm.KeyValue.WithPairs(putils.KeyValueFromStruct(build)...).Render()
func KeyValueFromStruct(v any) []pterm.KeyValuePair {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Pointer && !rv.IsNil() {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return nil
	}

	fields := structFields(rv.Type())
	pairs := make([]pterm.KeyValuePair, len(fields))

	for i, f := range fields {
		pairs[i] = pterm.KeyValuePair{Key: f.name, Value: format(rv.Field(f.index))}
	}

	return pairs
}

// KeyValueFromMap turns a map into key value pairs sorted by key, so the
// output is the same on every run.
func KeyValueFromMap[K cmp.Ordered, V any](m map[K]V) []pterm.KeyValuePair {
	keys := slices.Sorted(maps.Keys(m))
	pairs := make([]pterm.KeyValuePair, len(keys))

	for i, k := range keys {
		pairs[i] = pterm.KV(fmt.Sprint(k), m[k])
	}

	return pairs
}
