package v

import (
	"fmt"
	"io"
)

func Tag(name string, attr Attr, nodes ...Node) Element {
	return Element{Name: name, Attributes: attr, Nodes: nodes}
}

func Doctype(doctype string) NodeFunc {
	return func(w io.Writer) (int64, error) {
		n, err := fmt.Fprintf(w, "<!doctype %s>", doctype)
		return int64(n), err
	}
}
