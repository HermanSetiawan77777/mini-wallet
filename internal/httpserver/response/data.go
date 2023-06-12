package response

import "net/http"

func WithData(w http.ResponseWriter, status int, data any) {
	withJSON(w, status, map[string]any{
		"data": data,
	})
}
