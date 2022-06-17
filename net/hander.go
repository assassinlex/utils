package net

import "net/http"

// Health  健康检查 handler
func Health(w http.ResponseWriter, r *http.Request)  {
	_, _ = w.Write([]byte(""))
}
