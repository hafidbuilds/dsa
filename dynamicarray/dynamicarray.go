package dynamicarray

import (
	"fmt"
	"iter"
	"strings"

	"github.com/josestg/dsa/arrays"
)

type DynamicArray[T any] struct {
	backend *arrays.Array[T]
	length  int
}

func New[T any](capacity int) *DynamicArray[T] {
	if capacity <= 0 {
		panic("dynamicarray: must have at minimum 1 capacity")
	}
	return &DynamicArray[T]{
		backend: arrays.New[T](capacity),
		length:  0,
	}
}

func (d *DynamicArray[T]) Free() {
	if d.backend != nil {
		d.backend.Free()
		d.backend = nil
		d.length = 0
	}
}

func (d *DynamicArray[T]) Empty() bool { return d.Len() == 0 }

func (d *DynamicArray[T]) Len() int { return d.length }

func (d *DynamicArray[T]) Cap() int { return d.backend.Len() }

func (d *DynamicArray[T]) Get(index int) T {
	d.checkBounds(index)
	return d.backend.Get(index)
}

func (d *DynamicArray[T]) Set(index int, value T) {
	d.checkBounds(index)
	d.backend.Set(index, value)
}

func (d *DynamicArray[T]) Prepend(value T) {
	d.Append(value)
	for i := d.length - 1; i > 0; i-- {
		d.Swap(i, i-1)
	}
}

func (d *DynamicArray[T]) Shift() T {
	if v, ok := d.TryShift(); !ok {
		panic("dynamicarray: cannot shift from empty array")
	} else {
		return v
	}
}

func (d *DynamicArray[T]) TryShift() (T, bool) {
	var zero T
	n := d.Len()
	if n == 0 {
		return zero, false
	}

	v := d.Get(0)
	d.Set(0, zero)
	for i := 0; i < n-1; i++ {
		d.Swap(i, i+1)
	}
	d.length--
	return v, true
}

func (d *DynamicArray[T]) Swap(i, j int) {
	if i != j {
		x, y := d.Get(i), d.Get(j)
		d.Set(i, y)
		d.Set(j, x)
	}
}

func (d *DynamicArray[T]) Append(value T) {
	c := d.Cap()
	if d.length >= c {
		d.grow()
	}
	d.backend.Set(d.length, value)
	d.length++
}

func (d *DynamicArray[T]) grow() {
	// Go slice implementation only doubles the capacity if the current length is less than 256.
	// See: https://cs.opensource.google/go/go/+/refs/tags/go1.24.2:src/runtime/slice.go;l=289-322
	//
	// we can do the same with some this simple approximation: oldCap + (oldCap + 3*256) / 4
	// See: https://victoriametrics.com/blog/go-slice/
	const threshold = 256
	capacity := d.Cap()
	newCapacity := 2 * capacity
	if capacity >= threshold {
		newCapacity = capacity + (capacity+3*threshold)/4
	}
	newBackend := arrays.New[T](newCapacity)
	for i, v := range d.backend.Iter(false) {
		newBackend.Set(i, v)
	}
	d.backend.Free()
	d.backend = newBackend
}

func (d *DynamicArray[T]) Pop() T {
	if v, ok := d.TryPop(); !ok {
		panic("dynamicarray: cannot pop from empty list")
	} else {
		return v
	}
}

func (d *DynamicArray[T]) TryPop() (T, bool) {
	var zero T
	if d.Len() == 0 {
		return zero, false
	}
	val := d.backend.Get(d.length - 1)
	d.backend.Set(d.length-1, zero) // clear the slot.
	d.length--
	return val, true
}

func (d *DynamicArray[T]) Clip() {
	if d.Empty() {
		panic("dynamicarray: cannot clip on empty array")
	}

	if d.length == d.Cap() {
		return
	}

	newBackend := arrays.New[T](d.length)
	for i := range d.Len() {
		newBackend.Set(i, d.backend.Get(i))
	}
	d.backend.Free()
	d.backend = newBackend
}

func (d *DynamicArray[T]) Iter(reversed bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		if reversed {
			d.iterBackward(yield)
		} else {
			d.iterForward(yield)
		}
	}
}

func (d *DynamicArray[T]) iterForward(yield func(T) bool) {
	for i := range d.Len() {
		if !yield(d.Get(i)) {
			break
		}
	}
}

func (d *DynamicArray[T]) iterBackward(yield func(T) bool) {
	for i := d.Len() - 1; i >= 0; i-- {
		if !yield(d.Get(i)) {
			break
		}
	}
}

func (d *DynamicArray[T]) String() string {
	var buf strings.Builder
	buf.WriteRune('[')
	for i := range d.length {
		if i > 0 {
			buf.WriteRune(' ')
		}
		_, _ = fmt.Fprint(&buf, d.Get(i))
	}
	buf.WriteRune(']')
	return buf.String()
}

func (d *DynamicArray[T]) checkBounds(index int) {
	if index < 0 || index >= d.Len() {
		panic("dynamicarray: index out of range")
	}
}
