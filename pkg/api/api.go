package api

import "net/http"

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}
