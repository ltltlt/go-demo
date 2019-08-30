// http://npat-efault.github.io/programming/2016/10/10/escape-analysis-and-interfaces.html
package main

import "os"
import "net"

func Ok(f os.File) []byte {
	var x [128]byte
	println(&x)
	b := x[:]
	n, _ := f.Read(b)
	r := make([]byte, n)
	copy(r, b[:n])
	return r
}

func NotOk(c net.Conn) []byte {
	var x [128]byte
	println(&x)
	b := x[:]
	n, _ := c.Read(b)
	r := make([]byte, n)
	copy(r, b[:n])
	return r
}

func main() {
	f, _ := os.Open("test31.go")
	Ok(*f)
	conn, _ := net.Dial("tcp", "bing.com:443")
	NotOk(conn)
}

type S struct {
	s1 int
}

func (s *S) M1(i int) { s.s1 = i }

type I interface {
	M1(int)
}

func g() {
	var s1 S // this escapes
	var s2 S // this does not
	f1(&s1)
	f2(&s2)
}

func f1(s I)  { s.M1(42) }
func f2(s *S) { s.M1(42) }
