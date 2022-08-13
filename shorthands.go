package v

import (
	"fmt"
	"io"
	"strings"

	"github.com/samber/lo"
)

func Document(doctype, lang string, head Node, body Node) Nodes {
	return Nodes{
		Doctype(doctype),
		HTML(Attr{"lang": lang},
			head,
			body,
		),
	}
}

func If(condition bool, f NodeFunc) NodeFunc {
	return func(w io.Writer) (int64, error) {
		if condition {
			return f.WriteTo(w)
		}

		return 0, nil
	}
}

func Not(condition bool, f NodeFunc) NodeFunc {
	return If(!condition, f)
}

func IfErr(err error, f NodeFunc) NodeFunc {
	return If(err != nil, f)
}

type Classes map[string]any

func (c Classes) String() string {
	return strings.Join(lo.Keys(lo.PickBy(c, func(_ string, value any) bool { return !nilOrFalse(value) })), " ")
}

type CSS map[string]any

func (c CSS) String() string {
	return strings.Join(lo.Map(lo.Entries(c), func(entry lo.Entry[string, any], _ int) string { return fmt.Sprintf("%s:%v;", entry.Key, entry.Value) }), "")
}

func JS(src string, attr ...Attr) Element {
	return Script(Attr{"src": src, "type": "application/javascript"}.Merge(attr...))
}

func Stylesheet(href string, attr ...Attr) Element {
	return Link(Attr{"href": href, "rel": "stylesheet"}.Merge(attr...))
}

func Datas(attr Attr) Attr {
	return PrefixAttr("data-", attr)
}

func PrefixAttr(prefix string, attr Attr) Attr {
	return lo.MapKeys(attr, func(_ any, key string) string {
		return fmt.Sprintf("%s%s", prefix, key)
	})
}
