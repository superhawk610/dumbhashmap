package main

import (
	"fmt"
	"time"

	"github.com/superhawk610/hashmap/dumbhashmap"
	"github.com/superhawk610/hashmap/randstr"
)

const (
	load = 150000
)

func main() {
	benchmark()
}

func benchmark() {
	keys := make([]string, load)
	for i := 0; i < load; i++ {
		keys[i] = randstr.Get(6)
	}

	/* ---------- Set ---------- */
	step := time.Now()
	dh := dumbhashmap.New()
	for i := 0; i < load; i++ {
		dh.Set(keys[i], keys[i])
	}
	fmt.Printf("dumbhashmap store %v items: %v\n", load, time.Since(step))

	step = time.Now()
	m := make(map[string]string)
	for i := 0; i < load; i++ {
		m[keys[i]] = keys[i]
	}
	fmt.Printf(" native map store %v items: %v\n\n", load, time.Since(step))

	/* ---------- Get ---------- */
	var foo string
	step = time.Now()
	for i := 0; i < load; i++ {
		foo = dh.Get(keys[i]).(string)
	}
	fmt.Printf("  dumbhashmap get %v items: %v\n", load, time.Since(step))

	step = time.Now()
	for i := 0; i < load; i++ {
		foo = m[keys[i]]
	}
	fmt.Printf("   native map get %v items: %v\n\n", load, time.Since(step))
	_ = foo

	/* --------- Unset --------- */
	step = time.Now()
	for i := 0; i < load; i++ {
		dh.Unset(keys[i])
	}
	fmt.Printf("dumbhashmap unset %v items: %v\n", load, time.Since(step))

	for i := 0; i < load; i++ {
		delete(m, keys[i])
	}
	fmt.Printf(" native map unset %v items: %v\n", load, time.Since(step))
}
