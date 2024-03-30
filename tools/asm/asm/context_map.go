package asm

const (
	globalContext = ""
)

type ContextMap[T any] struct {
	ctx    []string
	values map[string]map[string]T
}

func NewContextMap[T any]() ContextMap[T] {
	return ContextMap[T]{
		ctx: []string{globalContext},
		values: map[string]map[string]T{
			globalContext: {},
		},
	}
}

func GetContextName(ptr SrcPointer) string {
	return ptr.String()
}

func (d *ContextMap[T]) PushContext(ctx string) {
	d.ctx = append(d.ctx, ctx)
}

func (d *ContextMap[T]) PopContext(ctx string) {
	if len(d.ctx) <= 1 { // note: do not pop the global context
		return
	}
	d.ctx = d.ctx[:len(d.ctx)-1]
}

func (d *ContextMap[T]) Set(name string, value T) {
	// todo: refactor this, the first element of ctx is always the globalContext
	topCtx := globalContext
	if len(d.ctx) > 0 {
		topCtx = d.ctx[len(d.ctx)-1]
	}

	if _, ok := d.values[topCtx]; !ok {
		d.values[topCtx] = map[string]T{}
	}

	d.values[topCtx][name] = value
}

func (c *ContextMap[T]) Get(name string) (v T, ok bool) {
	for i := len(c.ctx) - 1; i >= 0; i-- {
		if v, ok := c.values[c.ctx[i]][name]; ok {
			return v, true
		}
	}

	return v, false
}

func (d *ContextMap[T]) SetOnce(name string, value T) {
	if _, ok := d.Get(name); !ok {
		d.Set(name, value)
	}
}

func (d *ContextMap[T]) List() map[string]T {
	topCtx := d.ctx[len(d.ctx)-1]
	return d.values[topCtx]
}
