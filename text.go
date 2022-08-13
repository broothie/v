package v

import (
	"html"
	"io"
)

func Text(text string) NodeFunc {
	return func(w io.Writer) (int64, error) {
		n, err := w.Write([]byte(html.EscapeString(text)))
		return int64(n), err
	}
}
