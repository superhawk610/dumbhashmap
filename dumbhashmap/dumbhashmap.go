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

// dumbhashmap : a basic implementation of Go's builtin hashmap
type dumbhashmap struct {
	buckets []bucket
}

// New : create a new dumbhashmap
func New() (h *dumbhashmap) {
	return &dumbhashmap{make([]bucket, bucketCnt)}
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
func (h dumbhashmap) Get(key string) interface{} {
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
func (h dumbhashmap) Set(key string, value interface{}) {
	b := &h.buckets[hash(key)]
	b.entries = append(b.entries, []interface{}{key, value})
}

func (h dumbhashmap) String() string {
	var strValue = ""
	for i, b := range h.buckets {
		strValue += fmt.Sprintf("%v{", i)
		for j, e := range b.entries {
			strValue += fmt.Sprintf("\n  %v[%v, %v]\n", j, e[0], e[1])
		}
		strValue += "}\n"
	}

	if strValue == "" {
		return "<empty dumbhashmap>"
	}
	return strValue
}
