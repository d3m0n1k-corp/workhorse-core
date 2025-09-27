package pools

import (
	"bytes"
	"encoding/json"
	"sync"
)

// ObjectPool manages reusable resources to reduce GC pressure
type ObjectPool struct {
	buffers      sync.Pool
	jsonEncoders sync.Pool
	jsonDecoders sync.Pool
}

// GlobalPool is the shared instance
var GlobalPool = NewObjectPool()

// NewObjectPool creates a new object pool
func NewObjectPool() *ObjectPool {
	return &ObjectPool{
		buffers: sync.Pool{
			New: func() any {
				return &bytes.Buffer{}
			},
		},
		jsonEncoders: sync.Pool{
			New: func() any {
				buf := &bytes.Buffer{}
				return json.NewEncoder(buf)
			},
		},
		jsonDecoders: sync.Pool{
			New: func() any {
				return json.NewDecoder(&bytes.Buffer{})
			},
		},
	}
}

// GetBuffer returns a clean buffer from the pool
func (p *ObjectPool) GetBuffer() *bytes.Buffer {
	buf := p.buffers.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

// PutBuffer returns a buffer to the pool
func (p *ObjectPool) PutBuffer(buf *bytes.Buffer) {
	// Prevent memory leak from very large buffers
	if buf.Cap() > 64*1024 { // 64KB limit
		return
	}
	p.buffers.Put(buf)
}

// WithBuffer executes a function with a pooled buffer
func (p *ObjectPool) WithBuffer(fn func(*bytes.Buffer) error) error {
	buf := p.GetBuffer()
	defer p.PutBuffer(buf)
	return fn(buf)
}
