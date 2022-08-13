package v

import (
	"fmt"
	"html"
	"io"
	"strings"

	"github.com/samber/lo"
)

type Attr map[string]any

func (a Attr) WriteTo(w io.Writer) (int64, error) {
	attrs := lo.FilterMap(lo.Entries(a), func(entry lo.Entry[string, any], _ int) (string, bool) {
		if nilOrFalse(entry.Value) {
			return "", false
		}

		return fmt.Sprintf("%s=%q", html.EscapeString(entry.Key), html.EscapeString(fmt.Sprint(entry.Value))), true
	})

	n, err := w.Write([]byte(strings.Join(attrs, " ")))
	return int64(n), err
}

func Attrs(attr ...Attr) Attr {
	return lo.Assign(lo.Map(attr, func(attr Attr, i int) map[string]any { return attr })...)
}

func (a Attr) Merge(attr ...Attr) Attr {
	return Attrs(append([]Attr{a}, attr...)...)
}

func nilOrFalse(value any) bool {
	if value == nil {
		return true
	}

	if b, ok := value.(bool); ok && !b {
		return true
	}

	return false
}
