package main

import "encoding/binary"
import "sync/atomic"

func putuint64(b []byte, v uint64) {
	_ = b[7]
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
}

func test(a []int) {
	_ = a[3]
	v := a[2]
	print(v)
	v = a[3]
	print(v)
	a[3] = 2
}

func swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func main() {
	test([]int(nil))
	bts := [100]byte{}
	binary.BigEndian.PutUint64(bts[:], 122222)
	putuint64(bts[:], 12222)

	swap(1, 2)

	testload()
}

func testload() int64 {
	var v int64
	v = atomic.LoadInt64(&v)
	return v
}
