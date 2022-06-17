package net

import "net/http"

// health 健康检查 handler
func health(w http.ResponseWriter, r *http.Request)  {
	_, _ = w.Write([]byte(""))
}
