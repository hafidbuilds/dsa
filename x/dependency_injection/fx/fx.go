package fx

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/josestg/dsa/graph"
)

type Container struct {
	g            *graph.Graph[reflect.Type]
	instances    map[reflect.Type]reflect.Value
	constructors map[reflect.Type]reflect.Value
}

func NewContainer() *Container {
	return &Container{
		g:            graph.New[reflect.Type](true),
		instances:    make(map[reflect.Type]reflect.Value),
		constructors: make(map[reflect.Type]reflect.Value),
	}
}

func (c *Container) Provide(constructor any) error {
	val := reflect.ValueOf(constructor)
	typ := val.Type()
	if typ.Kind() != reflect.Func || typ.NumOut() != 1 {
		return fmt.Errorf("constructor must be a function returning 1 value")
	}

	out := typ.Out(0)
	if _, exists := c.constructors[out]; exists {
		return fmt.Errorf("constructor for %v already provided", out)
	}

	for i := 0; i < typ.NumIn(); i++ {
		c.g.AddEdge(typ.In(i), out)

	}

	c.constructors[out] = val
	return nil
}

func (c *Container) Build() error {
	var topoSorted []reflect.Type
	t := graph.NewWalker(c.g, graph.DFSPostOrder)
	if t.HasCycle() {
		return fmt.Errorf("cycle dependency detected")
	}

	t.WalkAll(func(n reflect.Type) { topoSorted = append(topoSorted, n) })

	slices.Reverse(topoSorted)
	for _, t := range topoSorted {
		ctor, ok := c.constructors[t]
		if !ok {
			return fmt.Errorf("misisng constructor")
		}

		var args []reflect.Value
		for j := 0; j < ctor.Type().NumIn(); j++ {
			depType := ctor.Type().In(j)
			dep, ok := c.instances[depType]
			if !ok {
				return fmt.Errorf("missing dependency: %v", depType)
			}
			args = append(args, dep)
		}

		result := ctor.Call(args)
		c.instances[t] = result[0]
	}

	return nil
}

func Resolve[T any](c *Container) (T, error) {
	var zero T
	t := reflect.TypeOf((*T)(nil)).Elem()
	val, ok := c.instances[t]
	if !ok {
		return zero, fmt.Errorf("type not found: %v", t)
	}
	return val.Interface().(T), nil
}
