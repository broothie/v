package v

import (
	"fmt"
	"html"
	"io"
)

func Text(text string) NodeFunc {
	return func(w io.Writer) (int64, error) {
		n, err := fmt.Fprint(w, html.EscapeString(text))
		return int64(n), err
	}
}
