// Basic example that demonstrates a type that doesn't implement an interface
// being flagged for a compile error when assigned to the interface.
package main

type Munger interface {
	Munge(int)
}

type Foo int
type Bar int

func (f Foo) Munge(int) {
}

func main() {
	var m Munger
	var f Foo
	m = f

	var b Bar
	m = b
}