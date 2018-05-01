## dumbhashmap

This is a very barebones implementation of a hashmap in Go.

The main package runs a basic benchmark of dumbhashmap vs. Go's native hashmap ([map](https://golang.org/src/runtime/hashmap.go)) using the provided load constant.

dumbhashmap uses a fixed set of 256 buckets for object storage, and falls back to linear list iteration for hash collisions. It uses the hashing algorithm from [string-hash](https://www.npmjs.com/package/string-hash), which is a slightly modified version of `dbj2` by Dan Bernstein ([link](http://www.cse.yorku.ca/%7Eoz/hash.html)). It does not support any form of resizing, and currently only implements `Get`, `Set`, and `Unset`. Overwritten key-value pairs remain in their original buckets.

### Benchmark

![benchmark_1](img/15000.png)

![benchmark_2](img/150000.png)

As one would expect, `dumbhashmap` outperforms native `map` on Set operations, since it has much less overhead without having to account for resizing, shuffling, rehashing, etc. However, it is _massively_ outclassed on Get operations on large datasets.

### Performance

Proper hashmap implementations will perform at O(1), tending toward O(n) as _load factor_ increases.

[Load factor](https://en.wikipedia.org/wiki/Hash_table#Key_statistics) is defined as

```
n / k
```

where

```
n = number of entries in table
k = number of buckets
```

In a fixed length hashmap like `dumbhashmap`, load factor becomes a major drawback in large data sets, as n is very high while k is fixed. Proper implementations, like Go's native `map`, will dynamically resize to accomodate larger datasets any time a threshold average load factor is reached.

When storing a key-value pair, a hashmap will calculate a hash to find the corresponding bucket. It is possible that multiple distinct keys will produce the same hash, called a [hash collision](https://en.wikipedia.org/wiki/Hash_table#Collision_resolution). To allow for this possibility, the hashmap stores key-value pairs indirectly in buckets instead of directly in the core array. This requires some secondary form of search algorithm to find the correct value in a Get operation if it is not the only pair in the bucket.

`dumbhashmap` falls back on linear array search for bucket searching, while `map` uses a further map-like structure, electing to use the lower-order bits of the key hash for the core array index and a portion of the higher-order bits to distinguish pairs in a bucket ([reference](https://golang.org/src/runtime/hashmap.go#L9)). Linear array search is O(n), while map search is O(1), which gives `map` a further edge.
