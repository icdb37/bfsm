package utils

type Getter[T any] interface {
	Get() T
}
type Setter[T any] interface {
	Set(T)
}

type FuncGetter[T any] func() T
type FuncSetter[T any] func(T)

type GetCache[T any] struct {
	state  bool
	info   *T
	getter FuncGetter[*T]
}

func NewGetCache[T any](g FuncGetter[*T]) *GetCache[T] {
	return &GetCache[T]{
		state:  false,
		info:   new(T),
		getter: g,
	}
}
func (c *GetCache[T]) Get() *T {
	if !c.state {
		c.info = c.getter()
		c.state = true
	}
	return c.info
}
