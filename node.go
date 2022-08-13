package v

import "io"

type Node interface {
	io.WriterTo
}

type NodeFunc func(w io.Writer) (int64, error)

func (f NodeFunc) WriteTo(w io.Writer) (int64, error) {
	return f(w)
}

type Nodes []Node

func (n Nodes) WriteTo(w io.Writer) (int64, error) {
	total := int64(0)
	for _, node := range n {
		if node == nil {
			continue
		}

		if n, err := node.WriteTo(w); err != nil {
			return total + n, err
		} else {
			total += n
		}
	}

	return total, nil
}

func Func(f func() (Node, error)) NodeFunc {
	return func(w io.Writer) (int64, error) {
		node, err := f()
		if err != nil {
			return 0, err
		}

		return node.WriteTo(w)
	}
}

func FromReader(r io.Reader) NodeFunc {
	return func(w io.Writer) (int64, error) {
		return io.Copy(w, r)
	}
}
