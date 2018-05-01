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
	keys := make([]string, load)
	for i := 0; i < load; i++ {
		keys[i] = randstr.Get(6)
	}

	step := time.Now()
	dh := dumbhashmap.New()
	for i := 0; i < load; i++ {
		dh.Set(keys[i], randstr.Get(6))
	}
	fmt.Printf("dumbhashmap store %v items: %v\n", load, time.Since(step))

	step = time.Now()
	h := make(map[string]string)
	for i := 0; i < load; i++ {
		h[keys[i]] = randstr.Get(6)
	}
	fmt.Printf("native map store %v items: %v\n\n", load, time.Since(step))

	foo := ""
	step = time.Now()
	for i := 0; i < load; i++ {
		foo += dh.Get(keys[i]).(string)
	}
	fmt.Printf("dumbhashmap get %v items: %v\n", load, time.Since(step))

	step = time.Now()
	for i := 0; i < load; i++ {
		foo += h[keys[i]]
	}
	fmt.Printf("native map get %v items: %v\n", load, time.Since(step))
}
