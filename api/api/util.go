package api

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
)

// json response
func respond(w http.ResponseWriter, data any) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// remote address
func getAddress(r *http.Request) string {
	address := r.Header.Get("x-forwarded-for")
	address = strings.TrimSpace(strings.Split(address, ",")[len(strings.Split(address, ","))-1])
	if address == "" {
		address, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return address
}
