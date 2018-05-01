package dumbhashmap

import "fmt"

const (
	bucketCnt = 256
)

type entry struct {
	key   string
	value interface{}
}

type bucket struct {
	entries [][]interface{}
}

// Dumbhashmap : a basic implementation of Go's builtin hashmap
type Dumbhashmap struct {
	buckets []bucket
}

// New : create a new Dumbhashmap
func New() (h *Dumbhashmap) {
	return &Dumbhashmap{make([]bucket, bucketCnt)}
}

func hash(key string) uint32 {
	var (
		hash  = int32(5381)
		i     = len(key)
		runes = []rune(key)
	)

	for i > 0 {
		i--
		hash = int32(hash*33) ^ int32(runes[i])
	}

	return uint32(hash) % bucketCnt
}

// Get : retrieve the value stored for a given key.
func (h Dumbhashmap) Get(key string) interface{} {
	b := h.buckets[hash(key)]
	if len(b.entries) == 0 {
		return nil
	}

	var value interface{}
	for _, e := range b.entries {
		if e[0] == key {
			value = e[1]
		}
	}
	return value
}

// Set : set a value for the given key
func (h Dumbhashmap) Set(key string, value interface{}) {
	b := &h.buckets[hash(key)]
	b.entries = append(b.entries, []interface{}{key, value})
}

// Unset : unset any stored value for the given key
func (h Dumbhashmap) Unset(key string) (ok bool) {
	b := &h.buckets[hash(key)]
	for i, e := range b.entries {
		if e[0] == key {
			ok = true
			b.entries = append(b.entries[:i], b.entries[i+1:]...)
		}
	}
	return ok
}

func (h Dumbhashmap) String() string {
	var strValue = ""
	for i, b := range h.buckets {
		strValue += fmt.Sprintf("%v{", i)
		for j, e := range b.entries {
			strValue += fmt.Sprintf("\n  %v[%v, %v]\n", j, e[0], e[1])
		}
		strValue += "}\n"
	}

	if strValue == "" {
		return "<empty Dumbhashmap>"
	}
	return strValue
}
