package handlers

import "net/http"

// Ping godoc
// @Router /ping [get]
// @Summary Ping.
// @Description server status.
// @Success 200
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
