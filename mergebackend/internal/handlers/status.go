package handlers

import (
	"io"
	"net/http"
)

func HandleGetStatus(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Ready")
}
