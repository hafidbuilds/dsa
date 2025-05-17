package dynamicarray_test

import (
	"fmt"
	"math"
	"slices"
	"testing"

	"github.com/josestg/dsa/dynamicarray"
	"github.com/stretchr/testify/assert"
)

func TestDynamicArray_EmptyState(t *testing.T) {
	assert.Panics(t, func() { dynamicarray.New[int](0) })
	assert.Panics(t, func() { dynamicarray.New[int](-1) })

	a := dynamicarray.New[int](1)
	t.Cleanup(a.Free)

	// empty state.
	assert.True(t, a.Empty())
	assert.Equal(t, 0, a.Len())
	assert.Equal(t, 1, a.Cap())
	assert.Equal(t, "[]", a.String())
	assert.Empty(t, slices.Collect(a.Iter(true)))
	assert.Empty(t, slices.Collect(a.Iter(false)))

	// must fail on empty state.
	assert.Panics(t, func() { a.Pop() })
	assert.Panics(t, func() { a.Shift() })
	assert.Panics(t, func() { a.Clip() })
	assert.Panics(t, func() { a.Get(0) })
	assert.Panics(t, func() { a.Set(0, 42) })
	assert.Panics(t, func() { a.Swap(0, 1) })

	// safe operation on empty state.
	shifted, ok := a.TryShift()
	assert.Zero(t, shifted)
	assert.False(t, ok)

	popped, ok := a.TryPop()
	assert.Zero(t, popped)
	assert.False(t, ok)
}

func TestDynamicArray_Append(t *testing.T) {
	t.Run("under doubling threshold", func(t *testing.T) {
		s := make([]int, 0, 1)
		a := dynamicarray.New[int](1)
		t.Cleanup(a.Free)

		n := 255 // the threshold - 1.
		for i := 0; i < n; i++ {
			v := 2*i + 1 // generating odd number.
			a.Append(v)
			s = append(s, v)
		}

		numOfGrow := int(math.Floor(math.Log2(float64(n)))) + 1
		assert.Equal(t, 1<<numOfGrow, a.Cap())
		assert.Equal(t, cap(s), a.Cap())
		assert.Equal(t, s, slices.Collect(a.Iter(false)))
		assert.Equal(t, fmt.Sprint(s), a.String())

		slices.Reverse(s)
		assert.Equal(t, s, slices.Collect(a.Iter(true)))

	})

	t.Run("hit the doubling threshold", func(t *testing.T) {
		s := make([]int, 0, 1)
		a := dynamicarray.New[int](1)
		t.Cleanup(a.Free)

		n := 513 // the threshold - 1.
		for i := 0; i < n; i++ {
			v := 2*i + 1 // generating odd number.
			a.Append(v)
			s = append(s, v)
		}

		delta := float64(max(cap(s), a.Cap())-min(cap(s), a.Cap())) / float64(max(cap(s), a.Cap()))
		assert.True(t, delta <= 0.1)
		assert.Equal(t, s, slices.Collect(a.Iter(false)))
		assert.Equal(t, fmt.Sprint(s), a.String())

		slices.Reverse(s)
		assert.Equal(t, s, slices.Collect(a.Iter(true)))
	})
}

func TestDynamicArray_Prepend(t *testing.T) {
	t.Run("under doubling threshold", func(t *testing.T) {
		s := make([]int, 0, 1)
		a := dynamicarray.New[int](1)
		t.Cleanup(a.Free)

		n := 255 // the threshold - 1.
		for i := 0; i < n; i++ {
			v := 2*i + 1 // generating odd number.
			a.Prepend(v)
			s = append(s, v)
		}

		numOfGrow := int(math.Floor(math.Log2(float64(n)))) + 1
		assert.Equal(t, 1<<numOfGrow, a.Cap())
		assert.Equal(t, cap(s), a.Cap())
		assert.Equal(t, s, slices.Collect(a.Iter(true)))

		slices.Reverse(s)
		assert.Equal(t, fmt.Sprint(s), a.String())
	})

	t.Run("hit the doubling threshold", func(t *testing.T) {
		s := make([]int, 0, 1)
		a := dynamicarray.New[int](1)
		t.Cleanup(a.Free)

		n := 513 // the threshold - 1.
		for i := 0; i < n; i++ {
			v := 2*i + 1 // generating odd number.
			a.Prepend(v)
			s = append(s, v)
		}

		delta := float64(max(cap(s), a.Cap())-min(cap(s), a.Cap())) / float64(max(cap(s), a.Cap()))
		assert.True(t, delta <= 0.1)
		assert.Equal(t, s, slices.Collect(a.Iter(true)))

		slices.Reverse(s)
		assert.Equal(t, fmt.Sprint(s), a.String())
	})
}

func TestDynamicArray_Pop(t *testing.T) {
	a := dynamicarray.New[int](1)
	t.Cleanup(a.Free)

	n := 255 // the threshold - 1.
	for i := 0; i < n; i++ {
		v := 2*i + 1 // generating odd number.
		a.Append(v)
	}

	exp := slices.Collect(a.Iter(true))
	got := make([]int, 0, 1)
	for !a.Empty() {
		got = append(got, a.Pop())
	}

	assert.Equal(t, exp, got)
}

func TestDynamicArray_Shift(t *testing.T) {
	a := dynamicarray.New[int](1)
	t.Cleanup(a.Free)

	n := 255 // the threshold - 1.
	for i := 0; i < n; i++ {
		v := 2*i + 1 // generating odd number.
		a.Append(v)
	}

	exp := slices.Collect(a.Iter(false))
	got := make([]int, 0, 1)
	for !a.Empty() {
		got = append(got, a.Shift())
	}
	assert.Equal(t, exp, got)
}

func TestDynamicArray_Clip(t *testing.T) {
	t.Run("cap > len", func(t *testing.T) {
		a := dynamicarray.New[int](1)
		t.Cleanup(a.Free)

		n := 10 // the threshold - 1.
		for i := 0; i < n; i++ {
			v := 2*i + 1 // generating odd number.
			a.Append(v)
		}

		assert.NotEqual(t, a.Cap(), a.Len())
		a.Clip()
		assert.Equal(t, a.Cap(), a.Len())
	})

	t.Run("cap == len", func(t *testing.T) {
		a := dynamicarray.New[int](1)
		t.Cleanup(a.Free)

		n := 4 // the threshold - 1.
		for i := 0; i < n; i++ {
			v := 2*i + 1 // generating odd number.
			a.Append(v)
		}

		assert.Equal(t, a.Cap(), a.Len())
		a.Clip()
		assert.Equal(t, a.Cap(), a.Len())
	})
}
