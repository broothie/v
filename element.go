package v

import (
	"fmt"
	"io"
)

type Element struct {
	Name       string
	Attributes Attr
	Nodes      Nodes
}

func (e Element) WriteTo(w io.Writer) (int64, error) {
	total := int64(0)

	if n, err := fmt.Fprintf(w, "<%s", e.Name); err != nil {
		return total + int64(n), err
	} else {
		total += int64(n)
	}

	if len(e.Attributes) > 0 {
		if n, err := fmt.Fprint(w, " "); err != nil {
			return total + int64(n), err
		} else {
			total += int64(n)
		}

		if n, err := e.Attributes.WriteTo(w); err != nil {
			return total + n, err
		} else {
			total += n
		}
	}

	if n, err := fmt.Fprint(w, ">"); err != nil {
		return total + int64(n), err
	} else {
		total += int64(n)
	}

	if n, err := e.Nodes.WriteTo(w); err != nil {
		return total + n, err
	} else {
		total += n
	}

	if n, err := fmt.Fprintf(w, "</%s>", e.Name); err != nil {
		return total + int64(n), err
	} else {
		total += int64(n)
	}

	return total, nil
}
