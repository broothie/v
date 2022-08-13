package v

import (
	"io"
	"net/http"
)

func Render(w http.ResponseWriter, statusCode int, nodes ...Node) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(statusCode)

	if _, err := WriteTo(w, nodes...); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WriteTo(w io.Writer, nodes ...Node) (int64, error) {
	return Nodes(nodes).WriteTo(w)
}
